package searcher

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	id       int
	category int
	url      string
	title    string
}

func readEntries(filePath string) map[int]*entry {
	fp, err := os.Open(filePath)
	check(err)
	defer fp.Close()
	scanner := bufio.NewScanner(fp)

	entries := map[int]*entry{}
	for scanner.Scan() {
		entry := toEntry(scanner.Text())
		entries[entry.id] = entry
	}
	check(scanner.Err())

	return entries
}

func toEntry(s string) *entry {
	texts := strings.Split(s, "\t")
	id, err := strconv.Atoi(texts[0])
	check(err)
	category, err := strconv.Atoi(texts[1])
	check(err)
	return &entry{
		id:       id,
		category: category,
		url:      texts[2],
		title:    texts[3]}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
