//https://www.hackerrank.com/challenges/recursive-digit-sum/problem
package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var (
		n string
		k int
	)
	fmt.Scan(&n, &k)

	var sum int
	for _, i := range strings.Split(n, "") {
		num, _ := strconv.Atoi(i)
		sum += num
	}

	fmt.Println(super(sum * k))
}

func super(sum int) int {
	if sum / 10 == 0 {
		return sum
	}
	newSum := 0
	for newSum = sum % 10; sum / 10 > 0; newSum += sum % 10 {
		sum /= 10
	}
	return super(newSum)
}
