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
	check(err)

	return &mecabParser{model: model}
}

func (m *mecabParser) destroy() {
	m.model.Destroy()
}

func (m *mecabParser) parse(s string) []string {
	tagger, err := m.model.NewMeCab()
	check(err)
	defer tagger.Destroy()

	ret, err := tagger.Parse(s)
	check(err)
	ret = strings.TrimSuffix(ret, "\n")
	ret = strings.TrimSpace(ret)
	if len(ret) == 0 {
		return []string{}
	}

	return strings.Fields(ret)
}
