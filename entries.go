package searcher

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Entry of diary.
type Entry struct {
	ID       int
	Category int
	URL      string
	Title    string
}

func readEntries(filePath string) map[int]*Entry {
	fp, err := os.Open(filePath)
	check(err)
	defer func () { check(fp.Close())}()
	scanner := bufio.NewScanner(fp)

	entries := map[int]*Entry{}
	for scanner.Scan() {
		entry := toEntry(scanner.Text())
		entries[entry.ID] = entry
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
		ID:       id,
		Category: category,
		URL:      texts[2],
		Title:    texts[3]}
}
