// https://atcoder.jp/contests/abs/tasks/abc088_b

package main

import (
	"fmt"
	"sort"
)

func main() {
	var len int
	fmt.Scanf("%d", &len)

	var nums = make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Scan(&nums[i])
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	var alice, bob int
	for i, num := range nums {
		if i%2 == 0 {
			alice += num
		} else {
			bob += num
		}
	}

	fmt.Println(alice - bob)
}
