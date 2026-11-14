// https://codeforces.com/problemset/problem/189/A

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

// func dp(n, a, b, c, ans int) int {
// 	if n <= 0 {
// 		return ans
// 	}
// 	A, B, C := 0, 0, 0
// 	if n-a >= 0 {
// 		A = dp(n-a, a, b, c, ans+1)
// 	}
// 	if n-b >= 0 {
// 		B = dp(n-b, a, b, c, ans+1)
// 	}
// 	if n-c >= 0 {
// 		C = dp(n-c, a, b, c, ans+1)
// 	}

// 	return max(A, B, C)
// }

func dp(n, a, b, c int, memo []int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -MaxInt
	}
	if memo[n] != -1 {
		return memo[n]
	}
	ans := max(dp(n-a, a, b, c, memo), dp(n-b, a, b, c, memo), dp(n-c, a, b, c, memo))
	if ans == -MaxInt {
		memo[n] = -MaxInt
		return -MaxInt
	}
	memo[n] = 1 + ans
	return 1 + ans
}

func solve() {
	n, a, b, c := readInt(), readInt(), readInt(), readInt()
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	println(dp(n, a, b, c, memo))
}

func main() {
	defer out.Flush()
	t := 1
	// if multiple test cases:
	// t = readInt()
	for ; t > 0; t-- {
		solve()
	}
}
