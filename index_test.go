package searcher

import (
	"reflect"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	expected := map[string][]byte{
		"寒い": []byte{123 + 128},
		"眠い": []byte{111 + 128},
		"もう": []byte{123 + 128, 12 + 128}}

	sut := newIndexBuilder()

	sut.put(123, []string{"もう", "寒い"})
	sut.put(111, []string{"もう", "眠い", "眠い"})

	actual := sut.build()

	for k, v := range actual {
		if !reflect.DeepEqual(expected[k], v) {
			t.Errorf("\nexpected: %v\nactual: %v on key: %v", expected[k], v, k)
		}
	}
}

func TestDecodeIndex(t *testing.T) {
	input := []byte{123 + 128, 12 + 128}
	expected := []int{123, 111}

	actual := decodeIndex(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}
