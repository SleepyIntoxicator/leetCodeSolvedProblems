package solutions

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashStr := [26]int{}
	for _, c := range s {
		hashStr[c-'a']++
	}
	for _, c := range t {
		hashStr[c-'a']--
		if hashStr[c-'a'] < 0 {
			return false
		}
	}

	return true
}
