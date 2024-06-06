package api

import (
	"fmt"
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
	"pandax/kit/biz"
	"pandax/kit/model"
	"pandax/kit/restfulx"
	"pandax/kit/utils"
	"pandax/pkg/global"
)

type DictApi struct {
	DictType services.SysDictTypeModel
	DictData services.SysDictDataModel
}

func (p *DictApi) GetDictTypeList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	status := restfulx.QueryParam(rc, "status")
	dictName := restfulx.QueryParam(rc, "dictName")
	dictType := restfulx.QueryParam(rc, "dictType")

	list, total, err := p.DictType.FindListPage(pageNum, pageSize, entity.SysDictType{Status: status, DictName: dictName, DictType: dictType})
	biz.ErrIsNil(err, "查询字典类型分页失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (p *DictApi) GetDictType(rc *restfulx.ReqCtx) {
	dictId := restfulx.PathParamInt(rc, "dictId")
	data, err := p.DictType.FindOne(int64(dictId))
	biz.ErrIsNil(err, "查询字典类型失败")
	rc.ResData = data
}

func (p *DictApi) InsertDictType(rc *restfulx.ReqCtx) {
	var dict entity.SysDictType
	restfulx.BindJsonAndValid(rc, &dict)

	dict.CreateBy = rc.LoginAccount.UserName
	_, err := p.DictType.Insert(dict)
	biz.ErrIsNil(err, "添加字典类型失败")
}

func (p *DictApi) UpdateDictType(rc *restfulx.ReqCtx) {
	var dict entity.SysDictType
	restfulx.BindJsonAndValid(rc, &dict)

	dict.CreateBy = rc.LoginAccount.UserName
	err := p.DictType.Update(dict)
	biz.ErrIsNil(err, "修改字典类型失败")
}

func (p *DictApi) DeleteDictType(rc *restfulx.ReqCtx) {
	dictId := restfulx.PathParam(rc, "dictId")
	dictIds := utils.IdsStrToIdsIntGroup(dictId)

	deList := make([]int64, 0)
	for _, id := range dictIds {
		one, err := p.DictType.FindOne(id)
		if err != nil {
			continue
		}
		list, err := p.DictData.FindList(entity.SysDictData{DictType: one.DictType})
		if err != nil {
			continue
		}
		if len(*list) == 0 {
			deList = append(deList, id)
		} else {
			global.Log.Info(fmt.Sprintf("dictId: %d 存在字典数据绑定无法删除", id))
		}
	}
	err := p.DictType.Delete(deList)
	biz.ErrIsNil(err, "删除字典类型失败")
}

func (p *DictApi) ExportDictType(rc *restfulx.ReqCtx) {
	filename := restfulx.QueryParam(rc, "filename")
	status := restfulx.QueryParam(rc, "status")
	dictName := restfulx.QueryParam(rc, "dictName")
	dictType := restfulx.QueryParam(rc, "dictType")

	list, err := p.DictType.FindList(entity.SysDictType{Status: status, DictName: dictName, DictType: dictType})
	biz.ErrIsNil(err, "查询字典类型列表失败")
	fileName := utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, fileName)
	rc.Download(fileName)
}

func (p *DictApi) GetDictDataList(rc *restfulx.ReqCtx) {
	dictLabel := restfulx.QueryParam(rc, "dictLabel")
	dictType := restfulx.QueryParam(rc, "dictType")
	status := restfulx.QueryParam(rc, "status")
	data, err := p.DictData.FindList(entity.SysDictData{Status: status, DictType: dictType, DictLabel: dictLabel})
	biz.ErrIsNil(err, "查询字典列表失败")
	rc.ResData = data
}

func (p *DictApi) GetDictDataListByDictType(rc *restfulx.ReqCtx) {
	dictType := restfulx.QueryParam(rc, "dictType")
	biz.IsTrue(dictType != "", "请传入字典类型")
	data, err := p.DictData.FindList(entity.SysDictData{DictType: dictType})
	biz.ErrIsNil(err, "查询字典列表失败")
	rc.ResData = data
}

func (p *DictApi) GetDictData(rc *restfulx.ReqCtx) {
	dictCode := restfulx.PathParamInt(rc, "dictCode")
	data, err := p.DictData.FindOne(int64(dictCode))
	biz.ErrIsNil(err, "查询字典失败")
	rc.ResData = data
}

func (p *DictApi) InsertDictData(rc *restfulx.ReqCtx) {
	var data entity.SysDictData
	restfulx.BindJsonAndValid(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName
	_, err := p.DictData.Insert(data)
	biz.ErrIsNil(err, "添加字典失败")
}

func (p *DictApi) UpdateDictData(rc *restfulx.ReqCtx) {
	var data entity.SysDictData
	restfulx.BindJsonAndValid(rc, &data)

	data.CreateBy = rc.LoginAccount.UserName
	err := p.DictData.Update(data)
	biz.ErrIsNil(err, "修改字典失败")
}

func (p *DictApi) DeleteDictData(rc *restfulx.ReqCtx) {
	dictCode := restfulx.PathParam(rc, "dictCode")
	err := p.DictData.Delete(utils.IdsStrToIdsIntGroup(dictCode))
	biz.ErrIsNil(err, "修改字典失败")
}
