package structcheck

import (
	"fmt"
	"reflect"
)

func Bool(v bool) *bool { return &v }

func Int8(v int8) *int8 { return &v }

func Int16(v int16) *int16 { return &v }

func Int32(v int32) *int32 { return &v }

func Int64(v int64) *int64 { return &v }

func Uint8(v uint8) *uint8 { return &v }

func Uint16(v uint16) *uint16 { return &v }

func Uint32(v uint32) *uint32 { return &v }

func Uint64(v uint64) *uint64 { return &v }

func Float32(v float32) *float32 { return &v }

func Float64(v float64) *float64 { return &v }

func String(v string) *string { return &v }

func IsExpected(expected interface{}, actual interface{}) (ok bool, msg []string) {
	ok, msg = isExpected(expected, actual, "struct")
	return
}

func isExpected(expected interface{}, actual interface{}, field string) (ok bool, msg []string) {
	ok = true
	msg = []string{}
	if ok = reflect.TypeOf(expected) == reflect.TypeOf(actual); !ok {
		msg = append(msg, fmt.Sprintf("%s type not same, expect %s but actual %s", field, reflect.TypeOf(expected).String(), reflect.TypeOf(actual).String()))
		return
	}
	exp := reflect.ValueOf(expected)
	act := reflect.ValueOf(actual)
	if exp.Kind() == reflect.Slice {
		// slice types
		// check capacity first
		if ok = act.Cap() >= exp.Cap(); !ok {
			msg = append(msg, fmt.Sprintf("%s slice size not enough, expect %d but actual %d", field, exp.Cap(), act.Cap()))
			return
		}
		// check if the expected slice is a sub-slice of actual slice
		for i := 0; i < exp.Cap(); i++ {
			o, m := isExpected(exp.Index(i).Interface(), act.Index(i).Interface(), fmt.Sprintf("%s.%d", field, i))
			ok = ok && o
			msg = append(msg, m...)
		}
	} else {
		// pointer types
		exp = exp.Elem()
		act = act.Elem()
		if exp.Kind() == reflect.Struct {
			// pointer to struct
			for i := 0; i < exp.NumField(); i++ {
				if !exp.Field(i).IsNil() {
					// if expected is nil, no need to check this field
					o, m := isExpected(exp.Field(i).Interface(), act.Field(i).Interface(), fmt.Sprintf("%s.%s", field, exp.Type().Field(i).Name))
					ok = ok && o
					msg = append(msg, m...)
				}
			}
		} else {
			// pointer to basic types
			if ok = reflect.DeepEqual(exp.Interface(), act.Interface()); !ok {
				msg = append(msg, fmt.Sprintf("%s not equal, expected %v but actual %v", field, exp.Interface(), act.Interface()))
				return
			}
		}
	}
	return
}
