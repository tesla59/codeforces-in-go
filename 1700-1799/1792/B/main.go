package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	MAXINT = int(^uint(0) >> 1)
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	a1, a2, a3, a4 := readInt(), readInt(), readInt(), readInt()
	if a1 == 0 {
		print(1)
		return
	}

	temp := min(a2, a3)

	ans := a1 + 2*temp
	a2 -= temp
	a3 -= temp

	x := max(a2, a3) + a4
	ans += min(a1, x)
	x -= min(a1, x)

	if x != 0 {
		ans++
	}
	print(ans)
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

func readArr(x ...int) (int, []int) {
	var n int
	if len(x) == 0 {
		fmt.Fscan(in, &n)
	} else {
		n = x[0]
	}
	var v = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
	}
	return n, v
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
