//https://www.hackerrank.com/challenges/the-power-sum/problem
package main

import (
	"math"

	"fmt"
)

var count = 0

func main() {
	var x, n int
	fmt.Scan(&x, &n)

	limit := int(math.Ceil(math.Pow(float64(x), 1/float64(n))))

	for i := 1; i <= limit; i++ {
		sum := int(math.Pow(float64(i), float64(n)))
		calc(i+1, sum, x, n, limit)
	}
	fmt.Println(count)
}

func calc(num, curSum, x, n, limit int) {
	if curSum < x {
		for i := num; i <= limit; i++ {
			calc(i+1, curSum+int(math.Pow(float64(i), float64(n))), x, n, limit)
		}
	} else if curSum == x {
		count += 1
	}
}
