package model

import (
	"fmt"
	"pandax/kit/biz"
	"pandax/kit/starter"
	"strconv"

	"strings"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	BaseAutoModel
	CreatorId  int64  `json:"creatorId"`
	Creator    string `json:"creator"` //创建者
	ModifierId int64  `json:"modifierId"`
	Modifier   string `json:"modifier"` //修改者
}

// 设置基础信息. 如创建时间，修改时间，创建者，修改者信息
func (m *Model) SetBaseInfo(account *LoginAccount) {
	nowTime := time.Now()
	isCreate := m.Id == 0
	if isCreate {
		m.CreatedAt = nowTime
	}
	m.UpdatedAt = nowTime

	if account == nil {
		return
	}
	id := account.UserId
	name := account.Username
	if isCreate {
		m.CreatorId = id
		m.Creator = name
	}
	m.Modifier = name
	m.ModifierId = id
}

// 事务
func Tx(funcs ...func(db *gorm.DB) error) (err error) {
	tx := starter.Db.Begin()
	defer func() {
		var err any
		err = recover()
		if err != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", err)
		}
	}()
	for _, f := range funcs {
		err = f(tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	err = tx.Commit().Error
	return
}

// 根据id获取实体对象。model需为指针类型（需要将查询出来的值赋值给model）
//
// 若error不为nil则为不存在该记录
func GetById(model any, id uint64, cols ...string) error {
	return starter.Db.Select(cols).Where("id = ?", id).First(model).Error
}

// 根据id列表查询
func GetByIdIn(model any, list any, ids []uint64, orderBy ...string) {
	var orderByStr string
	if orderBy == nil {
		orderByStr = "id desc"
	} else {
		orderByStr = strings.Join(orderBy, ",")
	}
	starter.Db.Model(model).Where("id in (?)", ids).Order(orderByStr).Find(list)
}

// 根据id列表查询 model可以是对象，也可以是map[string]interface{}
func CountBy(model any) int64 {
	var count int64
	starter.Db.Model(model).Where(model).Count(&count)
	return count
}

// 根据id更新model，更新字段为model中不为空的值，即int类型不为0，ptr类型不为nil这类字段值
func UpdateById(model any) error {
	return starter.Db.Model(model).Updates(model).Error
}
func UpdateByWhere(model any, where any) error {
	return starter.Db.Model(model).Where(where).Updates(model).Error
}

// 根据id删除model
func DeleteById(model any, id uint64) error {
	return starter.Db.Delete(model, "id = ?", id).Error
}

// 根据条件删除
func DeleteByCondition(model any) error {
	return starter.Db.Where(model).Delete(model).Error
}

// 插入model
func Insert(model any) error {
	return starter.Db.Create(model).Error
}

// 获取满足model中不为空的字段值条件的所有数据.
//
// list为数组类型 如 var users *[]User，可指定为非model结构体，即只包含需要返回的字段结构体
func ListBy(model any, list any, cols ...string) {
	starter.Db.Model(model).Select(cols).Where(model).Find(list)
}

// 获取满足model中不为空的字段值条件的所有数据.
//
// list为数组类型 如 var users *[]User，可指定为非model结构体
func ListByOrder(model any, list any, order ...string) {
	var orderByStr string
	if order == nil {
		orderByStr = "id desc"
	} else {
		orderByStr = strings.Join(order, ",")
	}
	starter.Db.Model(model).Where(model).Order(orderByStr).Find(list)
}

// 获取满足model中不为空的字段值条件的单个对象。model需为指针类型（需要将查询出来的值赋值给model）
//
// 若 error不为nil，则为不存在该记录
func GetBy(model any, cols ...string) error {
	return starter.Db.Select(cols).Where(model).First(model).Error
}

// 获取满足conditionModel中不为空的字段值条件的单个对象。model需为指针类型（需要将查询出来的值赋值给model）
// toModel  需要查询的字段
// 若 error不为nil，则为不存在该记录
func GetByConditionTo(conditionModel any, toModel any) error {
	return starter.Db.Model(conditionModel).Where(conditionModel).First(toModel).Error
}

// 获取分页结果
func GetPage(pageParam *PageParam, conditionModel any, toModels any, orderBy ...string) *ResultPage {
	var count int64
	err := starter.Db.Model(conditionModel).Where(conditionModel).Count(&count).Error
	biz.ErrIsNilAppendErr(err, " 查询错误：%s")
	if count == 0 {
		return &ResultPage{Total: 0, Data: []string{}}
	}

	page := pageParam.PageNum
	pageSize := pageParam.PageSize
	var orderByStr string
	if orderBy == nil {
		orderByStr = "id desc"
	} else {
		orderByStr = strings.Join(orderBy, ",")
	}
	err = starter.Db.Model(conditionModel).Where(conditionModel).Order(orderByStr).Limit(pageSize).Offset((page - 1) * pageSize).Find(toModels).Error
	biz.ErrIsNil(err, "查询失败")
	return &ResultPage{Total: count, Data: toModels}
}

// 根据sql获取分页对象
func GetPageBySql(sql string, param *PageParam, toModel any, args ...any) *ResultPage {
	db := starter.Db
	selectIndex := strings.Index(sql, "SELECT ") + 7
	fromIndex := strings.Index(sql, " FROM")
	selectCol := sql[selectIndex:fromIndex]
	countSql := strings.Replace(sql, selectCol, "COUNT(*) AS total ", 1)
	// 查询count
	var count int
	db.Raw(countSql, args...).Scan(&count)
	if count == 0 {
		return &ResultPage{Total: 0, Data: []string{}}
	}
	// 分页查询
	limitSql := sql + " LIMIT " + strconv.Itoa(param.PageNum-1) + ", " + strconv.Itoa(param.PageSize)
	err := db.Raw(limitSql).Scan(toModel).Error
	biz.ErrIsNil(err, "查询失败")
	return &ResultPage{Total: int64(count), Data: toModel}
}

func GetListBySql(sql string, params ...any) []map[string]any {
	var maps []map[string]any
	starter.Db.Raw(sql, params).Scan(&maps)
	return maps
}

func GetListBySql2Model(sql string, toEntity any, params ...any) error {
	return starter.Db.Raw(sql, params).Find(toEntity).Error
}
