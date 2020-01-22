package test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ken343/project-0/cmd/food/localpkg/place"
)

func TestProduceKey(t *testing.T) {
	myActualKey := place.ProduceKey("./testdata/test_key.txt")
	if myActualKey != "testkey123" {
		fmt.Println(myActualKey)
		t.Errorf("ProduceKey(\"./test_key.txt\" failed, expected %v, got %v", "testkey123", myActualKey)
	} else {
		t.Logf("ProduceKey(\"./test_key.txt\") success, expected %v, got %v", "testkey123", myActualKey)
	}

}

/*
func TestLoadPlaces(t *testing.T) {
	f, err := os.Open("./testdata/testplaces.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	myActual := food.LoadPlaces(f)

	myDedoder := json.NewDecoder(f)
	var myExpected place.Places
	err = myDedoder.Decode(&myExpected)
	if err != nil {
		log.Fatal(err)
	}

	if myActual != myExpected {
		t.Errorf("LoadPlaces(f), failed, expected %v, got %v", myExpected, myActual)
	} else {
		t.Logf("LoadPlaces(f), success, expected %v, got %v", myExpected, myActual)
	}
}
*/

func BenchmarkPrintSuggestions(b *testing.B) {
	f, err := os.Open("./testdata/testplaces.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	myDedoder := json.NewDecoder(f)
	var myExpected place.Places
	err = myDedoder.Decode(&myExpected)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		place.PrintSuggestions(myExpected)
	}
}
