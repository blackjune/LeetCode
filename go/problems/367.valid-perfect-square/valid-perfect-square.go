package square

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	if num == 1 || num == 4 {
		return true
	}
	if num < 4 {
		return false
	}
	up := num / 2
	down := 2

	for down+1 < up {
		trial := (down + up) / 2
		fmt.Println(up, down, trial)
		if trial*trial > num {
			up = trial
		} else if trial*trial < num {
			down = trial
		} else {
			return true
		}
	}
	return false
}
