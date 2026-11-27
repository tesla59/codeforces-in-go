// https://codeforces.com/problemset/problem/2292/D

package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

const (
	MaxInt = math.MaxInt
	MinInt = math.MinInt
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Real interface {
	Integer | ~float32 | ~float64
}

var (
	in  = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	_   = slices.Sort[[]int]
	_   = cmp.Compare[int]
)

func init() {
	in.Split(bufio.ScanWords)
	const maxBuf = 1024 * 1024
	in.Buffer(make([]byte, maxBuf), maxBuf)
}

func readInt() int {
	in.Scan()
	x, _ := strconv.Atoi(in.Text())
	return x
}

func readInt64() int64 {
	in.Scan()
	x, _ := strconv.ParseInt(in.Text(), 10, 64)
	return x
}

func readStr() string {
	in.Scan()
	return in.Text()
}

func readSlice(n int) []int {
	v := make([]int, n)
	for i := range n {
		v[i] = readInt()
	}
	return v
}

func print(args ...any) {
	fmt.Fprint(out, args...)
}

func println(args ...any) {
	fmt.Fprintln(out, args...)
}

func printSlice[T any](s []T) {
	for i, v := range s {
		if i > 0 {
			out.WriteByte(' ')
		}
		fmt.Fprint(out, v)
	}
	out.WriteByte('\n')
}

func abs[T Real](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func gcd[T Integer](a, b T) T {
	a, b = abs(a), abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm[T Integer](a, b T) T {
	return abs(a / gcd(a, b) * b)
}

func dp(n int, x, y []int, sunnyOnLeft bool, memo [][2]int, visited [][2]bool) int {
	if n >= len(x) {
		return 0
	}
	side := 0
	if !sunnyOnLeft {
		side = 1
	}
	if visited[n][side] {
		return memo[n][side]
	}
	buy, skip, ans := 0, 0, 0
	if sunnyOnLeft {
		buy = x[n] - y[n] + dp(n+1, x, y, true, memo, visited)
		skip = dp(n+1, x, y, false, memo, visited)
		ans = max(buy, skip)
	} else {
		buy = y[n] - x[n] + dp(n+1, x, y, false, memo, visited)
		skip = dp(n+1, x, y, true, memo, visited)
		ans = min(buy, skip)
	}
	visited[n][side] = true
	memo[n][side] = ans
	return memo[n][side]
}

func solve() {
	n := readInt()
	x, y := readSlice(n), readSlice(n)
	memo := make([][2]int, n)
	visited := make([][2]bool, n)
	d := dp(0, x, y, true, memo, visited)
	if d > 0 {
		println("Sunny")
	} else if d < 0 {
		println("Aubrey")
	} else {
		println("Draw")
	}
}

func main() {
	defer out.Flush()
	t := 1
	// if multiple test cases:
	t = readInt()
	for ; t > 0; t-- {
		solve()
	}
}
