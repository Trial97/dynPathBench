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
	nm := &DataNode{Type: NMMapType, Map: make(map[string]*DataNode)}
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
	nm := &DataNode{Type: NMMapType, Map: make(map[string]*DataNode)}
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
	path := make([][]utils.PathItems, len(gen)) // this is updated by field
	for j, g := range gen {
		path[j] = make([]utils.PathItems, b.N)
		for i := 0; i < b.N; i++ {
			path[j][i] = g.pathItems.Clone()
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j, data := range gen {
			if val, err := nm.Field(path[j][i]); err != nil {
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
	path := make([][]utils.PathItems, len(gen)) // this is updated by field
	for j, g := range gen {
		path[j] = make([]utils.PathItems, b.N)
		for i := 0; i < b.N; i++ {
			path[j][i] = g.pathItems.Clone()
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j, data := range gen {
			if val, err := nm.Field(path[j][i]); err != nil {
				b.Log(err)
			} else if val.Interface() != data.data {
				b.Errorf("Expected %q ,received: %q", data.data, val.Interface())
			}
		}
	}
}

func BenchmarkNodeField(b *testing.B) {
	nm := &DataNode{Type: NMMapType, Map: make(map[string]*DataNode)}
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

func BenchmarkOrderdNavigableMapRemove_(b *testing.B) {
	nms := make([]*utils.OrderedNavigableMap, b.N)
	for i := 0; i < b.N; i++ {
		nms[i] = utils.NewOrderedNavigableMap()
		for _, data := range gen {
			data.fp.PathItems = data.pathItems.Clone()
			if _, err := nms[i].Set(data.fp, utils.NewNMData(data.data)); err != nil {
				b.Log(err, data.path)
			}
		}
	}
	path := make([][]*utils.FullPath, len(gen)) // this is updated by field
	for j, g := range gen {
		path[j] = make([]*utils.FullPath, b.N)
		for i := 0; i < b.N; i++ {
			path[j][i] = &utils.FullPath{
				Path:      g.fp.Path,
				PathItems: g.pathItems.Clone(),
			}
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range gen {
			if err := nms[i].Remove(path[j][i]); err != nil {
				b.Log(err)
			}
		}
	}
}

func BenchmarkNavigableMapRemove(b *testing.B) {
	nms := make([]utils.NavigableMap, b.N)
	for i := 0; i < b.N; i++ {
		nms[i] = utils.NavigableMap{}
		for _, data := range gen {
			if _, err := nms[i].Set(data.pathItems.Clone(), utils.NewNMData(data.data)); err != nil {
				b.Log(err, data.path)
			}
		}
	}
	path := make([][]utils.PathItems, len(gen)) // this is updated by Remove
	for j, g := range gen {
		path[j] = make([]utils.PathItems, b.N)
		for i := 0; i < b.N; i++ {
			path[j][i] = g.pathItems.Clone()
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range gen {
			// fmt.Printf("%q\t", path[j][i][0].Field)
			// fmt.Printf("%q\n", path[j][i][0].Index)
			if err := nms[i].Remove(path[j][i]); err != nil {
				b.Log(err, path[j][i])
			}
		}
	}
}

func BenchmarkNodeRemove(b *testing.B) {
	nms := make([]*DataNode, b.N)
	for i := 0; i < b.N; i++ {
		nms[i] = &DataNode{Type: NMMapType, Map: make(map[string]*DataNode)}
		for _, data := range gen {
			if err := nms[i].Set(data.data, data.path); err != nil {
				b.Log(err, data.path)
			}
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, data := range gen {
			if err := nms[i].Remove(data.path); err != nil {
				b.Log(err)
			}
		}
	}
}
