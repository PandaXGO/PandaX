package api

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/pkg/ys"
)

type YsApi struct {
	Ys *ys.Ys
}

func (j *YsApi) GetDeviceList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	devices, total, err := j.Ys.GetDeviceList(pageNum, pageSize)
	biz.ErrIsNil(err, "设备列表获取失败,可能萤石Token过期，请联系管理员。。")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     devices,
	}
}

func (j *YsApi) GetDeviceChannelList(rc *restfulx.ReqCtx) {
	deviceSerial := restfulx.PathParam(rc, "deviceSerial")
	cameras, err := j.Ys.GetDeviceChannelList(deviceSerial)
	biz.ErrIsNilAppendErr(err, "设备通道列表获取失败")
	rc.ResData = cameras
}

func (j *YsApi) GetDeviceLiveAddress(rc *restfulx.ReqCtx) {
	deviceSerial := restfulx.PathParam(rc, "deviceSerial")
	channelNo := restfulx.QueryInt(rc, "channelNo", 1)
	live, err := j.Ys.GetDeviceLiveAddress(deviceSerial, channelNo)
	biz.ErrIsNilAppendErr(err, "设备直播地址获取失败")
	rc.ResData = live
}

func (j *YsApi) StartPtz(rc *restfulx.ReqCtx) {
	deviceSerial := restfulx.PathParam(rc, "deviceSerial")
	channelNo := restfulx.QueryInt(rc, "channelNo", 1)
	direction := restfulx.QueryInt(rc, "direction", 0)
	speed := restfulx.QueryInt(rc, "speed", 0)
	err := j.Ys.StartPtz(deviceSerial, channelNo, direction, speed)
	biz.ErrIsNilAppendErr(err, "操作摄像头失败")
}

func (j *YsApi) StopPtz(rc *restfulx.ReqCtx) {
	deviceSerial := restfulx.PathParam(rc, "deviceSerial")
	channelNo := restfulx.QueryInt(rc, "channelNo", 1)
	direction := restfulx.QueryInt(rc, "direction", 0)
	err := j.Ys.StopPtz(deviceSerial, channelNo, direction)
	biz.ErrIsNilAppendErr(err, "停止摄像头失败")
}
