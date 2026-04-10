package lexer

func (l *Lexer) takeUntil(predicate func(rune) bool) string {
	var result []rune

	for {
		r, err := l.peekRune()
		if err != nil {
			break
		}

		if predicate(r) {
			break
		}

		l.bump()

		result = append(result, r)
	}

	return string(result)
}

func (l *Lexer) takeWhile(predicate func(rune) bool) string {
	var result []rune

	for {
		r, err := l.peekRune()
		if err != nil {
			break
		}

		if !predicate(r) {
			break
		}

		l.bump()

		result = append(result, r)
	}

	return string(result)
}
