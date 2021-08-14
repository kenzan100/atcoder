// https://atcoder.jp/contests/abc126/tasks/abc126_a

package main

import (
	"fmt"
	"strings"
)

func main() {
	var size, idx int
	fmt.Scanf("%d %d", &size, &idx)

	var str string
	fmt.Scan(&str)

	strArr := strings.Split(str, "")

	strArr[idx-1] = strings.ToLower(strArr[idx-1])

	fmt.Println(strings.Join(strArr, ""))
}
