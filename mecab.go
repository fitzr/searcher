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

	tagger, err := model.NewMeCab()
	if err != nil {
		panic(err)
	}

	return &mecabParser{model: model, tagger: tagger}
}

func (m *mecabParser) destroy() {
	m.tagger.Destroy()
	m.model.Destroy()
}

func (m *mecabParser) parse(s string) []string {
	ret, err := m.tagger.Parse(s)
	if err != nil {
		panic(err)
	}
	ret = ret[:len(ret)-2] // remove last space
	return strings.Split(ret, " ")
}
