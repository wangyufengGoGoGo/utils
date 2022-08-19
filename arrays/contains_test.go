package arrays

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("字符串", func(t *testing.T) {
		test := []string{"aaa", "bbb", "ccc"}
		val := "ccc"
		if ContainsString(test, val) {
			fmt.Println("contain")
		}
	})
}
