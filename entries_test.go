package searcher

import (
	"reflect"
	"testing"
)

func TestReadEntries(t *testing.T) {

	expected := &entry{id: 15271537, category: 5, url: "http://strikewatches.sblo.jp/article/31222596.html", title: "31222596.html"}
	actual := readEntries("./data/10000entries.txt")

	if len(actual) != 10000 {
		t.Errorf("\nexpected: %v\nactual: %v", 10000, len(actual))
	}
	entry, _ := actual[15271537]
	if !reflect.DeepEqual(expected, entry) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, entry)
	}
}
