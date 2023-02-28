package manifest

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type Node struct {
	Type       string                 `json:"type" yaml:"type"`
	Id         string                 `json:"id" yaml:"id"`
	Properties map[string]interface{} `json:"properties" yaml:"properties"` //debugMode
}

type Edge struct {
	SourceNodeId string `json:"sourceNodeId" yaml:"sourceNodeId"`
	TargetNodeId string `json:"targetNodeId" yaml:"targetNodeId"`
	Type         string `json:"type" yaml:"type"` //success or fail
}

type Manifest struct {
	FirstRuleNodeId string `json:"firstRuleNodeId" yaml:"firstRuleNodeId"`
	Nodes           []Node `json:"nodes" yaml:"nodes"`
	Edges           []Edge `json:"edges" yaml:"edges"`
}

func New(data []byte) (*Manifest, error) {
	firstRuleNodeId := ""
	manifest := make(map[string]interface{})
	if err := json.Unmarshal(data, manifest); err != nil {
		logrus.WithError(err).Errorf("invalid node chain manifest file")
		return nil, err
	}
	nodes := make([]Node, 0)
	for _, node := range manifest["nodes"].([]map[string]interface{}) {
		if node["type"].(string) == "InputNode" {
			firstRuleNodeId = node["id"].(string)
		}
		nodes = append(nodes, Node{
			Id:         node["id"].(string),
			Type:       node["type"].(string),
			Properties: node["properties"].(map[string]interface{}),
		})
	}
	edges := make([]Edge, 0)
	for _, edge := range manifest["edges"].([]map[string]interface{}) {
		edges = append(edges, Edge{
			Type:         edge["type"].(string),
			SourceNodeId: edge["sourceNodeId"].(string),
			TargetNodeId: edge["targetNodeId"].(string),
		})
	}
	m := &Manifest{
		FirstRuleNodeId: firstRuleNodeId,
		Nodes:           nodes,
		Edges:           edges,
	}
	return m, nil
}
