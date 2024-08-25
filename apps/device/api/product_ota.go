package api

// ==========================================================================
import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/pkg/tool"
	"path"
	"strings"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/apps/device/util"
	devicerpc "pandax/pkg/device_rpc"
)

type ProductOtaApi struct {
	ProductOtaApp services.ProductOtaModel
	DeviceApp     services.DeviceModel
}

const filePath = "uploads/file"

// GetProductOtaList Ota列表数据
func (p *ProductOtaApi) GetProductOtaList(rc *restfulx.ReqCtx) {
	data := entity.ProductOta{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")
	data.Pid = restfulx.QueryParam(rc, "pid")

	list, total, err := p.ProductOtaApp.FindListPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "查询OTA信息列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetProductOta 获取Ota
func (p *ProductOtaApi) GetProductOta(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	one, err := p.ProductOtaApp.FindOne(id)
	biz.ErrIsNil(err, "查询OTA信息失败")
	rc.ResData = one
}

// InsertProductOta 添加Ota
func (p *ProductOtaApi) InsertProductOta(rc *restfulx.ReqCtx) {
	var data entity.ProductOta
	restfulx.BindJsonAndValid(rc, &data)
	// 生成文件MD5值
	md5, err := tool.GetFileMd5(path.Join(filePath, data.Url))
	biz.ErrIsNil(err, "读取文件md5校验值错误")
	data.Id = utils.GenerateID("ota")
	data.Check = md5
	p.ProductOtaApp.Insert(data)
}

// UpdateProductOta 修改Ota
func (p *ProductOtaApi) UpdateProductOta(rc *restfulx.ReqCtx) {
	var data entity.ProductOta
	restfulx.BindJsonAndValid(rc, &data)
	md5, err := tool.GetFileMd5(path.Join(filePath, data.Url))
	biz.ErrIsNil(err, "读取文件md5校验值错误")
	data.Check = md5
	p.ProductOtaApp.Update(data)
}

// DeleteProductOta 删除Ota
func (p *ProductOtaApi) DeleteProductOta(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.ProductOtaApp.Delete(ids)
}

// DeleteProductOta OTA升级
func (p *ProductOtaApi) OtaDown(rc *restfulx.ReqCtx) {
	// 固件包
	id := restfulx.PathParam(rc, "id")
	pid := restfulx.QueryParam(rc, "pid")
	ota, err := p.ProductOtaApp.FindOne(id)
	biz.ErrIsNil(err, "查询OTA信息失败")
	// 1、对比所有设备与OTA固件版本
	devices, err := p.DeviceApp.FindList(entity.Device{Pid: pid, LinkStatus: "online"})
	biz.ErrIsNil(err, "该产品下没有设备存在")
	// 2、对低版本的设备进行指令下发升级
	go func() {
		rpc := devicerpc.RpcPayload{
			Method: "ota",
			Params: map[string]any{
				"verison": ota.Version,
				"url":     ota.Url,
				"Id":      ota.Id,
				"sign":    ota.Check,
			},
		}
		for _, device := range *devices {
			util.BuildRunDeviceRpc(device.Id, "", rpc)
		}

	}()

}
