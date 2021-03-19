package main

import (
	"testing"

	"github.com/chrhlnd/dynjson"
)

// no set ignore

func TestDynJSON(t *testing.T) {
	dyn := dynjson.NewFromBytes([]byte("{}"))

	var node dynjson.DynNode
	var err error

	if node, err = dyn.Node("/some/path/that/is/cool"); err != nil {
		t.Fatalf("My path didn't parse error %v", err)
	}

	if _, err = node.Str(); err != nil {
		t.Fatalf("My value wasn't a string? %v", err)
	}
}

// func BenchmarkDynJSONSet(b *testing.B) {
// 	nm := dynjson.New() // ignore arrays for now when seting
// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		for _, data := range gen {
// 			if r := nm.Set(data.data, data.path...); r.Error() != nil {
// 				b.Log(r.Error(), data.path)
// 			}
// 		}
// 	}
// }

// func BenchmarkDynJSONField(b *testing.B) {
// 	nm := dynjson.New(map[string]interface{}{}) // ignore arrays for now when seting
// 	for _, data := range gen {
// 		if nm = nm.Set(data.data, data.path...); nm.Error() != nil {
// 			b.Log(nm.Error(), data.path)
// 		}
// 	}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		for _, data := range gen {
// 			if val := nm.Get(data.path...); val.Value() != data.data {
// 				b.Errorf("Expected %q ,received: %q", data.data, val.Value())
// 			}
// 		}
// 	}
// }
