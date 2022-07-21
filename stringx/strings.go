package stringx

import (
	"fmt"
	"github.com/wangyufengx/utils/errorx"
	"reflect"
	"sort"
)

func JsonSpliceToString(data any, seq string) (string, error) {
	var (
		str    string
		tmpStr string
		err    error
	)
	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.Map:
		tmpData := data.(map[string]any)
		keys := MapKeys(value.MapKeys())
		sort.Strings(keys)
		for k, v := range keys {
			tmpStr, err = BaseDataTypeToString(tmpData[v])
			if err != nil {
				return "", err
			}
			str += fmt.Sprintf("%s=%v", v, tmpStr)
			if k < len(keys)-1 {
				str += fmt.Sprintf("%s", seq)
			}
		}
	}
	return str, nil
}

func BaseDataTypeToString(data any) (string, error) {
	var (
		str string
		err error
	)
	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Float32, reflect.Float64:
		str = fmt.Sprintf("%v", data)
	case reflect.String:
		str = fmt.Sprintf("%v", data)
	case reflect.Bool:
		str = fmt.Sprintf("%v", data)
	case reflect.Slice, reflect.Array:
		str, err = SliceToString(data)
		if err != nil {
			return "", err
		}
	default:
		return "", errorx.ErrDataType
	}
	return str, nil
}

func SliceToString(data any) (string, error) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return "", errorx.ErrDataType
	}

	var str string

	switch data.(type) {
	case []string:
		str = Join(data.([]string), ",")
	case []int:
		str = Join(data.([]int), ",")

	}

	return str, nil
}

func MapKeys(mapKeys []reflect.Value) []string {
	keys := make([]string, 0)
	for _, v := range mapKeys {
		keys = append(keys, v.String())
	}
	return keys
}
