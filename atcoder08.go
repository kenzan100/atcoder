// https://atcoder.jp/contests/abs/tasks/abc085_c

package main

import "fmt"

func main() {
	var numBills, totalYen int
	fmt.Scanf("%d %d", &numBills, &totalYen)

	// i - 10_000 bill
	// j -  5_000 bill
	for i := 0; i <= numBills; i++ {

		for j := 0; j <= numBills; j++ {

			remaining := numBills - (i + j)

			if remaining < 0 {
				continue
			}

			sum := i*10_000 + j*5_000 + remaining*1_000

			if sum == totalYen {
				fmt.Printf("%d %d %d\n", i, j, remaining)
				return
			}
		}
	}

	fmt.Println("-1 -1 -1")
}
