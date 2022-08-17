package main

import (
	"fmt"
	"sort"
	"strings"
)

func main5() {
	var n, q int
	cards := make(map[string]int)
	fastCards := make(map[byte][]string)
	fmt.Scan(&n, &q)
	for i := 0; i < n; i++ {
		var input string
		fmt.Scan(&input)
		fastCards[input[0]] = append(fastCards[input[0]], input)
		cards[input] = i + 1
	}
	for key := range fastCards {
		sort.Strings(fastCards[key])
	}
	for i := 0; i < q; i++ {
		var number, counter int
		request := -1
		var prefix string
		fmt.Scan(&number, &prefix)
		for _, value := range fastCards[prefix[0]] {
			if strings.HasPrefix(value, prefix) {
				counter++
				if counter == number {
					request = cards[value]
				}
			}
		}
		fmt.Println(request)
	}
}
