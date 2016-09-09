package searcher

import (
	"bytes"
	"sort"
	"sync"
)

const indexMax = 50

type indexBuilder struct {
	wordMap map[string][]int
	mu      sync.Mutex
}

func newIndexBuilder() *indexBuilder {
	b := &indexBuilder{
		wordMap: make(map[string][]int),
		mu:      sync.Mutex{}}
	return b
}

// called parallel
func (b *indexBuilder) put(id int, words []string) {
	b.mu.Lock()
	for _, w := range words {
		v, ok := b.wordMap[w]
		if ok {
			b.wordMap[w] = append(v, id)
		} else {
			b.wordMap[w] = []int{id}
		}
	}
	b.mu.Unlock()
}

func (b *indexBuilder) build() map[string][]byte {
	index := map[string][]byte{}
	for k, v := range b.wordMap {
		index[k] = createIndex(v)
	}
	return index
}

func createIndex(ids []int) []byte {
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	if len(ids) > indexMax {
		ids = ids[:indexMax]
	}

	buff := new(bytes.Buffer)
	Encode(buff, ids[0])
	for i := 1; i < len(ids); i++ {
		Encode(buff, ids[i-1]-ids[i])
	}

	return buff.Bytes()
}
