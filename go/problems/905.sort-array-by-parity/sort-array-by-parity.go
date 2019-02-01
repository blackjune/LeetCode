package array

func sortArrayByParity(A []int) []int {
	B := make([]int, len(A))
	even := 0
	odd := len(A) - 1
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			B[even] = A[i]
			even++
		} else {
			B[odd] = A[i]
			odd--
		}
	}
	return B
}
