package main

import (
	"testing"
)

func BenchmarkLoadRestaurnats(b *testing.B) {
	testing.Init()
	for i := 0; i < b.N; i++ {
		loadRestaurants()
	}
}
