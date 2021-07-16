// https://atcoder.jp/contests/abs/tasks/abc083_b
package main

import "fmt"

func main() {
	var limit, low, high int
	var total int

	fmt.Scanf("%d %d %d", &limit, &low, &high)

	for i := 1; i <= limit; i++ {
		sum := sumOfDigits(i)

		if sum >= low && sum <= high {
			total += i
		}
	}

	fmt.Println(total)
}

func sumOfDigits(num int) int {
	var sum int

	for num > 0 {
		q := num / 10
		r := num % 10
		sum += r
		num = q
	}

	return sum
}
