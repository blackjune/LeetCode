package array

func twoSum(numbers []int, target int) []int {
	flags := make(map[int]int)
	flags[target-numbers[0]] = 0
	for j := 1; j < len(numbers); j++ {
		if i, ok := flags[numbers[j]]; ok {
			return []int{i + 1, j + 1}
		}
		flags[target-numbers[j]] = j
	}
	return []int{}
}
