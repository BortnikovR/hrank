//https://www.hackerrank.com/challenges/arithmetic-expressions/problem
package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	divider  = 101
	max      = math.MaxInt64 - 89
	min      = -max
	minus    = "-"
	plus     = "+"
	asterisk = "*"
)

var (
	operators = []string{"+", "-", "*"}
	memo      = map[string]struct{}{}
	numbers   []int64
)

func main() {
	n := 0
	fmt.Scan(&n)

	numbers = make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	if res, ok := buildExpression(0, 0, ""); ok {
		fmt.Println(res)
		//spew.Dump(memo)
		//spew.Dump(counter)
	} else {
		panic("second WTF?")
	}
}

func buildExpression(idx int, cur int64, expr string) (string, bool) {
	if idx != 0 && cur%divider == 0 {
		return multipleTheRest(expr, idx), true
	}

	key := fmt.Sprintf("%d_%d", idx, cur)
	if _, f := memo[key]; f {
		return "", false
	}

	if idx == len(numbers) {
		return expr, cur%divider == 0
	}

	for _, s := range operators {
		newCur, _ := calcCur(cur, numbers[idx], s)
		key = fmt.Sprintf("%d_%d", idx, cur)
		memo[key] = struct{}{}
		newExpr := expr
		if expr != "" {
			newExpr = expr + s
		}
		newExpr += strconv.FormatInt(numbers[idx], 10)
		if res, ok := buildExpression(idx+1, newCur%divider, newExpr); ok {
			return res, ok
		}
	}

	return "", false
}

func multipleTheRest(expr string, idx int) string {
	for i := idx; i < len(numbers); i++ {
		expr += "*" + strconv.FormatInt(numbers[i], 10)
	}
	return expr
}

func calcCur(cur int64, num int64, sign string) (int64, bool) {
	//if sign == plus && num+cur < max {
	//	return num + cur, true
	//} else if sign == asterisk && num*cur < max {
	//	return num * cur, true
	//} else if sign == minus && cur-num > min {
	//	return cur - num, true
	//}
	//return 0, false
	switch sign {
	case plus:
		return cur + num, true
	case minus:
		return cur - num, true
	case asterisk:
		return cur * num, true
	default:
		panic("WTF?")
	}
}
