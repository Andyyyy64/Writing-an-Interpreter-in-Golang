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

func (l *Lexer) readIdentifier() string { // 識別子を読んで、非英字に到達するまで字句解析器の位置を進める
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool { // 与えられた引数が英字かどうか判定する　ch == '_' は _ も英字に含めているということ -> foo_bar などの変数名が使える
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType,Literal:string(ch)}
}

func (l *Lexer) skipWhitespace(){
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool { // 与えられた引数が数字かどうか判定する 整数しか読まないことで単純化している
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token { // 現在検査中の文字l.chを見て、その文字が何であるかに応じてtokenを返す
	var tok token.Token

	l.skipWhitespace()
	
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN,l.ch)
	case ';':
		tok = newToken(token.SEMICOLON,l.ch)
	case '(':
		tok = newToken(token.LPAREN,l.ch)
	case ')':
		tok = newToken(token.RPAREN,l.ch)
	case ',':
		tok = newToken(token.COMMA,l.ch)
	case '+':
		tok = newToken(token.PLUS,l.ch)
	case '{':
		tok = newToken(token.LBRACE,l.ch)
	case '}':
		tok = newToken(token.RBRACE,l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch){
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL,l.ch)
		}
	}
	l.readChar()
	return tok

}
