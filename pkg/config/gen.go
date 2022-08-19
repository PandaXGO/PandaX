package config

/**
 * @Description
 * @Author Panda
 * @Date 2021/12/31 15:13
 **/
type Gen struct {
	Dbname    string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Frontpath string `mapstructure:"frontpath" json:"frontpath" yaml:"frontpath"`
}
