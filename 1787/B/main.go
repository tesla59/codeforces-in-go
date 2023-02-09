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
	n := readInt()
	factors := primeFactors(n)
	mp := make(map[int]int)
	for i := range factors {
		mp[factors[i]]++
	}
	sum := 0
	// a, p := make([]int, 0), make([]int, 0)
	MapMin := minMap(mp)
	for MapMin != -1 {
		prod := 1
		for k, v := range mp {
			prod *= k
			// mp[k] -= MapMin
			if v == MapMin {
				delete(mp, k)
			} else {
				mp[k] -= MapMin
			}
		}
		// a = append(a, prod)
		// p = append(p, MapMin)
		sum += prod * MapMin
		MapMin = minMap(mp)
	}
	// for i := range a {
	// 	sum += a[i] * p[i]
	// }
	print(sum)
}

func minMap(mp map[int]int) int {
	if len(mp) == 0 {
		return -1
	}
	ans := MAXINT
	for _, v := range mp {
		ans = min(ans, v)
	}
	return ans
}

func primeFactors(n int) []int {
	factors := []int{}
	for i := 2; i <= n/i; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
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
