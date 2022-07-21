package stringx

import (
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	test := []int{1, 2, 3, 4, 5}
	str := Join(test, ",")
	expected := "1,2,3,4,5"
	if str != expected {
		t.Fatal("data join error")
	}
}

func BenchmarkJoin(b *testing.B) {
	b.ReportAllocs()
	test := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	b.Run("base strings join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Join(test, ",")
		}
	})
	b.Run("for slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var str string
			for i := 0; i < len(test); i++ {
				str += fmt.Sprintf("%v", test[i])
				if i < len(test)-1 {
					str += fmt.Sprintf("%v", ",")
				}
			}
		}
	})

}
