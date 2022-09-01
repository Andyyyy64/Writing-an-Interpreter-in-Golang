package lexer

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
