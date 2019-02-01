package binary

func addBinary(a string, b string) string {
	var c, long, short string
	var ll, ls int
	la := len(a)
	lb := len(b)
	if la > lb {
		long, short = a, b
		ll, ls = la, lb
	} else {
		long, short = b, a
		ll, ls = lb, la
	}
	nums := map[byte]int{
		'0': 0,
		'1': 1,
	}
	chars := map[int]string{
		0: "0",
		1: "1",
	}
	var remander int
	for i := 0; i < ll; i++ {
		var sum int
		if i < len(short) {
			sum = remander + nums[long[ll-1-i]] + nums[short[ls-1-i]]
		} else {
			sum = remander + nums[long[ll-1-i]]
		}
		remander = sum / 2
		c = chars[sum%2] + c
	}
	if remander > 0 {
		c = chars[remander] + c
	}
	return c
}
