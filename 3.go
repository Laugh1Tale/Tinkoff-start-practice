package main

import "fmt"

func main3() {
	var max, min, n, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		if i == 0 {
			min = temp
			continue
		}
		if i == 1 {
			max = temp
			continue
		}
		if i%2 == 0 {
			sum += temp
			if temp < min {
				min = temp
			}
		} else {
			sum -= temp
			if temp > max {
				max = temp
			}
		}
	}
	fmt.Println(sum + max - min)
}
