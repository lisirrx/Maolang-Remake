package inputstream

import (
	"io"
	"io/ioutil"
)

type InputStream interface {
	Char() rune
	Next()
	EOF() bool
}

type MaoLangInputStream struct {
	reader     io.Reader
	buf        []rune
	currentPos int
	current    rune
}

func (m *MaoLangInputStream) Init() {
	buffer, err := ioutil.ReadAll(m.reader)
	if err != nil {
		panic(err)
	}
	m.buf = []rune(string(buffer))
	m.current = m.buf[0]
	m.currentPos = 0
}

func (m *MaoLangInputStream) Char() rune {
	return m.current
}
func (m *MaoLangInputStream) EOF() bool {
	return m.currentPos >= len(m.buf)
}

func (m *MaoLangInputStream) Next() {
	m.currentPos++
	if !m.EOF() {
		m.current = m.buf[m.currentPos]
	} else {
		m.current = 0
	}
}

func NewMaoLanInputStream(r io.Reader) InputStream {
	ml := MaoLangInputStream{reader: r}
	ml.Init()
	return &ml
}
