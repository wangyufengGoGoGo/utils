package enum

import (
	"fmt"
	"github.com/wangyufengx/utils/arrays"
	"testing"
)

func TestCompSurname(t *testing.T) {
	if arrays.ContainsString(CompSurname, "慕容") {
		fmt.Println("慕容")
	}
	fmt.Println("复姓如下:")
	fmt.Println(CompSurname)
}
