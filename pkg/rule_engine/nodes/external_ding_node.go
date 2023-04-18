package nodes

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/XM-GO/PandaKit/httpclient"
	"github.com/sirupsen/logrus"
	"net/url"
	"pandax/pkg/rule_engine/message"
	"time"
)

type externalDingNode struct {
	bareNode
	WebHook   string   `json:"webHook" yaml:"webHook"`
	Secret    string   `json:"secret"`
	MsgType   string   `json:"msgType" yaml:"msgType"`
	Content   string   `json:"content" yaml:"content"`
	IsAtAll   bool     `json:"isAtAll" yaml:"isAtAll"`
	AtMobiles []string `json:"atMobiles" yaml:"atMobiles"`
}

type externalDingNodeFactory struct{}

func (f externalDingNodeFactory) Name() string     { return "DingNode" }
func (f externalDingNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalDingNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalDingNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalDingNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalDingNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	//获取消息
	template, err := ParseTemplate(n.Content, msg.GetAllMap())
	if err != nil {
		return err
	}
	sendData := map[string]interface{}{
		"msgtype": "text",
		"text":    map[string]string{"content": template},
	}
	if n.IsAtAll {
		sendData["at"] = map[string]interface{}{"isAtAll": n.IsAtAll}
	} else {
		sendData["at"] = map[string]interface{}{"atMobiles": n.AtMobiles}
	}
	marshal, _ := json.Marshal(sendData)

	timestamp := time.Now().UnixMilli()
	sign := getSign(timestamp, n.Secret)

	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", n.WebHook, timestamp, sign)

	postJson := httpclient.NewRequest(url).Header("Content-Type", "application/json").PostJson(string(marshal))
	if postJson.StatusCode != 200 {
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return err
		}
	}
	if successLabelNode != nil {
		return successLabelNode.Handle(msg)
	}
	return nil
}

func getSign(timestamp int64, secret string) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(stringToSign))
	signData := hash.Sum(nil)
	return url.QueryEscape(base64.StdEncoding.EncodeToString(signData))
}
