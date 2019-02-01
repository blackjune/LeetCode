package array

func sortedSquares(A []int) []int {
	n := len(A)
	a := make([]int, n)
	j := 0
	for j < n && A[j] < 0 {
		j++
	}
	i := j - 1
	t := 0
	for i >= 0 && j < n {
		if A[i]*A[i] < A[j]*A[j] {
			a[t] = A[i] * A[i]
			i--
		} else {
			a[t] = A[j] * A[j]
			j++
		}
		t++
	}
	for ; i >= 0; i-- {
		a[t] = A[i] * A[i]
		t++
	}
	for ; j < n; j++ {
		a[t] = A[j] * A[j]
		t++
	}
	return a
}
