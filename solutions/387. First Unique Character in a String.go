package solutions

func firstUniqCharFromLeetCode(s string) int {
	hashStr := make([]int, 26)
	for i := range s {
		hashStr[s[i]-'a']++
	}
	for i := range s {
		if hashStr[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqCharFirst(s string) int {
	hashStr := make(map[byte]int)
	for i := range s {
		if _, ok := hashStr[s[i]]; ok {
			hashStr[s[i]] = hashStr[s[i]] + 1
		} else {
			hashStr[s[i]] = 1
		}
	}
	for i := range s {
		if c := hashStr[s[i]]; c == 1 {
			return i
		}
	}

	return -1
}
