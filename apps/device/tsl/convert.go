package tsl

import (
	"strconv"
	"time"
)

func (t DefineAttribute) ConvertAttributeValue(v interface{}) interface{} {
	if *t.Rw == "w" {
		return ""
	}
	if v == nil {
		return t.DefaultValue
	}
	return v
}

func (t ValueType) ConvertValue(v interface{}) interface{} {
	var transfer Transfer
	switch t.Type {
	case TypeInt:
		transfer = TInt(t.DefineBase)
	case TypeFloat:
		transfer = TFloat(t.DefineBase)
	case TypeString:
		transfer = TText(t.DefineBase)
	case TypeBool:
		transfer = TBoolean(t.DefineBase)
	case TypeDate:
		transfer = TDate(t.DefineBase)
	case TypeEnum:
		transfer = TEnum(t.DefineBase)
	default:
		return nil
	}
	return transfer.Convert(v)
}

type Transfer interface {
	Convert(interface{}) interface{}
}

type TInt DefineBase

func (tInt TInt) Convert(v interface{}) interface{} {
	number, ok := v.(int64)
	if !ok {
		floatNumber, floatNumberOk := v.(float64)
		if !floatNumberOk {
			return 0
		}
		number = int64(floatNumber)
	}
	if tInt.Min != nil && int64(*tInt.Min) > number {
		return *tInt.Min
	}
	if tInt.Max != nil && int64(*tInt.Max) > number {
		return *tInt.Max
	}
	return number
}

type TFloat DefineBase

func (tFloat TFloat) Convert(v interface{}) interface{} {
	number, ok := v.(float64)
	if !ok {
		return 0
	}
	if tFloat.Min != nil && *tFloat.Min > number {
		number = *tFloat.Min
	}
	if tFloat.Max != nil && *tFloat.Max > number {
		number = *tFloat.Max
	}
	defaultDecimal := 2
	if tFloat.Decimals != nil {
		defaultDecimal = *tFloat.Decimals
	}
	number64, _ := strconv.ParseFloat(strconv.FormatFloat(number, 'f', defaultDecimal, 64), 32)
	return number64
}

type TText DefineBase

func (tText TText) Convert(v interface{}) interface{} {
	text, ok := v.(string)
	if !ok {
		return ""
	}
	if tText.MaxLength != nil && *tText.MaxLength > 0 && len(text) > *tText.MaxLength {
		return text[:*tText.MaxLength-1]
	} else {
		return text
	}
}

type TBoolean DefineBase

func (tBoolean TBoolean) Convert(v interface{}) interface{} {
	b, ok := v.(bool)
	if !ok {
		return ""
	}
	if b {
		return tBoolean.DefineBool[1].Value
	} else {
		return tBoolean.DefineBool[0].Value
	}
	return ""
}

type TDate DefineBase

const layout = "2006-01-02 15:04:05"

func (TDate TDate) Convert(v interface{}) interface{} {
	str, ok := v.(string)
	if !ok {
		return time.Time{}
	}
	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}
	} else {
		return t
	}
}

type TEnum DefineBase

func (tEnum TEnum) Convert(v interface{}) interface{} {
	tE, ok := v.(string)
	if !ok {
		return ""
	}
	if tEnum.Enums == nil {
		return ""
	}
	for _, node := range tEnum.Enums {
		if node.Key == tE {
			return node.Value
		}
	}
	return ""
}
