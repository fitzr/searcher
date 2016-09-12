package searcher

import (
	"strconv"
)

type Searcher interface {
	Search(input string) (results []*Entry)
}

type searcher struct {
	entries map[int]*Entry
	index   map[string][]byte
}

func NewSearcher(entriesPath string, textsPath string) Searcher {
	entries, index := load(entriesPath, textsPath)
	return &searcher{entries: entries, index: index}
}

func (s *searcher) Search(input string) (results []*Entry) {
	results = []*Entry{}
	b, ok := s.index[input]
	if !ok {
		return
	}

	ids := decodeIndex(b)
	for _, id := range ids {
		e, ok := s.entries[id]
		if !ok {
			panic("index " + strconv.Itoa(id) + " have no entry")
		}
		results = append(results, e)
	}
	return
}
