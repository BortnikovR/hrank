//https://www.hackerrank.com/challenges/password-cracker/problem
package main

import (
	"fmt"
)

var found = false
var memo = map[string]struct{}{}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		//crack()
		crack2()
	}
}

func crack2() {
	var n int
	fmt.Scan(&n)
	passwords := map[string]struct{}{}
	var pass string
	minLen, maxLen := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&pass)

		l := len(pass)
		if l < minLen || minLen == 0 {
			minLen = l
		}
		if l > maxLen {
			maxLen = l
		}

		passwords[pass] = struct{}{}
	}

	var testPass string
	fmt.Scan(&testPass)

	memo = map[string]struct{}{}
	found = false
	s, ok := findSolution2(passwords, testPass, "", minLen, maxLen)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("WRONG PASSWORD")
	}
}

func findSolution2(passwords map[string]struct{}, testPass, solution string, minLen, maxLen int) (string, bool) {
	if found {
		return "", true
	}

	if len(testPass) == 0 {
		found = true
		return solution, true
	}

	if _, ok := memo[testPass]; ok {
		return "", false
	}

	for i := minLen; i <= maxLen; i++ {
		if i > len(testPass) {
			return "", false
		}

		if _, ok := passwords[testPass[0:i]]; ok {
			memo[testPass] = struct{}{}
			if s, ok := findSolution2(passwords, testPass[i:], solution+testPass[0:i]+" ", minLen, maxLen); ok {
				return s, true
			}
		}
	}

	return "", false
}
