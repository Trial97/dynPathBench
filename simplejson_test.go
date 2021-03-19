package main

import (
	"bytes"
	"fmt"
	"testing"

	simplejson "github.com/bitly/go-simplejson"
)

func Test(t *testing.T) {
	buf := bytes.NewBuffer([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 9223372036854775807,
			"uint64": 18446744073709551615
		}
	}`))
	fmt.Println(simplejson.NewFromReader(buf))
}

func BenchmarkSimpleJSONSet(b *testing.B) {
	nm := simplejson.New() // ignore arrays for now when seting
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, data := range gen {
			nm.SetPath(data.path, data.data)
		}
	}
}

func BenchmarkSimpleJSONField(b *testing.B) {
	nm := simplejson.New() // ignore arrays for now when seting
	for _, data := range gen {
		nm.SetPath(data.path, data.data)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, data := range gen {
			if val := nm.GetPath(data.path...); val.Interface() != data.data {
				b.Errorf("Expected %q ,received: %q", data.data, val.Interface())
			}
		}
	}
}
