package array

func transpose(A [][]int) [][]int {
	if len(A) == 0 {
		return A
	}
	row := len(A)
	col := len(A[0])
	a := make([][]int, col)
	for j := 0; j < col; j++ {
		a[j] = make([]int, row)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			a[j][i] = A[i][j]
		}
	}
	return a
}
