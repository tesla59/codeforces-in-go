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
	n, m := readInt(), readInt()
	mat := makeMatrix(n, m)
	for i := range mat {
		str := readStr()
		mat[i] = strings.Split(str, "")
	}
	// print(mat)
	x1, y1 := 0, 0
out:
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == "R" {
				x1 = i
				y1 = j
				break out
			}
		}
	}
	x2, y2 := 0, 0
in:
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[j][i] == "R" {
				x2 = j
				y2 = i
				break in
			}
		}
	}
	if x1 == x2 && y1 == y2 {
		print("YES")
	} else {
		print("NO")
	}
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

func makeMatrix(n, m int) [][]string {
	x := make([][]string, n)
	for i := range x {
		x[i] = make([]string, m)
	}
	return x
}
