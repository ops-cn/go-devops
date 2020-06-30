package util

import (
	"fmt"
	"reflect"

	"github.com/ops-cn/go-devops/common/thirdparty/structs"
)

// StructMapToStruct 结构体映射
func StructMapToStruct(s, ts interface{}) error {

	if !structs.IsStruct(s) || !structs.IsStruct(ts) {
		fmt.Println(structs.IsStruct(s))
		fmt.Println(structs.IsStruct(ts))
		return nil
	}

	ss, tss := structs.New(s), structs.New(ts)
	for _, tfield := range tss.Fields() {
		fmt.Println("===============" + tfield.Name())
		if !tfield.IsExported() {
			continue
		}

		if tfield.IsEmbedded() && tfield.Kind() == reflect.Struct {
			for _, tefield := range tfield.Fields() {
				if f, ok := ss.FieldOk(tefield.Name()); ok {
					tefield.Set2(f.Value())
				}
			}
		} else if f, ok := ss.FieldOk(tfield.Name()); ok {
			tfield.Set2(f.Value())
		}
	}

	return nil
}
