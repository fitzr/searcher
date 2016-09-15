package searcher

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type loader struct {
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

func load(entriesPath string, textsPath string) (map[int]*Entry, map[string][]byte) {

	l := loader{
		builder:      newIndexBuilder(),
		mecab:        newMeCab(),
		wordsChannel: make(chan words),
		wg:           sync.WaitGroup{}}
	defer l.destroy()

	files := l.listFiles(textsPath)

	l.wg.Add(len(files) + 1)

	go l.readEntry(entriesPath)
	for _, file := range files {
		go l.readText(textsPath, file.Name())
	}
	go l.collect()

	l.wg.Wait()

	return l.entries, l.builder.build()
}

func (l *loader) listFiles(root string) []os.FileInfo {
	files, err := ioutil.ReadDir(root)
	check(err)
	return files
}

func (l *loader) readEntry(entries string) {
	l.entries = readEntries(entries)
	for _, e := range l.entries {
		l.wordsChannel <- words{e.ID, l.mecab.parse(e.Title)}
	}
	defer l.wg.Done()
}

func (l *loader) readText(root string, file string) {
	id, err := strconv.Atoi(file)
	check(err)

	data, err := ioutil.ReadFile(filepath.Join(root, file))
	check(err)
	l.wordsChannel <- words{id, l.mecab.parse(string(data[:]))}

	defer l.wg.Done()
}

func (l *loader) collect() {
	for words := range l.wordsChannel {
		l.builder.put(words.id, words.words)
	}
}

func (l *loader) destroy() {
	l.mecab.destroy()
	close(l.wordsChannel)
}
