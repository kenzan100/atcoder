// https://atcoder.jp/contests/abs/tasks/abc081_b
package main

import "fmt"

func main() {
	var cnt int
	fmt.Scanf("%d", &cnt)

	var nums = make([]int, cnt)
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	var numOps int
out:
	for {
		for i, num := range nums {
			// if modulo is other than 0, break
			if num%2 > 0 {
				break out
			}

			// otherwise, divide them
			nums[i] = num / 2
		}
		numOps++
	}

	fmt.Println(numOps)
}
