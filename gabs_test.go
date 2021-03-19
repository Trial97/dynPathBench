package main

import (
	"fmt"
	"testing"

	"github.com/Jeffail/gabs/v2"
)

func TestGabs(t *testing.T) {
	// doesn't support directly setting []slices using set

	jsonObj := gabs.New()
	jsonObj.Set(10, "outter", "inner", "value")
	jsonObj.SetP(20, "outter.inner.value2")
	jsonObj.Set(30, "outter", "inner2", "value3")
	// if we want to set an array we use something similar
	// but we need to create them a leafs in the tree
	jsonObj.ArrayAppendP(40, "outter.inner3.value2")
	fmt.Println(jsonObj.String())
}

func BenchmarkGabsSet(b *testing.B) {
	nm := gabs.New() // ignore arrays for now when seting
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, data := range gen {
			if _, err := nm.Set(data.data, data.path...); err != nil {
				b.Log(err, data.path)
			}
		}
	}
}

func BenchmarkGabsField(b *testing.B) {
	nm := gabs.New() // ignore arrays for now when seting
	for _, data := range gen {
		if _, err := nm.Set(data.data, data.path...); err != nil {
			b.Log(err, data.path)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, data := range gen {
			if val := nm.Search(data.path...); val.Data() != data.data {
				b.Errorf("Expected %q ,received: %q", data.data, val.Data())
			}
		}
	}
}
