package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}
	k := len(nums1) - 1
	var i int
	for i = n - 1; i >= 0 && k >= m && m > 0; {
		var flg bool
		if nums2[i] > nums1[m-1] {
			nums1[k] = nums2[i]
			flg = true
		} else {
			nums1[m-1], nums1[k] = nums1[k], nums1[m-1]
			m--
			flg = false
		}
		if flg {
			i--
		}
		k--
	}
	for j := 0; j <= i; j++ {
		nums1[j] = nums2[j]
	}
}
