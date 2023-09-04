package nodes

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

const (
	NODE_CONFIG_MESSAGE_TYPE_KEY    = "messageTypeKey"
	NODE_CONFIG_ORIGINATOR_TYPE_KEY = "originatorTypeKey"
)

// Metadata 前端 参数 Properties
type Metadata interface {
	Keys() []string
	With(key string, val interface{}) Metadata
	Value(key string) (interface{}, error)
	DecodePath(rawVal interface{}) error
}

type nodeMetadata struct {
	keypairs map[string]interface{}
}

func NewMetadata() Metadata {
	return &nodeMetadata{
		keypairs: make(map[string]interface{}),
	}
}

func NewMetadataWithString(vals string) Metadata {
	return &nodeMetadata{}
}

func NewMetadataWithValues(vals map[string]interface{}) Metadata {
	return &nodeMetadata{
		keypairs: vals,
	}
}

func (c *nodeMetadata) Keys() []string {
	keys := []string{}
	for key, _ := range c.keypairs {
		keys = append(keys, key)
	}
	return keys
}

func (c *nodeMetadata) Value(key string) (interface{}, error) {
	if val, found := c.keypairs[key]; found {
		return val, nil
	}
	return nil, fmt.Errorf("key '%s' not found", key)
}

func (c *nodeMetadata) With(key string, val interface{}) Metadata {
	c.keypairs[key] = val
	return c
}

func (c *nodeMetadata) DecodePath(rawVal interface{}) error {
	//return utils.Map2Struct(c.keypairs, rawVal)
	return mapstructure.Decode(c.keypairs, rawVal)
}
