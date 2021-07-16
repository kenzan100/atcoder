// https://atcoder.jp/contests/abs/tasks/abc087_b

package main

import "fmt"

func main() {
	var target, cnt int
	var a500, b100, c50 int

	fmt.Scanf("%d", &a500)
	fmt.Scanf("%d", &b100)
	fmt.Scanf("%d", &c50)
	fmt.Scanf("%d", &target)

	for i := a500; i >= 0; i-- {
		for j := b100; j >= 0; j-- {
			for k := c50; k >= 0; k-- {
				sum := i*500 + j*100 + k*50
				if sum == target {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}
