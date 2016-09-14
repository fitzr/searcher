package searcher

import (
	"bytes"
	"sort"
)

const indexMax = 50

type indexBuilder struct {
	wordMap map[string][]int
}

func newIndexBuilder() *indexBuilder {
	b := &indexBuilder{wordMap: make(map[string][]int)}
	return b
}

func (b *indexBuilder) put(id int, words []string) {
	for _, w := range words {
		v, ok := b.wordMap[w]
		if ok {
			b.wordMap[w] = append(v, id)
		} else {
			b.wordMap[w] = []int{id}
		}
	}
}

func (b *indexBuilder) build() map[string][]byte {
	index := map[string][]byte{}
	for k, v := range b.wordMap {
		index[k] = encodeIndex(v)
	}
	return index
}

func encodeIndex(ids []int) []byte {
	ids = deduplicate(ids)
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	if len(ids) > indexMax {
		ids = ids[:indexMax]
	}

	buff := new(bytes.Buffer)
	err := Encode(buff, ids[0])
	check(err)
	for i := 1; i < len(ids); i++ {
		err = Encode(buff, ids[i-1]-ids[i])
		check(err)
	}

	return buff.Bytes()
}

func deduplicate(ids []int) []int {
	m := map[int]struct{}{}
	for _, id := range ids {
		m[id] = struct{}{}
	}
	ids = make([]int, len(m))
	i := 0
	for k := range m {
		ids[i] = k
		i++
	}
	return ids
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
			}
			panic(err)
		}
		if id == 0 {
			id = i
		} else {
			id -= i
		}
		ids = append(ids, id)
	}
}
