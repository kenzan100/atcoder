// https://atcoder.jp/contests/abs/tasks/abc085_b

package main

import "fmt"

func main() {

	var len int
	fmt.Scanf("%d", &len)

	var nums = make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Scan(&nums[i])
	}

	memo := make(map[int]bool)
	var stack int

	for _, num := range nums {
		if memo[num] {
			continue
		} else {
			stack += 1
			memo[num] = true
		}
	}

	fmt.Println(stack)
}
