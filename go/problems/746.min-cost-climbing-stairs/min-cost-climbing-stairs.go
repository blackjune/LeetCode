package dynamic_programming

func minCostClimbingStairs(cost []int) int {
	f1, f2 := 0, 0
	for i := len(cost) - 1; i >= 0; i-- {
		f0 := cost[i] + minInt(f1, f2)
		f2 = f1
		f1 = f0
	}
	return minInt(f1, f2)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
