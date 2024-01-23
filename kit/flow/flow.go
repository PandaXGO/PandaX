package flow

type Flow struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

func (f *Flow) GetStartNode(ty string) *Node {
	for _, node := range f.Nodes {
		if node.IsStartNode(ty) {
			return &node
		}
	}
	return nil
}

func (f *Flow) GetTargetNodeId(sourceNodeId string) string {
	for _, edge := range f.Edges {
		return edge.getTargetNode(sourceNodeId)
	}
	return ""
}

func (f *Flow) GetTargetNode(sourceNodeId string) *Node {
	for _, edge := range f.Edges {
		if edge.SourceNodeId == sourceNodeId {
			return f.GetNode(edge.TargetNodeId)
		}
	}
	return nil
}

func (f *Flow) GetNode(nodeId string) *Node {
	for _, node := range f.Nodes {
		if node.Id == nodeId {
			return &node
		}
	}
	return nil
}
