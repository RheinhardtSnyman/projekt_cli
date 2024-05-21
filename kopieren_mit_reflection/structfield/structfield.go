// package structfield is used for struct fields
package structfield

import (
	"fmt"
	"reflect"
)

// copy is used for copying diffrent fields
func Copy(dst, src interface{}) error {
	typeDst := reflect.TypeOf(dst)
	if typeDst.Kind() != reflect.Ptr {
		return fmt.Errorf("dst is not a pointer")
	}
	valDst := reflect.ValueOf(dst).Elem() //Elem nur mit pointer
	valSrc := reflect.ValueOf(src)
	typeSrc := reflect.TypeOf(src)
	for i := 0; i < valSrc.NumField(); i++ {
		srcField := typeSrc.Field(i)
		dstField := valDst.FieldByName(srcField.Name)
		if !dstField.IsValid() {
			continue
		}
		srcTag := typeSrc.Field(i).Tag
		if srcTag.Get("structField") == "nocopy" {
			continue
		}
		dstField.Set(valSrc.Field(i))
	}
	return nil
}
