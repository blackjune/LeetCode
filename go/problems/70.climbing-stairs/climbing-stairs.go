package dynamic_programming

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var (
		mNMinus2 int // f(n-2)
		mMinus1  int // f(n - 1)
		m        int // f(n)
	)
	mNMinus2 = 1
	mMinus1 = 2
	for i := 3; i <= n; i++ {
		m = mMinus1 + mNMinus2
		mNMinus2 = mMinus1
		mMinus1 = m
	}
	return m
}
