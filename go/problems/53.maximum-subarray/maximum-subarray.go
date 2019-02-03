package dynamic_programming

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+sum > nums[i] {
			sum = nums[i] + sum
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
