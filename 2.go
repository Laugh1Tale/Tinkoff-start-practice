package main

import (
	"fmt"
	"sort"
)

func main2() {
	var n, max int
	fmt.Scanf("%d\n", &n)
	winners := make(map[[3]string]int, n)
	for i := 0; i < n; i++ {
		var first, second, third string
		fmt.Scanf("%s %s %s\n", &first, &second, &third)
		temp := [3]string{first, second, third}
		sort.Strings(temp[0:3])
		winners[temp] += 1
		if max < winners[temp] {
			max = winners[temp]
		}
	}
	fmt.Println(max)
}
