package main

import (
	"fmt"
	"testing"

	"github.com/cgrates/cgrates/utils"
)

func TestNavMAp(t *testing.T) {
	nm := utils.NavigableMap{}
	for _, data := range gen {
		if _, err := nm.Set(data.pathItems, utils.NewNMData(data.data)); err != nil {
			t.Error(err, data.path)
		}
	}
}
func TestNode(t *testing.T) {
	nm := &Node{Type: NMMapType, Map: make(map[string]*Node)}
	// fmt.Println(nm.Set(gen[0].data, gen[0].path))
	fmt.Println(nm.Set(gen[0].data, gen[0].path))
	fmt.Println(nm.Field(gen[0].path))
	fmt.Println(utils.ToJSON(nm))
}

func BenchmarkOrderdNavigableMapSet(b *testing.B) {
	nm := utils.NewOrderedNavigableMap()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, data := range gen {
			if _, err := nm.Set(data.fp, utils.NewNMData(data.data)); err != nil {
				b.Log(err, data.path)
			}
		}
	}
}

func BenchmarkNavigableMapSet(b *testing.B) {
	nm := utils.NavigableMap{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, data := range gen {
			if _, err := nm.Set(data.pathItems, utils.NewNMData(data.data)); err != nil {
				b.Log(err, data.path)
			}
		}
	}
}

func BenchmarkNodeSet(b *testing.B) {
	nm := &Node{Type: NMMapType, Map: make(map[string]*Node)}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, data := range gen {
			if err := nm.Set(data.data, data.path); err != nil {
				b.Log(err, data.path)
			}
		}
	}
}

func BenchmarkOrderdNavigableMapField(b *testing.B) {
	nm := utils.NewOrderedNavigableMap()
	for _, data := range gen {
		data.fp.PathItems = data.pathItems.Clone()
		if _, err := nm.Set(data.fp, utils.NewNMData(data.data)); err != nil {
			b.Log(err, data.path)
		}
	}
	path := make([]utils.PathItems, b.N) // this is updated by field
	for i := 0; i < b.N; i++ {
		path[i] = gen[0].pathItems.Clone()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, data := range gen {
			if val, err := nm.Field(path[i]); err != nil {
				b.Log(err)
			} else if val.Interface() != data.data {
				b.Errorf("Expected %q ,received: %q", data.data, val.Interface())
			}
		}
	}
}

func BenchmarkNavigableMapField(b *testing.B) {
	nm := utils.NavigableMap{}
	for _, data := range gen {
		if _, err := nm.Set(data.pathItems.Clone(), utils.NewNMData(data.data)); err != nil {
			b.Log(err, data.path)
		}
	}
	path := make([]utils.PathItems, b.N) // this is updated by field
	for i := 0; i < b.N; i++ {
		path[i] = gen[0].pathItems.Clone()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, data := range gen {
			if val, err := nm.Field(path[i]); err != nil {
				b.Log(err)
			} else if val.Interface() != data.data {
				b.Errorf("Expected %q ,received: %q", data.data, val.Interface())
			}
		}
	}
}

func BenchmarkNodeField(b *testing.B) {
	nm := &Node{Type: NMMapType, Map: make(map[string]*Node)}
	for _, data := range gen {
		if err := nm.Set(data.data, data.path); err != nil {
			b.Log(err, data.path)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, data := range gen {
			if val, err := nm.Field(data.path); err != nil {
				b.Log(err)
			} else if val != data.data {
				b.Errorf("Expected %q ,received: %q", data.data, val)
			}
		}
	}
}
