package nodes

import (
	"encoding/json"
	"github.com/PandaXGO/PandaKit/httpclient"
	"pandax/pkg/rule_engine/message"
)

type externalWechatNode struct {
	bareNode
	WebHook   string   `json:"webHook" yaml:"webHook"`
	MsgType   string   `json:"msgType" yaml:"msgType"`
	Content   string   `json:"content" yaml:"content"`
	IsAtAll   bool     `json:"isAtAll" yaml:"isAtAll"`
	AtMobiles []string `json:"atMobiles" yaml:"atMobiles"`
}

type externalWechatNodeFactory struct{}

func (f externalWechatNodeFactory) Name() string     { return "WechatNode" }
func (f externalWechatNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalWechatNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalWechatNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &externalWechatNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalWechatNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	template, err := ParseTemplate(n.Content, msg.GetAllMap())

	sendData := map[string]interface{}{
		"msgtype": "text",
		"text":    map[string]interface{}{"content": template},
	}
	if n.IsAtAll {
		sendData["text"].(map[string]interface{})["mentioned_mobile_list"] = []string{"@all"}
	} else {
		sendData["text"].(map[string]interface{})["mentioned_mobile_list"] = n.AtMobiles
	}
	marshal, _ := json.Marshal(sendData)
	postJson := httpclient.NewRequest(n.WebHook).Header("Content-Type", "application/json").PostJson(string(marshal))
	if postJson.StatusCode != 200 {
		n.Debug(msg, message.DEBUGOUT, "请求微信机器人hook接口失败")
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return err
		}
	}
	if successLabelNode != nil {
		n.Debug(msg, message.DEBUGOUT, "")
		return successLabelNode.Handle(msg)
	}
	return nil
}
