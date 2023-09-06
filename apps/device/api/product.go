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
	"github.com/kakuilan/kgo"
	"pandax/pkg/global"
	"strings"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
)

type ProductApi struct {
	ProductApp  services.ProductModel
	DeviceApp   services.DeviceModel
	TemplateApp services.ProductTemplateModel
	OtaAPP      services.ProductOtaModel
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

	list, total := p.ProductApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
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

	rc.ResData = p.ProductApp.FindList(data)
}

// GetProduct 获取Product
func (p *ProductApi) GetProduct(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.ProductApp.FindOne(id)
}

// InsertProduct 添加Product
func (p *ProductApi) InsertProduct(rc *restfulx.ReqCtx) {
	var data entity.Product
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = kgo.KStr.Uniqid("p_")
	data.Owner = rc.LoginAccount.UserName
	p.ProductApp.Insert(data)
	// 创建taos数据库超级表
	createDeviceStable(data.Id)
}

// UpdateProduct 修改Product
func (p *ProductApi) UpdateProduct(rc *restfulx.ReqCtx) {
	var data entity.Product
	restfulx.BindJsonAndValid(rc, &data)

	p.ProductApp.Update(data)
}

// DeleteProduct 删除Product
func (p *ProductApi) DeleteProduct(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	for _, id := range ids {
		list := p.DeviceApp.FindList(entity.Device{Pid: id})
		biz.IsTrue(!(list != nil && len(*list) > 0), fmt.Sprintf("产品ID%s下存在设备，不能被删除", id))
	}
	// 删除产品
	p.ProductApp.Delete(ids)
	for _, id := range ids {
		// 删除所有模型，固件
		p.TemplateApp.Delete([]string{id})
		p.OtaAPP.Delete([]string{id})
		// 删除超级表
		deleteDeviceStable(id)
	}
}

func createDeviceStable(productId string) {
	err := global.TdDb.CreateStable(productId + "_" + entity.ATTRIBUTES_TSL)
	biz.ErrIsNil(err, "创建时序属性超级表失败")
	err = global.TdDb.CreateStable(productId + "_" + entity.TELEMETRY_TSL)
	biz.ErrIsNil(err, "创建时序遥测超级表失败")
}

func deleteDeviceStable(productId string) {
	err := global.TdDb.DropStable(productId + "_" + entity.ATTRIBUTES_TSL)
	biz.ErrIsNil(err, "删除时序属性超级表失败")
	err = global.TdDb.DropStable(productId + "_" + entity.TELEMETRY_TSL)
	biz.ErrIsNil(err, "删除时序遥测超级表失败")
}
