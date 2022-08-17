package main

import "fmt"

func test() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	if d <= b {
		fmt.Println(a)
	} else {
		fmt.Println(a + c*(d-b))
	}
}
