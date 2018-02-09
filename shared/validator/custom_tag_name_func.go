package validator

import (
	"reflect"
)

// FormTagName for register new name field by tag
func FormTagName(fld reflect.StructField) string {
	name := fld.Tag.Get("form")
	if name == "" {
		return fld.Name
	}
	return name
}
