package arrays

func intersection(nums1 []int, nums2 []int) []int {
	var mem = make(map[int]bool)
	var res []int
	for i := 0; i < len(nums1); i++ {
		if _, ok := mem[nums1[i]]; !ok {
			mem[nums1[i]] = true
		}
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := mem[nums2[i]]; ok {
			delete(mem, nums2[i])
			res = append(res, nums2[i])
		}
	}
	return res
}
