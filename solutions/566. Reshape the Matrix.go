package solutions

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	resMat := make([][]int, 0, r)
	resMat = append(resMat, make([]int, c))

	rI, rJ := 0, 0

	for i := range mat {
		for _, jt := range mat[i] {
			resMat[rI][rJ] = jt
			rJ++

			if rJ == c {
				rI++
				rJ = 0
				if rI != r {
					resMat = append(resMat, make([]int, c))
				}
			}
		}
	}

	return resMat
}

func matrixReshapeFromLeetCode(mat [][]int, r int, c int) [][]int {
	itemsNumber := len(mat) * len(mat[0])
	if itemsNumber != r*c || len(mat) == r {
		return mat
	}

	tmp := make([]int, 0, itemsNumber)
	res := make([][]int, 0, r)

	for i := 0; i < len(mat); i++ {
		tmp = append(tmp, mat[i]...) // The elements are copied to the new memory area
	}

	for i := 0; i < itemsNumber; i += c {
		res = append(res, tmp[i:i+c])
	}

	return res
}

// r  - высота,	c  - длина
// rI - высота, rJ - длина
// rI, rJ -> resMat[rI][rJ]
