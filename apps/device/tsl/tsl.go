package tsl

const (
	TypeEnum   = "enum" // 枚举类型
	TypeInt    = "int64"
	TypeString = "string"
	TypeBool   = "bool"
	TypeFloat  = "float64"
	TypeDate   = "date"
	TypeStruct = "struct"
)

// DefineAttribute 属性拓展
type DefineAttribute struct {
	DefaultValue *string `json:"defaultValue"` // 属性时
	Rw           *string `json:"rw"`
}

// DefineBase 基础类型参数
type DefineBase struct {
	Max      *float64 `json:"max" `     // 最大,数字类型:int64、float64
	Min      *float64 `json:"min" `     // 最小,数字类型:int64、float64
	Step     *float64 `json:"step"`     // 小数位数,数字类型:int64、float64
	Decimals *int     `json:"decimals"` // 小数位数,数字类型:float64
	Unit     *string  `json:"unit"`     // 单位,数字类型: int64、float64

	MaxLength *int `json:"maxLength"` // 最大长度,字符类型:string

	DefineBool []DefineBool   `json:"boolDefine"`
	Enums      []DefineEnum   `json:"enumDefine"`    // 枚举类型:enum
	Struct     []DefineStruct `json:"structDefine" ` // 对象类型:Struct
}

type DefineBool struct {
	Key   string `json:"key"`   // 键 0、1
	Value string `json:"value"` // 枚举值
}

type DefineEnum struct {
	Key   string `json:"key"`   // 键 0、1
	Value string `json:"value"` // 枚举值
}

// DefineStruct 扩展类型参数:对象型
type DefineStruct struct {
	Key       string    `json:"key"`       //参数标识
	Name      string    `json:"name"`      // 参数名称
	ValueType ValueType `json:"valueType"` // 参数值
	Desc      string    `json:"desc"`      // 描述
}

type ValueType struct {
	Type       string `json:"type" dc:"数据类型" v:"required#请选择数据类型"` // 类型
	DefineBase        // 参数
}

// DefineCommands 命令
type DefineCommands struct {
	Key    string                `json:"key" dc:"功能标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入功能标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name   string                `json:"name" dc:"功能名称" v:"required#请输入功能名称"` // 功能名称
	Inputs []DefineCommandsInput `json:"inputs" dc:"输入参数"`                    // 输入参数
	Output ValueType             `json:"output" dc:"输出参数"`                    // 输出参数
	Desc   string                `json:"desc" dc:"描述"`                        // 描述
}

// DefineCommandsInput 命令:输入参数
type DefineCommandsInput struct {
	Key       string    `json:"key"`
	Name      string    `json:"name"`      // 输入参数名称
	ValueType ValueType `json:"valueType"` // 参数值
	Desc      string    `json:"desc"`      // 描述
}
