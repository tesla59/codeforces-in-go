package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	MAXINT = int(^uint(0) >> 1)
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	_, arr := readArr(n)
	_, arrdash := readArr(n)
	l, r := -1, -1
	for i := 0; i < n; i++ {
		if arr[i] != arrdash[i] {
			l = i
			break
		}
	}
	for i := n - 1; i > 0; i-- {
		if arr[i] != arrdash[i] {
			r = i
			break
		}
	}
	// if l != 0 && arrdash[l] == arrdash[l-1] {
	// 	l--
	// }
	// if r != n-1 && arrdash[r] == arrdash[r+1] {
	// 	r++
	// }
	// print(l, r)
	for l > 0 {
		if arrdash[l] >= arrdash[l-1] {
			l--
			continue
		}
		break
	}
	for r < n-1 {
		if arrdash[r] <= arrdash[r+1] {
			r++
			continue
		}
		break
	}
	print(l+1, r+1)
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
	fmt.Fprintln(out, strings.Trim(fmt.Sprint(s), "[]"))
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

func binarySearch(arr []int, target int) int {
	index := sort.Search(len(arr), func(i int) bool { return arr[i] >= target })
	if index < len(arr) && arr[index] == target {
		return index
	}
	return -1
}
