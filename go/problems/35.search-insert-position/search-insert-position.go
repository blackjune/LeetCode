package position

func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	var i int
	for i = 0; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
	}
	return i
}
