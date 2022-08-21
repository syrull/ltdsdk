package ltdsdk

import (
	"fmt"
	"reflect"
	"strconv"
)

// Transforms an object's fields with `qs` tag and (int|string|float|bool) type
// into map[string]string
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
		case bool:
			queryStringMap[k] = strconv.FormatBool(v)
		}
	}
	return queryStringMap
}

// Parses a string to Float32
func parseStringToFloat32(str string, valueOnError float64) float32 {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		f = valueOnError
	}
	return float32(f)
}

// Parses a string to Integer
func parseStringToInt(str string, valueOnError int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		i = valueOnError
	}
	return i
}
