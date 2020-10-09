package token

import (
	"fmt"
	"strings"

	mlInputeStream "me.lisirrx/maolang/inputstream"
)

type Lexer interface {
	Token() Token
	Next()
	EOF() bool
}

type MaoLangLexer struct {
	current     Token
	previous    Token
	eof         bool
	inputStream mlInputeStream.InputStream
}

func NewMaoLangLexer(inputStream mlInputeStream.InputStream) Lexer {
	return &MaoLangLexer{inputStream: inputStream, eof: false}
}

func (m *MaoLangLexer) Token() Token {
	return m.current
}

func (m *MaoLangLexer) Next() {

	current := m.readNext()

	m.previous = m.current
	m.current = current

	if current.tokenType == TOKEN_EOF {
		m.eof = true
	}
}

func (m *MaoLangLexer) EOF() bool {
	return m.eof
}

func (m *MaoLangLexer) readNext() Token {
	m.skipBlank()

	if m.inputStream.EOF() {
		return Token{tokenType: TOKEN_EOF}
	}

	c := m.inputStream.Char()

	if isSemi(c) {
		m.inputStream.Next()
		return Token{tokenType: TOKEN_SEMI, value: ";"}
	}

	if isOp(c) {
		m.inputStream.Next()
		return Token{tokenType: TOKEN_OP, value: string(c)}
	}

	if isLeftPar(c) {
		m.inputStream.Next()
		return Token{tokenType: TOKEN_LEFT_PARENTHESE, value: string(c)}

	}

	if isRightPar(c) {
		m.inputStream.Next()
		return Token{tokenType: TOKEN_RIGHT_PARENTHESE, value: string(c)}

	}

	if isComma(c) {
		m.inputStream.Next()
		return Token{tokenType: TOKEN_COMMMA, value: string(c)}
	}

	if isWordBegin(c) {
		word := m.readAWord()
		if isTypeKw(word) {
			return Token{tokenType: TOKEN_PRI_TYPE, value: word}
		} else {
			return Token{tokenType: TOKEN_VAR_NAME, value: word}
		}
	}

	if isNumberBegin(c) {
		word := m.readNumber()
		return Token{tokenType: TOKEN_NUMBER, value: word}
	}

	panic(fmt.Sprintf("Invalid Char %s", string(c)))
}

func (m *MaoLangLexer) readAWord() string {
	return m.readUntil(func(c rune) bool {
		return !isPartOfWord(c)
	})
}

func (m *MaoLangLexer) readUntil(judge func(rune) bool) string {
	var tokenStr strings.Builder
	for !m.inputStream.EOF() {
		c := m.inputStream.Char()

		if judge(c) {
			break
		} else {
			tokenStr.WriteRune(c)
			m.inputStream.Next()
		}
	}
	return tokenStr.String()
}

// return when current char is not blank
func (m *MaoLangLexer) skipBlank() {
	for !m.inputStream.EOF() {
		c := m.inputStream.Char()
		if !isBlank(c) {
			break
		}
		m.inputStream.Next()
	}
}

func (m *MaoLangLexer) readNumber() string {
	return m.readUntil(func(c rune) bool {
		return !isNumberBegin(c)
	})
}

func isBlank(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func isNumberBegin(c rune) bool {
	return isNumberChar(c) || c == '.'

}

func isTypeKw(word string) bool {
	tmp := strings.ToUpper(word)
	switch tmp {
	case "INT":
		fallthrough
	case "DOUBLE":
		return true
	}

	return false
}

func isWordBegin(c rune) bool {
	return isLowerLetter(c) || isUpperLetter(c)
}

func isRightPar(c rune) bool {
	return c == ')'
}

func isLeftPar(c rune) bool {
	return c == '('
}

func isOp(c rune) bool {
	switch c {
	case '+':
		fallthrough
	case '-':
		fallthrough
	case '*':
		fallthrough
	case '/':
		fallthrough
	case '=':
		return true
	}
	return false
}

func isSemi(c rune) bool {
	return c == ';'
}

func isPartOfWord(c rune) bool {
	return isLowerLetter(c) || isUpperLetter(c) || isNumberChar(c) || isUnderLine(c)
}

func isLowerLetter(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func isUpperLetter(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func isNumberChar(c rune) bool {
	return c >= '0' && c <= '9'
}

func isUnderLine(c rune) bool {
	return c == '_'
}

func isComma(c rune) bool {
	return c == ','
}
