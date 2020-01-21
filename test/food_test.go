package tests

import (
	"fmt"
	"testing"

	"github.com/ken343/project-0/cmd/food/localpkg/place"
)

func TestProduceKey(t *testing.T) {
	myActualKey := place.ProduceKey("./test_key.txt")
	if myActualKey != "testkey123" {
		fmt.Println(myActualKey)
		t.Fail()
	}

}
