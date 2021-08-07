package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d", &size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	fmt.Println(arr)

	var cnt int
	for {
		success := divide_half_all(arr)
		if success {
			cnt++
		} else {
			break
		}
	}

	fmt.Println(cnt)
}

func divide_half_all(arr []int) bool {
	for i, elem := range arr {
		if elem%2 != 0 {
			return false
		}
		arr[i] = elem / 2
	}
	return true
}
