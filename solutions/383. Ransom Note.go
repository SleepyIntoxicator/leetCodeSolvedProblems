package solutions

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) { //from leetCode
		return false
	}

	hashStr := make([]int, 26)
	for i := range magazine {
		hashStr[magazine[i]-'a']++
	}
	for i := range ransomNote {
		hashStr[ransomNote[i]-'a']--
		if hashStr[ransomNote[i]-'a'] < 0 {
			return false
		}
	}

	return true
}
