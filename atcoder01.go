// https://atcoder.jp/contests/abs/tasks/abc086_a

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	product := a * b

	var output string
	if product^1 == product+1 {
		output = "Even"
	} else {
		output = "Odd"
	}

	fmt.Println(output)
}
