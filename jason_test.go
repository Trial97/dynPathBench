package main

import (
	"log"
	"testing"

	"github.com/antonholmquist/jason"
)

// no set ignore
func TestJason(t *testing.T) {
	v, _ := jason.NewObjectFromBytes([]byte("{}"))

	name, _ := v.GetString("name")
	age, _ := v.GetNumber("age")
	occupation, _ := v.GetString("other", "occupation")
	years, _ := v.GetNumber("other", "years")

	log.Println("age:", age)
	log.Println("name:", name)
	log.Println("occupation:", occupation)
	log.Println("years:", years)
}

// func BenchmarkJasonSet(b *testing.B) {
// no set
// nm, _ := jason.NewObjectFromBytes([]byte("{}"))
// b.ResetTimer()
// for n := 0; n < b.N; n++ {
// 	for _, data := range gen {
// 		if r := nm.Set(data.data, data.path...); r.Error() != nil {
// 			b.Log(r.Error(), data.path)
// 		}
// 	}
// }
// }
/*
func BenchmarkJasonField(b *testing.B) {
	nm, _ := jason.NewObjectFromBytes([]byte("{}"))
	for _, data := range gen {
		if nm = nm.Set(data.data, data.path...); nm.Error() != nil {
			b.Log(nm.Error(), data.path)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, data := range gen {
			if val := nm.Get(data.path...); val.Value() != data.data {
				b.Errorf("Expected %q ,received: %q", data.data, val.Value())
			}
		}
	}
}
*/
