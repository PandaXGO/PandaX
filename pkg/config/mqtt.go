package config

type Mqtt struct {
	Broker     string `mapstructure:"broker" json:"broker" yaml:"broker"`
	HttpBroker string `mapstructure:"httpBroker" json:"httpBroker" yaml:"httpBroker"`
	Qos        int    `mapstructure:"qos" json:"qos" yaml:"qos"`
	Username   string `mapstructure:"username" json:"username" yaml:"username"`
	Password   string `mapstructure:"password" json:"password" yaml:"password"`
}
