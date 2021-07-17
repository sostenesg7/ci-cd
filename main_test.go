package main

import (
	"testing"
)

func Test_ResponseOk(t *testing.T) {
	response := ResponseOk()
	if response != "OK" {
		t.Fatal("ResponseOk must return \"OK\"")
	}
}
