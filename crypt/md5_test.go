package crypt

import (
	"testing"
)

func TestGetMD5(t *testing.T) {
	md5 := GetMD5([]byte("hello md5!"))
	ok := "ef67ee85bbec92353b5cbd5c25b04263"
	if ok != md5 {
		t.Fatalf("md5 error, %s not equal %s", ok, md5)
	}
}

func TestGetMD5ByString(t *testing.T) {
	md5 := GetMD5ByString("hello md5!")
	ok := "ef67ee85bbec92353b5cbd5c25b04263"
	if ok != md5 {
		t.Fatalf("md5 error, %s not equal %s", ok, md5)
	}
}
