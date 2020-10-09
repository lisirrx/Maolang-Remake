package token

import "fmt"

type TokenType uint8

const (
	// 类型
	TOKEN_PRI_TYPE TokenType = 1
	// 变量名
	TOKEN_VAR_NAME TokenType = 2
	// 操作符
	TOKEN_OP TokenType = 3
	// 分号
	TOKEN_SEMI TokenType = 4
	// 括号
	TOKEN_LEFT_PARENTHESE TokenType = 5

	TOKEN_RIGHT_PARENTHESE TokenType = 6

	// 数字
	TOKEN_NUMBER TokenType = 7

	//EOF
	TOKEN_EOF TokenType = 8

	TOKEN_COMMMA TokenType = 9
)

type Token struct {
	tokenType TokenType
	value     string
}

func (t Token) String() string {
	return fmt.Sprintf("[%s] : [%s]", t.tokenType, t.value)
}

func (t Token) Value() string {
	return t.value
}

func (t Token) Type() TokenType {
	return t.tokenType
}

func (t Token) typeEqual(tokenType TokenType) bool {
	return t.tokenType == tokenType
}
