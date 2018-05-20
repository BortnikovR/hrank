//https://www.hackerrank.com/challenges/morgan-and-a-string/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var t int
	fmt.Scan(&t)

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < t; i++ {
		a, _ := reader.ReadString('\n')
		a = strings.Trim(a, "\n")

		b, _ := reader.ReadString('\n')
		b = strings.Trim(b, "\n")

		findString([]rune(a), []rune(b))
		//findString2([]rune(a), []rune(b))
	}
}

func findString2(a, b []rune) {
	res := make([]rune, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		k := i
		for i < len(a) && a[i] < b[j] {
			i++
		}
		res = append(res, a[k:i]...)
		if i == len(a) {
			break
		}

		l := j
		for j < len(b) && a[i] > b[j] {
			j++
		}
		res = append(res, b[l:j]...)
		if j == len(b) {
			break
		}

		if a[i] == b[j] {
			if string(a[i:]) < string(b[j:]) {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
	}

	if i == len(a) {
		res = append(res, b[j:]...)
	} else {
		res = append(res, a[i:]...)
	}
	fmt.Println(string(res))
}

func findString(a, b []rune) {
	res := make([]rune, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else if a[i] > b[j] {
			res = append(res, b[j])
			j++
		} else {
			if string(a[i:]) < string(b[j:]) {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
	}

	if i == len(a) {
		res = append(res, b[j:]...)
	} else {
		res = append(res, a[i:]...)
	}
	fmt.Println(string(res))
}
