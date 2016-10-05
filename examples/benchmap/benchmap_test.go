package main

import "testing"

// use go test -bench=. -benchmem

var capitals = map[string]string{
	"Algeria":                "Algiers",
	"Argentina":              "Buenos Aires",
	"Australia":              "Canberra",
	"Austria":                "Vienna",
	"Bahamas":                "Nassau",
	"Belarus":                "Minsk",
	"Bosnia and Herzegovina": "Sarajevo",
	"Brazil":                 "Brasilia",
	"Bulgaria":               "Sofia",
	"Canada":                 "Ottawa",
	"China":                  "Beijing",
	"Croatia":                "Zagreb",
	"Cuba":                   "Havana",
	"Egypt":                  "Cairo",
	"France":                 "Paris",
	"Germany":                "Berlin",
	"Indonesia":              "Jakarta",
	"Ireland":                "Dublin",
	"Jamaica":                "Kingston",
	"Japan":                  "Tokyo",
	"Luxembourg":             "Luxembourg",
}

var sink string

func BenchmarkMapLookup(b *testing.B) {
	var key = []byte{'F', 'r', 'a', 'n', 'c', 'e'}
	var r string
	for n := 0; n < b.N; n++ {
		r = capitals[string(key)]
	}
	sink = r
}

func BenchmarkMapLookup2(b *testing.B) {
	var key = []byte{'F', 'r', 'a', 'n', 'c', 'e'}
	var r string
	for n := 0; n < b.N; n++ {
		k := string(key)
		r = capitals[k]
	}
	sink = r
}
