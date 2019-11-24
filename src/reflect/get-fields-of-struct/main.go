package main

import (
	"fmt"
	"reflect"
)

type Row struct {
	ID     string
	Name   string
	status string
}

func main() {
	row := Row{"1", "go", "enabled"}
	publicFields := structKeys(&row, false)
	fmt.Printf("publicFields=%v, count=%v\n", publicFields, len(publicFields))

	allFields := structKeys(&row, true)
	fmt.Printf("allFields=%v, count=%v\n", allFields, len(allFields))

	values := structValues(&row)
	fmt.Printf("values=%v\n", values)

}

func structKeys(s interface{}, withPrivate bool) []string {
	value := reflect.ValueOf(s).Elem()
	//fmt.Printf("value.NumField()=%v\n", value.NumField())
	var fields []string
	for i := 0; i < value.NumField(); i++ {
		name := value.Type().Field(i).Name
		if withPrivate {
			fields = append(fields, name)
		} else {
			f := value.FieldByName(name)
			if f.CanSet() {
				fields = append(fields, name)
			}
		}
	}
	return fields
}

func structValues(s interface{}) []interface{} {
	val := reflect.ValueOf(s).Elem()
	var values []interface{}
	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Name
		f := val.FieldByName(name)
		fmt.Printf("%v can set = %v\n", name, f.CanSet())
		if f.CanSet() {
			values = append(values, val.Field(i).CanInterface())
		}
	}
	return values
}
