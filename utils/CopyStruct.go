// Package utils
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-13 17:55
package utils

import (
	"reflect"
)

//
//// ergodicstructergodic
//func structByReflect(data map[string]interface{}, inStructPtr interface{}) {
//	rType := reflect.TypeOf(inStructPtr)
//	rVal := reflect.ValueOf(inStructPtr)
//	if rType.Kind() == reflect.Ptr {
//		// IncominginStructPtrIncoming，Incoming.Elem()Incomingvalue
//		rType = rType.Elem()
//		rVal = rVal.Elem()
//	} else {
//		panic("inStructPtr must be ptr to struct")
//	}
//	// Traversal structure
//	for i := 0; i < rType.NumField(); i++ {
//		t := rType.Field(i)
//		f := rVal.Field(i)
//		// obtaintagobtain
//		key := t.Tag.Get("key")
//		if v, ok := data[key]; ok {
//			// Check if type conversion is required
//			dataType := reflect.TypeOf(v)
//			structType := f.Type()
//			if structType == dataType {
//				f.Set(reflect.ValueOf(v))
//			} else {
//				if dataType.ConvertibleTo(structType) {
//					// Conversion Type
//					f.Set(reflect.ValueOf(v).Convert(structType))
//				} else {
//					panic(t.Name + " type mismatch")
//				}
//			}
//		} else {
//			panic(t.Name + " not found")
//		}
//	}
//}

// CopyFields
// usebuseause
// IffieldsIf, IfbIfaIf
// aShould be a structural pointer
func CopyFields(a interface{}, b interface{}, fields ...string) (err error) {
	at := reflect.TypeOf(a)
	av := reflect.ValueOf(a)
	bt := reflect.TypeOf(b)
	bv := reflect.ValueOf(b)

	// To make a simple judgment
	if at.Kind() != reflect.Ptr {
		return Errorf("a must be a struct pointer")
	}
	if bt.Kind() == reflect.Ptr {
		bt = bt.Elem()
		bv = bv.Elem()
	}
	av = reflect.ValueOf(av.Interface())
	// Which fields to copy
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bt.NumField(); i++ {
			_fields = append(_fields, bt.Field(i).Name)
		}
	}
	if len(_fields) == 0 {
		return Errorf("no fields to copy")
	}
	// copy
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := av.Elem().FieldByName(name)
		bValue := bv.FieldByName(name)
		// aCopy only when there is a field with the same name and the type is consistent in the
		if f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		}
	}
	return
}
