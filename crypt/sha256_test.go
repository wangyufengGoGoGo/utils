package crypt

import (
	"fmt"
	"testing"
)

func TestGetSha256(t *testing.T) {
	data := "hello sha256"
	sha256Active := "433855b7d2b96c23a6f60e70c655eb4305e8806b682a9596a200642f947259b1"
	sha256 := GetSha256([]byte(data))
	if sha256 != sha256Active {
		t.Fatalf("sha256 error, %s not equal %s", sha256, sha256Active)
	}
	fmt.Println(sha256)
}
