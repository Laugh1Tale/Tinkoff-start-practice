package main

import (
	"fmt"
	"math"
)

func main1() {
	var x1, y1, x2, y2, x11, y11, x22, y22, a, b, max float64
	fmt.Scanf("%f %f %f %f\n%f %f %f %f", &x1, &y1, &x2, &y2, &x11, &y11, &x22, &y22)
	a = math.Max(math.Abs(x22-x1), math.Abs(x2-x11))
	b = math.Max(math.Abs(y22-y1), math.Abs(y2-y11))
	max = math.Max(a, b)
	fmt.Println(max * max)
}
