package searcher

import (
	"bytes"
	"github.com/deckarep/golang-set"
	"sort"
)

const indexMax = 50

type indexBuilder struct {
	wordMap map[string]mapset.Set
}

func newIndexBuilder() *indexBuilder {
	b := &indexBuilder{wordMap: make(map[string]mapset.Set)}
	return b
}

// called parallel
func (b *indexBuilder) put(id int, words []string) {
	for _, w := range words {
		v, ok := b.wordMap[w]
		if !ok {
			v = mapset.NewThreadUnsafeSet()
			b.wordMap[w] = v
		}
		v.Add(id)
	}
}

func (b *indexBuilder) build() map[string][]byte {
	index := map[string][]byte{}
	for k, v := range b.wordMap {
		index[k] = createIndex(v)
	}
	return index
}

func createIndex(set mapset.Set) []byte {
	slice := set.ToSlice()
	ids := make([]int, len(slice))
	for i, id := range slice {
		ids[i] = id.(int)
	}

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

func decodeIndex(b []byte) (ids []int) {
	ids = []int{}
	reader := bytes.NewReader(b)
	id := 0
	for {
		i, err := Decode(reader)
		if err != nil {
			if err.Error() == "EOF" {
				return
			} else {
				panic(err)
			}
		}
		if id == 0 {
			id = i
		} else {
			id -= i
		}
		ids = append(ids, id)
	}
	return
}
