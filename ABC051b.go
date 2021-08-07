package main

import "fmt"

func main() {
	var upto, target int
	fmt.Scanf("%d %d", &upto, &target)

	var result int
	for i := 0; i <= upto; i++ {
		for j := 0; j <= upto; j++ {

			remaining := target - (i + j)

			if remaining >= 0 && remaining <= upto {
				result++
			}
		}
	}

	fmt.Println(result)
}
