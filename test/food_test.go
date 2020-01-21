package tests

import (
	"testing"

	"github.com/ken343/project-0/cmd/food/localpkg/place"
)

func TestProduceKey(t *testing.T) {
	myActualKey := place.ProduceKey("./test_key.go")
	if myActualKey != "testkey123" {
		t.Fail()
	}

}
