// https://atcoder.jp/contests/abs/tasks/abc081_a
package main

import "fmt"

func main() {
	// get them as a number n

	var count, num int
	// bit shift
	fmt.Scanf("%b", &num)

	for num > 0 {
		count += num & 1
		num = num >> 1
	}

	fmt.Println(count)
}
