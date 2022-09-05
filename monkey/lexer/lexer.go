package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  //現在の位置を示す
	readPosition int  //次の文字
	ch           byte //検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() { //次の位置文字を読んでinput文字列の現在位置を進める

	if l.readPosition >= len(l.input) { //入力が終端に到達したかの吟味
		l.ch = 0 //asciiコードの"NULL"に対応している -> まだ何も読み込んでいない or EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition // positionの更新
	l.readPosition += 1 // readpositionの更新

}

func (l *Lexer) NextToken() token.Token { // 現在検査中の文字l.chを見て、その文字が何であるかに応ずてtokenを返す
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newtoken(token.ASSIGN,l.ch)
	case ';':
		tok = newtoken(token.SEMICOLON,l.ch)
	case '(':
		tok = newtoken(token.LPAREN,l.ch)
	case ')':
		tok = newtoken(token.RPAREN,l.ch)
	case ',':
		tok = newtoken(token.COMMA,l.ch)
	case '+':
		tok = newtoken(token.PLUS,l.ch)
	case '{':
		tok = newtoken(token.LBRACE,l.ch)
	case '}':
		tok = newtoken(token.RBRACE,l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newtoken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType,Literal:string(ch)}
}