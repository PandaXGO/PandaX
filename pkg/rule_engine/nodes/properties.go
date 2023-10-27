package nodes

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

const (
	NODE_CONFIG_MESSAGE_TYPE_KEY    = "messageTypeKey"
	NODE_CONFIG_ORIGINATOR_TYPE_KEY = "originatorTypeKey"
)

// Properties 前端 参数 Properties
type Properties interface {
	Keys() []string
	With(key string, val interface{}) Properties
	Value(key string) (interface{}, error)
	DecodePath(rawVal interface{}) error
}

type nodeProperties struct {
	keypairs map[string]interface{}
}

func NewProperties() Properties {
	return &nodeProperties{
		keypairs: make(map[string]interface{}),
	}
}

func NewPropertiesWithString(vals string) Properties {
	return &nodeProperties{}
}

func NewPropertiesWithValues(vals map[string]interface{}) Properties {
	return &nodeProperties{
		keypairs: vals,
	}
}

func (c *nodeProperties) Keys() []string {
	keys := []string{}
	for key, _ := range c.keypairs {
		keys = append(keys, key)
	}
	return keys
}

func (c *nodeProperties) Value(key string) (interface{}, error) {
	if val, found := c.keypairs[key]; found {
		return val, nil
	}
	return nil, fmt.Errorf("key '%s' not found", key)
}

func (c *nodeProperties) With(key string, val interface{}) Properties {
	c.keypairs[key] = val
	return c
}

func (c *nodeProperties) DecodePath(rawVal interface{}) error {
	//return utils.Map2Struct(c.keypairs, rawVal)
	return mapstructure.Decode(c.keypairs, rawVal)
}
