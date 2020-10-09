package main

import (
	"bufio"
	"fmt"
	"os"

	"me.lisirrx/maolang/token"

	mlInputeStream "me.lisirrx/maolang/inputstream"
)

func main() {
	args := os.Args[1:]

	fileName := args[0]

	f, err := os.Open(fileName)
	if err != nil {
		panic("Read Error")
	}

	fileReader := bufio.NewReader(f)

	m := mlInputeStream.NewMaoLanInputStream(fileReader)

	lexer := token.NewMaoLangLexer(m)
	lexer.Next()

	for !lexer.EOF() {
		fmt.Println(lexer.Token())
		lexer.Next()
	}

}
