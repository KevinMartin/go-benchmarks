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

// Matrix creates a slice of slices.
func Matrix(foo, bar string) string {
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
