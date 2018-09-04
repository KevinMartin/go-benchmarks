package keyValueToString

import (
	"testing"
)

var tests = []struct {
	name string
	test func(string, string) string
}{
	{"MapWithSort", MapWithSort},
	{"MatrixStringJoin", MatrixStringJoin},
	{"MatrixStringBuilder", MatrixStringBuilder},
}

func Test(t *testing.T) {
	expect := "bar=BAR,foo=FOO"

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if v.test("FOO", "BAR") != expect {
				t.Error(v.name + " did not match expected output.")
			}
		})
	}
}

func Benchmark(b *testing.B) {
	for _, v := range tests {
		b.Run(v.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v.test("FOO", "BAR")
			}
		})
	}
}
