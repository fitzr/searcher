package searcher

import (
	"github.com/shogo82148/go-mecab"
	"strings"
)

type mecabParser struct {
	model  mecab.Model
	tagger mecab.MeCab
}

func newMeCab() *mecabParser {
	model, err := mecab.NewModel(map[string]string{"output-format-type": "wakati"})
	if err != nil {
		panic(err)
	}

	return &mecabParser{model: model}
}

func (m *mecabParser) destroy() {
	m.tagger.Destroy()
	m.model.Destroy()
}

func (m *mecabParser) parse(s string) []string {
	tagger, err := m.model.NewMeCab()
	if err != nil {
		panic(err)
	}
	defer tagger.Destroy()

	ret, err := tagger.Parse(s)
	if err != nil {
		panic(err)
	}
	ret = strings.TrimRight(ret, "\n")
	ret = strings.Trim(ret, " ")
	if len(ret) == 0 {
		return []string{}
	}

	return strings.Split(ret, " ")
}
