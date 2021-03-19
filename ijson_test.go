package main

import (
	"fmt"
	"testing"

	"github.com/akshaybharambe14/ijson"
)

func TestIJSON(t *testing.T) {
	r := ijson.ParseBytes([]byte{}).
		GetP("#0.friends.#~name"). // list the friend names for 0th record -
		// []interface {}{"Justine Bird", "Justine Bird", "Marianne Rutledge"}

		Del("#0"). // delete 0th record
		// []interface {}{"Marianne Rutledge", "Justine Bird"}

		Set("tom", "#") // append "tom" in the list
		// []interface {}{"Marianne Rutledge", "Justine Bird", "tom"}

	fmt.Printf("%#v\n", r.Value())

	// returns error if the data type differs than the type expected by query
	fmt.Println(r.Set(1, "name").Error())
}

func BenchmarkIJSONSet(b *testing.B) {
	nm := ijson.New(map[string]interface{}{}) // ignore arrays for now when seting
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, data := range gen {
			if r := nm.Set(data.data, data.path...); r.Error() != nil {
				b.Log(r.Error(), data.path)
			}
		}
	}
}

func BenchmarkIJSONField(b *testing.B) {
	nm := ijson.New(map[string]interface{}{}) // ignore arrays for now when seting
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
