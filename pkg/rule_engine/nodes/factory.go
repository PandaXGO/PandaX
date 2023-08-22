package nodes

import (
	"fmt"
)

const (
	NODE_CATEGORY_FILTER     = "filter"
	NODE_CATEGORY_ACTION     = "action"
	NODE_CATEGORY_ENRICHMENT = "enrichment"
	NODE_CATEGORY_TRANSFORM  = "transform"
	NODE_CATEGORY_EXTERNAL   = "external"
	NODE_CATEGORY_OTHERS     = "others"
	NODE_CATEGORY_FLOWS      = "flows"
)

type Factory interface {
	Name() string
	Category() string
	Labels() []string
	Create(id string, meta Metadata) (Node, error)
}

var (
	// allNodeFactories hold all node's factory
	allNodeFactories map[string]Factory = make(map[string]Factory)

	// allNodeCategories hold node's metadata by category
	allNodeCategories map[string][]map[string]interface{} = make(map[string][]map[string]interface{})
	allCategories     []map[string]interface{}            = make([]map[string]interface{}, 0)
)

func RegisterFactory(f Factory) {
	allNodeFactories[f.Name()] = f

	if allNodeCategories[f.Category()] == nil {
		allNodeCategories[f.Category()] = []map[string]interface{}{}
	}
	allNodeCategories[f.Category()] = append(allNodeCategories[f.Category()], map[string]interface{}{"name": f.Name(), "labels": f.Labels()})
	allCategories = append(allCategories, map[string]interface{}{"name": f.Name(), "labels": f.Labels()})
}

// NewNode is the only way to create a new node
func NewNode(nodeType string, id string, meta Metadata) (Node, error) {
	if f, found := allNodeFactories[nodeType]; found {
		return f.Create(id, meta)
	}
	return nil, fmt.Errorf("invalid node type '%s'", nodeType)
}

// GetCategoryNodes return specified category's all nodes
func GetCategoryNodes() map[string][]map[string]interface{} { return allNodeCategories }

func GetCategory() []map[string]interface{} { return allCategories }
