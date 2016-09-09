package searcher

import (
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	input := "もう眠い"
	expected := []string{"もう", "眠い"}

	sut := newMeCab()
	actual := sut.parse(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}
