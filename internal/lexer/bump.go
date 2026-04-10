package lexer

func (l *Lexer) bump() {
	_, _ = l.nextRune()
}

func (l *Lexer) bumpUntil(fn func(rune) bool) {
	for {
		r, err := l.peekRune()
		if err != nil {
			break
		}

		if fn(r) {
			break
		}

		l.bump()
	}
}

func (l *Lexer) bumpWhile(fn func(rune) bool) {
	for {
		r, err := l.peekRune()
		if err != nil {
			break
		}

		if !fn(r) {
			break
		}

		l.bump()
	}
}
