package flow

import "pandax/kit/utils"

type Edge struct {
	Id           string     `json:"id"`
	Type         string     `json:"type"`
	SourceNodeId string     `json:"sourceNodeId"` //当前节点
	TargetNodeId string     `json:"targetNodeId"` //下一节点
	StartPoint   Point      `json:"startPoint"`
	EndPoint     Point      `json:"endPoint"`
	Properties   Properties `json:"properties"`
	PointsList   []Point    `json:"pointsList"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (edge *Edge) getTargetNode(sourceNodeId string) string {
	if edge.SourceNodeId == sourceNodeId {
		return edge.TargetNodeId
	}
	return ""
}

func (edge *Edge) GetProperties(data any) error {
	if err := utils.Map2Struct(edge.Properties, data); err != nil {
		return err
	}
	return nil
}
