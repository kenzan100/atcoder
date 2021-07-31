// https://atcoder.jp/contests/abs/tasks/arc089_a

package main

import (
	"fmt"
)

func main() {
	var row_len int
	fmt.Scanln(&row_len)

	var current_x, current_y, current_time int
	result := true

	for i := 0; i < row_len; i++ {
		var t, x, y int
		fmt.Scanf("%d %d %d", &t, &x, &y)

		success := solve(current_x, current_y, t-current_time, x, y)
		if !success {
			result = false
		}

		current_x = x
		current_y = y
		current_time = t
	}

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func solve(x0, y0, time, x1, y1 int) bool {
	steps := abs(x1-x0) + abs(y1-y0)

	if steps > time {
		return false
	} else if steps == time {
		return true
	}

	remaining := time - steps
	if remaining%2 == 0 {
		return true
	} else {
		return false
	}
}
