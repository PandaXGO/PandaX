package utils

import (
	"fmt"
	"github.com/kakuilan/kgo"
	"github.com/xuri/excelize/v2"
	"reflect"
	"time"
)

func ExportExcel(head []string, datas [][]interface{}, filePath string) error {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &head)
	for i, data := range datas {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &data)
	}
	excel.SaveAs(filePath)
	return nil
}

func GetFileName(path, model string) string {
	return path + fmt.Sprintf("%s_%s.xlsx", model, kgo.KTime.Date("20060102150405", time.Now()))
}

func InterfaceToExcel(data interface{}, fileName string) {
	heads := make([]string, 0)
	datas := make([][]interface{}, 0)
	v := reflect.ValueOf(data)
	max := int64(v.Len())
	for i := 0; i < int(max); i++ {
		u := v.Index(i).Interface()
		var key = reflect.TypeOf(u)
		var val = reflect.ValueOf(u)
		num := key.NumField()
		if i == 0 {
			for i := 0; i < num; i++ {
				heads = append(heads, key.Field(i).Name)
			}
		}
		field := make([]interface{}, 0)
		for i := 0; i < num; i++ {
			field = append(field, val.Field(i).Interface())
		}
		datas = append(datas, field)
	}
	ExportExcel(heads, datas, fileName)
}
