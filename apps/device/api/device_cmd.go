package api

// ==========================================================================
import (
	"context"
	"encoding/json"
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	ruleEntity "pandax/apps/rule/entity"
	ruleService "pandax/apps/rule/services"
	"pandax/pkg/global"
	"pandax/pkg/global_model"
	"pandax/pkg/rule_engine"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/tool"
	"strings"
	"time"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
)

type DeviceCmdLogApi struct {
	DeviceCmdLogApp services.DeviceCmdLogModel
	DeviceApp       services.DeviceModel
}

// GetDeviceCmdLogList 告警列表数据
func (p *DeviceCmdLogApi) GetDeviceCmdLogList(rc *restfulx.ReqCtx) {
	data := entity.DeviceCmdLog{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.DeviceId = restfulx.QueryParam(rc, "deviceId")
	data.State = restfulx.QueryParam(rc, "state")
	data.Type = restfulx.QueryParam(rc, "type")

	list, total := p.DeviceCmdLogApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// InsertDeviceCmdLog 添加DeviceCmdLog
func (p *DeviceCmdLogApi) InsertDeviceCmdLog(rc *restfulx.ReqCtx) {
	var data entity.DeviceCmdLog
	restfulx.BindJsonAndValid(rc, &data)
	//验证指令格式
	ms := make(map[string]interface{})
	err := json.Unmarshal([]byte(data.CmdContent), &ms)
	biz.ErrIsNil(err, "指令格式不正确")

	data.Id = global_model.GenerateID()
	data.State = "2"
	data.RequestTime = time.Now().Format("2006-01-02 15:04:05")
	one := p.DeviceApp.FindOne(data.DeviceId)
	biz.IsTrue(one.LinkStatus == global.ONLINE, "设备不在线无法下发指令")
	// 查询规则链
	findOne := ruleService.RuleChainModelDao.FindOne(one.Product.RuleChainId)
	ruleData := ruleEntity.RuleDataJson{}
	err = tool.StringToStruct(findOne.RuleDataJson, &ruleData)
	biz.ErrIsNil(err, "规则链数据转化失败")
	dataCode := ruleData.LfData.DataCode
	code, err := json.Marshal(dataCode)
	//新建规则链实体
	instance, errs := rule_engine.NewRuleChainInstance(code)
	if len(errs) > 0 {
		global.Log.Error("规则链初始化失败", errs[0])
		return
	}
	go func() {
		// 构建规则链消息
		metadataVals := map[string]interface{}{
			"deviceId":       data.DeviceId,
			"mode":           data.Mode,
			"deviceName":     one.Name,
			"deviceType":     one.DeviceType,
			"deviceProtocol": one.Product.ProtocolName,
			"productId":      one.Pid,
			"orgId":          one.OrgId,
			"owner":          one.Owner,
		}
		msg := message.NewMessage(one.Owner, message.RpcRequestToDevice, map[string]interface{}{
			"method": data.CmdName,
			"params": ms,
		}, metadataVals)
		err = instance.StartRuleChain(context.Background(), msg)
		if err != nil {
			global.Log.Error("规则链执行失败", errs)
			data.State = "1"
		} else {
			data.State = "0"
		}
		data.ResponseTime = time.Now().Format("2006-01-02 15:04:05.000")
		err = p.DeviceCmdLogApp.Insert(data)
		biz.ErrIsNil(err, "添加指令记录失败")
	}()

}

// DeleteDeviceCmdLog 删除告警
func (p *DeviceCmdLogApi) DeleteDeviceCmdLog(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.DeviceCmdLogApp.Delete(ids)
}
