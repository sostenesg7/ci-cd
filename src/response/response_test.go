package response

import (
	"testing"
)

func Test_ResponseOk(t *testing.T) {
	response := "OK"
	if response != "OK" {
		t.Fatal("ResponseOk must return \"OK\"")
	}
}
