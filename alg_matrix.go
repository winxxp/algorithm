package algorithm

// zWalkMatrix按z字形遍历矩阵
// r is row, c is column
func ZWalkMatrix(a [][]int, r, c int) []int {
	zmatrix := make([]int, 0, r*c)

	i, j := 0, 0
	for {
		zmatrix = append(zmatrix, a[i][j])

		if (i^j)&1 == 0 {
			if j++; j == c {
				j = c - 1
				if i++; i == r {
					break
				}
			} else {
				if i > 0 {
					i--
				}
			}
		} else {
			if i++; i == r {
				i = r - 1
				if j++; j == c {
					break
				}
			} else {
				if j > 0 {
					j--
				}
			}
		}
	}

	return zmatrix
}
