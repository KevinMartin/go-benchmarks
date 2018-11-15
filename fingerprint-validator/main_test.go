package fingerprintValidator

import (
	"testing"
)

var tests = []struct {
	name string
	run  func(int, string) bool
}{
	{"ValidateStringsSplitStringsCount", ValidateStringsSplitStringsCount},
	{"ValidateStringsSplitStrings", ValidateStringsSplitStrings},
	{"ValidateLoopPairsCount", ValidateLoopPairsCount},
}

func Test(t *testing.T) {
	testCases := map[string]struct {
		pairs  int
		fp     string
		expect bool
	}{
		"too_short":  {3, "12:34:56:78", false},
		"too_long":   {5, "12:34:56:78", false},
		"too_many":   {4, "12:34:56:789", false},
		"too_few":    {4, "12:34:5:78", false},
		"bad_hex":    {4, "12:34:56:7z", false},
		"pipe_sep":   {4, "12:34:56|78", false},
		"pipe_zero":  {4, "12:34:56078", false},
		"just_right": {4, "12:34:56:78", true},
	}

	for _, fn := range tests {
		for k, tc := range testCases {
			t.Run(fn.name+"_"+k, func(t *testing.T) {
				if fn.run(tc.pairs, tc.fp) != tc.expect {
					t.Error(fn.name + " did not match expected output.")
				}
			})
		}
	}
}

func Benchmark(b *testing.B) {
	testCases := map[string]struct {
		pairs int
		fp    string
	}{
		"small":  {4, "12:34:56:78"},
		"medium": {8, "12:34:56:78:12:34:56:78"},
		"large":  {16, "12:34:56:78:12:34:56:78:12:34:56:78:12:34:56:78"},
	}

	for _, fn := range tests {
		for k, tc := range testCases {
			b.Run(fn.name+"_"+k, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					fn.run(tc.pairs, tc.fp)
				}
			})
		}
	}
}
