package array

func flipAndInvertImage(A [][]int) [][]int {
	B := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		for j := len(A[i]) - 1; j >= 0; j-- {
			B[i] = append(B[i], A[i][j]^1)
		}
	}
	return B
}
