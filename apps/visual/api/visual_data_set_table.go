package api

// ==========================================================================
// 生成日期：2023-04-10 03:05:21 +0800 CST
// 生成路径: apps/visual/api/visual_data_set_table.go
// 生成人：panda
// ==========================================================================
import (
	"encoding/json"
	"fmt"
	"github.com/XM-GO/PandaKit/biz"
	filek "github.com/XM-GO/PandaKit/file"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
	"pandax/apps/visual/driver"
	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
	"pandax/pkg/global"
	"pandax/pkg/tool"
	"path"
	"strings"
	"time"
)

const filePath = "uploads/excel"

var ColumnTypeStr = []string{"delete_time", "update_time"}

type VisualDataSetTableApi struct {
	VisualDataSetTableApp services.VisualDataSetTableModel
	VisualDataSourceApp   services.VisualDataSourceModel
	VisualDataSetFieldApi services.VisualDataSetFieldModel
}

// GetVisualDataSetTableList DataSetTable列表数据
func (p *VisualDataSetTableApi) GetVisualDataSetTableList(rc *restfulx.ReqCtx) {
	data := entity.VisualDataSetTable{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")
	data.TableType = restfulx.QueryParam(rc, "tableType")
	list, total := p.VisualDataSetTableApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualDataSetTableListAll DataSetTable列表数据
func (p *VisualDataSetTableApi) GetVisualDataSetTableListAll(rc *restfulx.ReqCtx) {
	data := entity.VisualDataSetTable{}
	data.Name = restfulx.QueryParam(rc, "name")
	data.TableType = restfulx.QueryParam(rc, "tableType")
	rc.ResData = p.VisualDataSetTableApp.FindList(data)
}

// GetVisualDataSetTable 获取DataSetTable
func (p *VisualDataSetTableApi) GetVisualDataSetTable(rc *restfulx.ReqCtx) {
	tableId := restfulx.PathParam(rc, "tableId")
	rc.ResData = p.VisualDataSetTableApp.FindOne(tableId)
}

// InsertVisualDataSetTable 添加DataSetTable
func (p *VisualDataSetTableApi) InsertVisualDataSetTable(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetTable
	restfulx.BindQuery(rc, &data)
	data.TableId = kgo.KStr.Uniqid("px")
	p.VisualDataSetTableApp.Insert(data)
	// 协程执行添加字段
	go p.operateTable("add", data)
}

// UpdateVisualDataSetTable 修改DataSetTable
func (p *VisualDataSetTableApi) UpdateVisualDataSetTable(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetTable
	restfulx.BindQuery(rc, &data)
	p.VisualDataSetTableApp.Update(data)

	go p.operateTable("edit", data)
}

// DeleteVisualDataSetTable 删除DataSetTable
func (p *VisualDataSetTableApi) DeleteVisualDataSetTable(rc *restfulx.ReqCtx) {
	tableId := restfulx.PathParam(rc, "tableId")
	tableIds := strings.Split(tableId, ",")
	p.VisualDataSetTableApp.Delete(tableIds)
}

// GetVisualDataSetTableResult 获取运行结果
func (p *VisualDataSetTableApi) GetVisualDataSetTableResult(rc *restfulx.ReqCtx) {
	dsr := entity.VisualDataSetRequest{}
	restfulx.BindQuery(rc, &dsr)
	info := make(map[string]string)
	biz.ErrIsNil(json.Unmarshal([]byte(dsr.Info), &info), "获取表信息失败")
	if dsr.Type == "excel" {
		filePath := info["filePath"]
		fields, data := tool.ReadExcel(filePath)
		rc.ResData = entity.VisualDataSetRes{
			Data:   data,
			Fields: fields,
		}
		return
	}
	one := p.VisualDataSourceApp.FindOne(dsr.SourceId)
	instance := driver.NewDbInstance(one)
	var sql string
	if dsr.Type == "db" {
		sql = fmt.Sprintf(`select * from %s LIMIT 10`, info["table"])
	}
	if dsr.Type == "sql" {
		sql = info["sql"] + " LIMIT 10"
	}
	fields, data, err := instance.SelectData(sql)
	biz.ErrIsNil(err, "查询表信息失败")
	rc.ResData = entity.VisualDataSetRes{
		Data:   data,
		Fields: fields,
	}
}

// GetVisualDataSetTableByExcelResult 通过excel读取结果
func (p *VisualDataSetTableApi) GetVisualDataSetTableByExcelResult(rc *restfulx.ReqCtx) {
	_, fileHeader, err := rc.Request.Request.FormFile("excelFile")
	ext := path.Ext(fileHeader.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(fileHeader.Filename, ext)
	name = kgo.KStr.Md5(name, 10)
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	dst := fmt.Sprintf("%s/%s", filePath, filename)
	filek.SaveUploadedFile(fileHeader, dst)
	biz.ErrIsNil(err, "文件上传失败")
	rc.ResData = dst
}

// GetDataSetTableFun 获取sql函数
func (p *VisualDataSetTableApi) GetDataSetTableFun(rc *restfulx.ReqCtx) {
	sourceId := restfulx.QueryParam(rc, "sourceId")
	one := p.VisualDataSourceApp.FindOne(sourceId)
	rc.ResData = p.VisualDataSetTableApp.FindFunList(one.SourceType)
}

// GetDataSetSeries 获取图表设置的数据集查询数据
func (p *VisualDataSetTableApi) GetDataSetSeries(rc *restfulx.ReqCtx) {
	var data entity.DataSetDataReq
	restfulx.BindQuery(rc, &data)
	one := p.VisualDataSetTableApp.FindOne(data.TableId)
	biz.NotNil(one, "未查到该数据集")
	info := make(map[string]string)
	biz.ErrIsNil(json.Unmarshal([]byte(one.Info), &info), "解析info失败")
	// 数据源查询
	source := p.VisualDataSourceApp.FindOne(one.DataSourceId)
	// 拼接维度
	dataDs := make([]string, 0)
	dataQs := make([]string, 0)
	sorts := make([]string, 0)
	filters := make([]string, 0)

	for _, ds := range data.DataDs {
		dataDs = append(dataDs, ds.Value)
		dataQs = append(dataQs, ds.Value)
	}
	// 拼接指标
	for _, qs := range data.DataQs {
		q := qs.Value
		rename := qs.Func + "_" + qs.Value
		if qs.Func != "" {
			q = fmt.Sprintf("%s(%s)", qs.Func, q)
		}

		if qs.Decimal > 0 {
			q = fmt.Sprintf("ROUND(%s,%d) as %s", q, qs.Decimal, rename)
		} else {
			q = fmt.Sprintf("%s as %s", q, rename)
		}

		if qs.Sort != "" {
			//升序
			if qs.Sort == "2" {
				sorts = append(sorts, fmt.Sprintf("%s asc", rename))
			}
			//降序
			if qs.Sort == "3" {
				sorts = append(sorts, fmt.Sprintf("%s desc", rename))
			}
		}
		if len(qs.Filters) > 0 {
			for _, filter := range qs.Filters {
				filters = append(filters, fmt.Sprintf("%s %s %d", filter.Name, filter.Rule, filter.Num))
			}
		}
		dataQs = append(dataQs, q)
	}

	//
	if one.TableType == "excel" {
		rc.ResData = "占时不支持"
		return
	}

	var sql string
	//根据信息拼接sql
	if one.TableType == "db" {
		sql = fmt.Sprintf(`select %s from %s`,
			strings.Join(dataQs, ","),
			info["table"],
		)
	}
	if one.TableType == "sql" {
		sql = fmt.Sprintf(`select %s from (%s) as table`,
			strings.Join(dataQs, ","),
			info["sql"],
		)
	}
	if len(filters) > 0 {
		sql += fmt.Sprintf(" WHERE %s", strings.Join(filters, " AND "))
	}
	if len(dataDs) > 0 {
		sql += fmt.Sprintf(" GROUP BY %s", strings.Join(dataDs, ","))
	}
	if len(sorts) > 0 {
		sql += fmt.Sprintf(" ORDER BY %s", strings.Join(sorts, ","))
	}
	if data.ShowNumType == "2" {
		sql += fmt.Sprintf(" LIMIT %d", data.ShowNum)
	}
	// 执行
	instance := driver.NewDbInstance(source)
	_, series, err := instance.SelectData(sql)
	biz.ErrIsNil(err, "查询表信息失败")
	rc.ResData = series
}

func (p *VisualDataSetTableApi) operateTable(opType string, data entity.VisualDataSetTable) {
	info := make(map[string]string)
	err := json.Unmarshal([]byte(data.Info), &info)
	if err != nil {
		return
	}
	// 保存 excel field
	if data.TableType == "excel" {
		filePath := info["filePath"]
		_, datas := tool.ReadExcel(filePath)
		fields := getDataSetField(datas)
		if len(fields) > 0 {
			for _, field := range fields {
				field.TableId = data.TableId
				if opType == "edit" {
					p.VisualDataSetFieldApi.DeleteByTable(data.TableId)
				}
				p.VisualDataSetFieldApi.Insert(field)
			}
		}
		return
	}

	one := p.VisualDataSourceApp.FindOne(data.DataSourceId)
	if driver.TestConnection(one) != nil {
		one.Status = "0"
		one := p.VisualDataSourceApp.Update(*one)
		global.Log.Errorf("数据源:%s不在线", one.SourceName)
		return
	}
	instance := driver.NewDbInstance(one)

	sql := ""
	if data.TableType == "db" {
		sql = fmt.Sprintf(`select * from %s LIMIT 1`, info["table"])
	}
	if data.TableType == "sql" {
		sql = info["sql"] + " LIMIT 1"
	}
	_, datas, err := instance.SelectData(sql)

	fields := getDataSetField(datas)
	if len(fields) > 0 {
		for _, field := range fields {
			field.TableId = data.TableId
			if opType == "edit" {
				p.VisualDataSetFieldApi.DeleteByTable(data.TableId)
			}
			p.VisualDataSetFieldApi.Insert(field)
		}
	}
}

// 根据读取的Data获取字段
func getDataSetField(datas []map[string]interface{}) []entity.VisualDataSetField {
	fields := make([]entity.VisualDataSetField, 0)
	if len(datas) > 0 {
		for k, v := range datas[0] {
			field := entity.VisualDataSetField{}
			if needDelete(k) {
				continue
			}
			field.FieldId = kgo.KStr.Uniqid("px")
			field.Name = k
			field.DeName = k
			switch v.(type) {
			case int, uint, int8, uint8, int64, uint64, float32, float64:
				field.DeType = "2"
				field.GroupType = "q"
			case string:
				if _, err := time.Parse("2006-01-02", v.(string)); err != nil {
					field.DeType = "1"
					field.GroupType = "d"
				} else {
					field.DeType = "0"
					field.GroupType = "d"
				}
			}
			fields = append(fields, field)
		}
	}
	return fields
}

func needDelete(name string) bool {
	if strings.Contains(name, "id") || strings.Contains(name, "Id") {
		return true
	}
	return isExistInArray(name, ColumnTypeStr)
}

func isExistInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
