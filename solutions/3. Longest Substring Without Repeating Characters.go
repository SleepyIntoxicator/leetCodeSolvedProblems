package solutions

func lengthOfLongestSubstring(s string) int {
	repHash := make(map[byte]int, 0)
	maxLen := 0

	for i := range s {
		if dupIndex, ok := repHash[s[i]]; ok {
			if len(repHash) > maxLen {
				maxLen = len(repHash)
			}
			for k, v := range repHash {
				if v < dupIndex {
					delete(repHash, k)
				}
				repHash[s[i]] = i
			}
		} else {
			repHash[s[i]] = i
		}
	}

	if len(repHash) > maxLen {
		return len(repHash)
	}

	return maxLen
}
