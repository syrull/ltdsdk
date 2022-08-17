package ltdsdk

import (
	"fmt"
	"reflect"
	"strconv"
)

func toQueryString(obj interface{}) map[string]string {
	elem := reflect.ValueOf(obj).Elem()
	queryStringMap := make(map[string]string)
	for i := 0; i < elem.NumField(); i++ {
		k := elem.Type().Field(i).Tag.Get("qs")
		v := elem.Field(i).Interface()
		switch v := v.(type) {
		case int:
			if v != 0 {
				queryStringMap[k] = strconv.Itoa(v)
			}
		case string:
			if v != "" {
				queryStringMap[k] = v
			}
		case float32:
			if v != 0 {
				queryStringMap[k] = fmt.Sprintf("%g", v)
			}
		case float64:
			if v != 0 {
				queryStringMap[k] = fmt.Sprintf("%g", v)
			}
		}
	}
	return queryStringMap
}
