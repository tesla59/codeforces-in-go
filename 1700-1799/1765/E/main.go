package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	MAXINT = int(^uint(0) >> 1)
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	a := readInt()
	b := readInt()
	if a > b {
		print(1)
		return
	}
	print(int(math.Ceil(float64(n) / float64(a))))
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		solve()
	}
}

func readInt() int {
	var x int
	fmt.Fscan(in, &x)
	return x
}

func readStr() string {
	var x string
	fmt.Fscan(in, &x)
	return x
}

func print(x ...any) {
	fmt.Fprintln(out, x...)
}

func printSlice(s []int) {
	fmt.Println(strings.Trim(fmt.Sprint(s), "[]"))
}

func max(x ...int) int {
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] > x[mi] {
			mi = i
		}
	}
	return x[mi]
}

func min(x ...int) int {
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] < x[mi] {
			mi = i
		}
	}
	return x[mi]
}

func GCD(x, y int) int {
	x, y = abs(x), abs(y)
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func LCM(x, y int) int {
	return abs(x*y) / GCD(x, y)
}

func sum(x ...int) int {
	res := 0
	for _, i := range x {
		res += i
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverseInts(x []int) {
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}
}

func reverseStr(s string) string {
	x := []rune(s)
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}
	return string(x)
}
