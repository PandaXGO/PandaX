package api

// ==========================================================================
// 生成日期：2023-06-30 08:57:48 +0800 CST
// 生成路径: apps/device/api/products.go
// 生成人：panda
// ==========================================================================
import (
	"fmt"
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/pkg/global"
	"strings"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	ruleService "pandax/apps/rule/services"
)

type ProductApi struct {
	ProductApp  services.ProductModel
	DeviceApp   services.DeviceModel
	TemplateApp services.ProductTemplateModel
	OtaAPP      services.ProductOtaModel
	RuleApp     ruleService.RuleChainModel
}

// GetProductList Product列表数据
func (p *ProductApi) GetProductList(rc *restfulx.ReqCtx) {
	data := entity.Product{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Status = restfulx.QueryParam(rc, "status")
	data.ProductCategoryId = restfulx.QueryParam(rc, "productCategoryId")
	data.ProtocolName = restfulx.QueryParam(rc, "protocolName")
	data.DeviceType = restfulx.QueryParam(rc, "deviceType")
	data.Name = restfulx.QueryParam(rc, "name")

	list, total, err := p.ProductApp.FindListPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "查询产品列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (p *ProductApi) GetProductListAll(rc *restfulx.ReqCtx) {
	data := entity.Product{}
	data.Status = restfulx.QueryParam(rc, "status")
	data.ProductCategoryId = restfulx.QueryParam(rc, "productCategoryId")
	data.ProtocolName = restfulx.QueryParam(rc, "protocolName")
	data.DeviceType = restfulx.QueryParam(rc, "deviceType")
	data.Name = restfulx.QueryParam(rc, "name")
	list, err := p.ProductApp.FindList(data)
	biz.ErrIsNil(err, "查询产品列表失败")
	rc.ResData = list
}

// GetProductTsl Template列表数据
func (p *ProductApi) GetProductTsl(rc *restfulx.ReqCtx) {
	data := entity.ProductTemplate{}
	data.Pid = restfulx.PathParam(rc, "id")
	templates, err := p.TemplateApp.FindList(data)
	biz.ErrIsNil(err, "查询产品模板列表失败")
	attributes := make([]map[string]interface{}, 0)
	telemetry := make([]map[string]interface{}, 0)
	commands := make([]map[string]interface{}, 0)
	events := make([]map[string]interface{}, 0)
	for _, template := range *templates {
		tslData := map[string]interface{}{
			"name":       template.Name,
			"identifier": template.Key,
			"type":       template.Type,
			"define":     template.Define,
		}
		if template.Classify == global.TslAttributesType {
			attributes = append(attributes, tslData)
		}
		if template.Classify == global.TslTelemetryType {
			telemetry = append(telemetry, tslData)
		}
		if template.Classify == global.TslCommandsType {
			commands = append(commands, tslData)
		}
		if template.Classify == global.TslEventType {
			events = append(events, tslData)
		}

	}

	rc.ResData = map[string]interface{}{
		"attributes": attributes,
		"telemetry":  telemetry,
		"commands":   commands,
		"events":     events,
	}
}

// GetProduct 获取Product
func (p *ProductApi) GetProduct(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	one, err := p.ProductApp.FindOne(id)
	biz.ErrIsNil(err, "查询失败")
	rc.ResData = one
}

// InsertProduct 添加Product
func (p *ProductApi) InsertProduct(rc *restfulx.ReqCtx) {
	var data entity.Product
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = utils.GenerateID("p")
	data.Owner = rc.LoginAccount.UserName
	data.OrgId = rc.LoginAccount.OrganizationId
	// 如果未设置规则链，默认为主链
	if data.RuleChainId == "" {
		root, err := p.RuleApp.FindOneByRoot()
		biz.ErrIsNil(err, "规则链查询错误")
		data.RuleChainId = root.Id
	}

	_, err := p.ProductApp.Insert(data)
	biz.ErrIsNilAppendErr(err, "")
}

// UpdateProduct 修改Product
func (p *ProductApi) UpdateProduct(rc *restfulx.ReqCtx) {
	var data entity.Product
	restfulx.BindQuery(rc, &data)

	_, err := p.ProductApp.Update(data)
	biz.ErrIsNilAppendErr(err, "")
}

func (p *ProductApi) UpdateProductStatue(rc *restfulx.ReqCtx) {
	var data entity.Product
	restfulx.BindJsonAndValid(rc, &data)

	_, err := p.ProductApp.Update(data)
	biz.ErrIsNilAppendErr(err, "")
}

// DeleteProduct 删除Product
func (p *ProductApi) DeleteProduct(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	for _, id := range ids {
		list, _ := p.DeviceApp.FindList(entity.Device{Pid: id})
		biz.IsTrue(!(list != nil && len(*list) > 0), fmt.Sprintf("产品ID%s下存在设备，不能被删除", id))
	}
	// 删除产品
	err := p.ProductApp.Delete(ids)
	biz.ErrIsNil(err, "产品删除失败")
}
