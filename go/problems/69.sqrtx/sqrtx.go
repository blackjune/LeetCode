package sqrtx

import (
	"fmt"
)

func mySqrt(x int) int {
	for i := 0; i < x; i++ {
		fmt.Println(i)
		if i*i > x {
			return i - 1
		}
		if i*i == x {
			return i
		}
	}
	return 1
}
