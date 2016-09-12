package searcher

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	Id       int
	Category int
	Url      string
	Title    string
}

func readEntries(filePath string) map[int]*Entry {
	fp, err := os.Open(filePath)
	check(err)
	defer fp.Close()
	scanner := bufio.NewScanner(fp)

	entries := map[int]*Entry{}
	for scanner.Scan() {
		entry := toEntry(scanner.Text())
		entries[entry.Id] = entry
	}
	check(scanner.Err())

	return entries
}

func toEntry(s string) *Entry {
	texts := strings.Split(s, "\t")
	id, err := strconv.Atoi(texts[0])
	check(err)
	category, err := strconv.Atoi(texts[1])
	check(err)
	return &Entry{
		Id:       id,
		Category: category,
		Url:      texts[2],
		Title:    texts[3]}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
