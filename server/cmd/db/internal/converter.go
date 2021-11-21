package internal

import (
	"errors"
	"reflect"
)

// MakeArrFromStruct struct -> []interface
func MakeArrFromStruct(data interface{}) []interface{} {
	v := reflect.ValueOf(data)
	arr := []interface{}{}

	if v.Kind() == reflect.Ptr {
		return arr
	}

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsZero() {
			continue
		}
		arr = append(arr, v.Field(i).Interface())
	}
	return arr
}

// FillStructFromMap return pointer to filled struct by sample struct and map
func FillStructFromMap(sampleStruct interface{}, dataSrc map[string]interface{}) (interface{}, []error) {
	// copy to new
	typeStruct := reflect.TypeOf(sampleStruct)
	if reflect.ValueOf(sampleStruct).Kind() == reflect.Ptr {
		typeStruct = typeStruct.Elem() // if sampleStruct is pointer then remove pointer
	}
	dstStruct := reflect.Indirect(reflect.New(typeStruct))

	errs := []error{}
	for k, v := range dataSrc {
		dstFieldValue := dstStruct.FieldByName(k)

		if !dstFieldValue.IsValid() {
			errs = append(errs, errors.New("no such field: "+k+" in obj."))
			continue
		}
		if !dstFieldValue.CanSet() {
			errs = append(errs, errors.New("can't set "+k+" field value."))
			continue
		}

		val := reflect.ValueOf(v)
		if dstFieldValue.Type() != val.Type() {
			errs = append(errs, errors.New("provided value type didn't match "+k+" field type."))
			continue
		}

		dstFieldValue.Set(val)
	}

	return dstStruct.Addr().Interface(), errs
}
