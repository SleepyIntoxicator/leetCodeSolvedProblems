package solutions

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var sp int
	stack := make([]byte, len(s)/2+1)

	for _, c := range []byte(s) {
		switch c {
		case '(':
			stack[sp] = ')'
			sp++
		case '[':
			stack[sp] = ']'
			sp++
		case '{':
			stack[sp] = '}'
			sp++
		case ')', ']', '}':
			if sp > 0 && c == stack[sp-1] {
				sp--
			} else {
				return false
			}
		}
	}
	if sp != 0 {
		return false
	}

	return true
}
