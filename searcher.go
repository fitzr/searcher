package searcher

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type Searcher interface {
	Search(input string) (results []*Entry)
}

type searcher struct {
	index   map[string][]byte
	entries map[int]*Entry
}

type initializer struct {
	builder      *indexBuilder
	mecab        *mecabParser
	wordsChannel chan words
	wg           sync.WaitGroup
	entries      map[int]*Entry
}

type words struct {
	id    int
	words []string
}

func NewSearcher(entries string, texts string) Searcher {
	init := initializer{
		builder:      newIndexBuilder(),
		mecab:        newMeCab(),
		wordsChannel: make(chan words),
		wg:           sync.WaitGroup{}}
	defer init.destroy()

	files := init.listFiles(texts)

	init.wg.Add(len(files) + 1)

	go init.readEntry(entries)
	for _, file := range files {
		go init.readText(texts, file.Name())
	}
	go init.collect()

	init.wg.Wait()

	return &searcher{
		entries: init.entries,
		index:   init.builder.build()}
}

func (i *initializer) listFiles(root string) []os.FileInfo {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err)
	}
	return files
}

func (i *initializer) readEntry(entries string) {
	i.entries = readEntries(entries)
	for _, e := range i.entries {
		i.wordsChannel <- words{e.Id, i.mecab.parse(e.Title)}
	}
	defer i.wg.Done()
}

func (i *initializer) readText(root string, file string) {
	id, err := strconv.Atoi(file)
	if err != nil {
		panic(err)
	}

	fp, err := os.Open(filepath.Join(root, file))
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	for sc.Scan() {
		w := i.mecab.parse(sc.Text())
		i.wordsChannel <- words{id, w}
	}
	defer i.wg.Done()
}

func (i *initializer) collect() {
	for words := range i.wordsChannel {
		i.builder.put(words.id, words.words)
	}
}

func (i *initializer) destroy() {
	i.mecab.destroy()
	close(i.wordsChannel)
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
