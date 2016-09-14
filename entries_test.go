package searcher

import (
	"reflect"
	"testing"
)

func TestReadEntries(t *testing.T) {

	expected := &Entry{ID: 15271537, Category: 5, URL: "http://strikewatches.sblo.jp/article/31222596.html", Title: "31222596.html"}
	actual := readEntries("./testdata/10000entries.txt")

	if len(actual) != 10000 {
		t.Errorf("\nexpected: %v\nactual: %v", 10000, len(actual))
	}
	entry, _ := actual[15271537]
	if !reflect.DeepEqual(expected, entry) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, entry)
	}
}
