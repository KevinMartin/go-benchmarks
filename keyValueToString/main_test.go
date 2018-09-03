package keyValueToString

import (
	"testing"
)

var str string

var tests = []struct{
	name string
	test func(string, string) string
}{
	{"MapWithSort", MapWithSort},
	{"Matrix", Matrix},
}

func Test(t *testing.T) {
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			str = v.test("FOO", "BAR")
			if str != "bar=BAR,foo=FOO" {
				t.Error(v.name + " did not match expected output.")
			}
		})
	}
}

func Benchmark(b *testing.B) {
	for _, v := range tests {
		b.Run(v.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				str = v.test("FOO", "BAR")
			}
		})
	}
}
