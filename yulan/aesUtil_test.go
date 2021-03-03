package main

import (
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	crypted, err := AesEncrypt([]byte("abc"), []byte("321423u9y8d2fwf9"))
	if err != nil {
		t.Error(err)
	}
	println("crypted:", crypted)
}
