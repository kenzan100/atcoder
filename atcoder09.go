// https://atcoder.jp/contests/abs/tasks/arc065_a
package main

import "fmt"

func main() {
	var word string
	fmt.Scanln(&word)

	dict := map[string][]string{
		"d": {"dream", "dreamer"},
		"e": {"erase", "eraser"},
	}

	result := solve(word, dict)

	if result {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(word string, dict map[string][]string) bool {
	if word == "" {
		return true
	}

	ch := word[0:1]
	candidates := dict[ch]

	for _, dict_word := range candidates {
		rest, success := match(word, dict_word)
		if success {
			if solve(rest, dict) {
				return true
			}
		}
	}
	return false
}

func match(word string, dict_word string) (string, bool) {
	if len(word) < len(dict_word) {
		return "", false
	}

	for i := 0; i < len(dict_word); i++ {
		if word[i] != dict_word[i] {
			return "", false
		}
	}

	return word[len(dict_word):], true
}
