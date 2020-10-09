package parser

import tk "me.lisirrx/maolang/token"

type MaoAstRoot struct {
	maoDecl  MaoAstDecl
	maoOp    MaoAstOp
	maoPrint MapAstPrint
}

type MaoAstDecl struct {
	maoAstDeclStatements []MaoAstDeclStatement
}

type MaoAstOp struct {
	maoAstOpStatements []MaoAstOpStatement
}

type MapAstPrint struct {
	maoAstPrintStatements []MaoAstPrintStatement
}

type MaoAstDeclStatement struct {
	typeToken tk.Token
	varTokens []tk.Token
}

type MaoAstOpStatement struct{

}

type MaoAstPrintStatement struct {
	
}
