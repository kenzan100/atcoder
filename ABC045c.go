package main

import (
	"fmt"
	"strconv"
)

// https://shoman.hatenablog.com/entry/2020/02/18/191933
func main() {
	var s string
	fmt.Scan(&s)

	// "125", startIdx
	sum := dfs(s, 0, 0, 0)
	fmt.Println(sum)
}

func dfs(s string, i int, lasti int, sum int) int {
	if i == len(s)-1 {
		last, _ := strconv.Atoi(s[lasti:len(s)])
		return sum + last
	}

	n, _ := strconv.Atoi(s[lasti : i+1])

	sum1 := dfs(s, i+1, lasti, sum)
	sum2 := dfs(s, i+1, i+1, sum+n)

	return sum1 + sum2
}
