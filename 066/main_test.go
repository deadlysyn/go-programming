package main

import "testing"

// sample failing test
// func TestFoo(t *testing.T) {
// 	r := Foo()
// 	for i := range r {
// 		if r[i] != "foo" {
// 			t.Errorf("Want foo got %v", r[i])
// 		}
// 	}
// }

// sample passing test
func TestFoo(t *testing.T) {
	r := Foo()
	for i := range r {
		if r[i] != "foo" && r[i] != "bar" {
			t.Errorf("Want foo/bar got %v", r[i])
		}
	}
}

// sample benchmark
// run with `go test -bench=.`
func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bar(1000)
	}
}
