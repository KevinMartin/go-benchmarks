package keyValueToString

import (
	"sort"
	"strings"
)

// MapWithSort creates a map and sorts the end-value.
func MapWithSort(foo, bar string) string {
	myMap := map[string]string{
		"foo": foo,
		"bar": bar,
	}

	mySlice := make([]string, 0, len(myMap))
	for k, v := range myMap {
		mySlice = append(mySlice, k+"="+v)
	}

	sort.Strings(mySlice)

	return strings.Join(mySlice, ",")
}

// MatrixStringJoin creates a slice of slices and uses strings.Join.
func MatrixStringJoin(foo, bar string) string {
	myMatrix := [][]string{
		{"bar", bar},
		{"foo", foo},
	}

	mySlice := make([]string, 0, len(myMatrix))
	for _, v := range myMatrix {
		mySlice = append(mySlice, strings.Join(v, "="))
	}

	return strings.Join(mySlice, ",")
}

// MatrixStringBuilder creates a slice of slices and uses a strings.Builder
// buffer.
func MatrixStringBuilder(foo, bar string) string {
	myMatrix := [][2]string{
		{"bar", bar},
		{"foo", foo},
	}

	b := strings.Builder{}
	for i, v := range myMatrix {
		b.WriteString(v[0])
		b.WriteRune('=')
		b.WriteString(v[1])
		if i+1 < len(myMatrix) {
			b.WriteRune(',')
		}
	}

	return b.String()
}

// MatrixStringBuilderGrow creates a slice of slices, uses a strings.Builder
// buffer, and pre-allocates a large buffer.
func MatrixStringBuilderGrow(foo, bar string) string {
	myMatrix := [][2]string{
		{"bar", bar},
		{"foo", foo},
	}

	var buflen int
	for i, v := range myMatrix {
		buflen += len(v[0]) + len(v[1]) + 1
		if i+1 < len(myMatrix) {
			buflen++
		}
	}

	b := strings.Builder{}
	b.Grow(buflen)
	for i, v := range myMatrix {
		b.WriteString(v[0])
		b.WriteRune('=')
		b.WriteString(v[1])
		if i+1 < len(myMatrix) {
			b.WriteRune(',')
		}
	}

	return b.String()
}
