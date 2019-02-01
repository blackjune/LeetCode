package prefix

func longestCommonPrefix(strs []string) (r string) {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return
	}
	var curr string
	for i := 0; i < len(strs[0]); i++ {
		curr = strs[0][i : i+1]
		for _, str := range strs {
			if i >= len(str) || str[i:i+1] != curr {
				return
			}
		}
		r += curr
	}
	return
}
