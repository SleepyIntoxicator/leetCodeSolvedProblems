package solutions

import "strconv"

func isPalindromeString(x int) bool {
	strX := strconv.Itoa(x)
	for i := 0; i < len(strX)/2; i++ {
		if strX[i] != strX[len(strX)-1-i] {
			return false
		}
	}
	return true
}

func isPalindromeBruteforce(x int) bool {
	if x < 0 {
		return false
	}
	digits := make([]int, 0)
	digitDelim := 10

	for {
		if digitDelim == 10 {
			digits = append(digits, x%10)
		} else if digitDelim > x {
			digits = append(digits, x/(digitDelim/10))
		} else {
			dig := (x % digitDelim) / (digitDelim / 10)
			digits = append(digits, dig)
		}

		if digitDelim > x {
			break
		}
		digitDelim *= 10
	}

	isPalindrom := true
	j := len(digits) - 1
	for i := 0; i < len(digits); i++ {
		if digits[i] != digits[j] {
			isPalindrom = false
		}
		j--
	}

	return isPalindrom
}

func isPalindromeOptimal(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed, num := 0, x

	for reversed < num {
		reversed = reversed*10 + num%10
		num /= 10
	}

	return reversed == num || reversed/10 == num
}
