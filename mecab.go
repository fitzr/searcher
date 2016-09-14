package searcher

import (
	"strings"

	"github.com/shogo82148/go-mecab"
)

type mecabParser struct {
	model mecab.Model
}

func newMeCab() *mecabParser {
	model, err := mecab.NewModel(map[string]string{"output-format-type": "wakati"})
	if err != nil {
		panic(err)
	}

	return &mecabParser{model: model}
}

func (m *mecabParser) destroy() {
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
