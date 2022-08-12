package stringx

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSpliceString(t *testing.T) {
	m := map[string]any{
		"test":  []string{"aaa", "bbb"},
		"test1": 1,
		"test2": 1.1,
		"test3": true,
		"test5": "测试",
	}

	value := reflect.ValueOf(m)
	fmt.Println(value.Kind())
	typeValue := reflect.TypeOf(m)
	fmt.Println(typeValue.Key())
	fmt.Println(value.MapKeys())
	for _, v := range value.MapKeys() {
		fmt.Println(m[v.String()])
	}

	spliceString, err := JsonSpliceToString(m, "&", true)
	if err != nil {
		t.Fatal(err)
	}
	expected := "test=aaa%2Cbbb&test1=1&test2=1.1&test3=true&test5=%E6%B5%8B%E8%AF%95"
	if spliceString != expected {
		t.Fatal(fmt.Sprintf("预期{%s}, 实际{%s}", expected, spliceString))
	}
}

func TestBaseDataTypeToString(t *testing.T) {

	var (
		str string
		err error
	)
	t.Run("数字", func(t *testing.T) {
		num := 1
		str, err = BaseDataTypeToString(num)
		if err != nil {
			t.Fatal(err)
		}
		if str != "1" {
			t.Fatal("data convert error")
		}
	})
	t.Run("字符串", func(t *testing.T) {
		num := "test"
		str, err = BaseDataTypeToString(num)
		if err != nil {
			t.Fatal(err)
		}
		if str != "test" {
			t.Fatal("data convert error")
		}
	})
	t.Run("布尔", func(t *testing.T) {
		str, err = BaseDataTypeToString(true)
		if err != nil {
			t.Fatal(err)
		}
		if str != "true" {
			t.Fatal("data convert error")
		}
	})
	t.Run("切片(字符串)", func(t *testing.T) {
		data := []string{"aaa", "bbb"}
		str, err = BaseDataTypeToString(data)
		if err != nil {
			t.Fatal(err)
		}
		if str != "aaa,bbb" {
			t.Fatal("data convert error")
		}
	})
	t.Run("切片(数值)", func(t *testing.T) {
		data := []int{1, 2}
		str, err = BaseDataTypeToString(data)
		if err != nil {
			t.Fatal(err)
		}
		if str != "1,2" {
			t.Fatal("data convert error")
		}
	})
}
