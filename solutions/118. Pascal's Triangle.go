package solutions

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)

	for i := 1; i <= numRows; i++ {
		res = append(res, make([]int, i))
		for r := 0; r < i; r++ {
			if r == 0 || r == i-1 {
				res[i-1][r] = 1
			} else {
				res[i-1][r] = res[i-2][r-1] + res[i-2][r]
			}
		}
	}

	return res
}
