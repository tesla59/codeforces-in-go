// http://codeforces.com/problemset/problem/2288/B1

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

func dp(i, p1, p2 int, spell []byte, memo [][105][105]int) int {
	if i >= len(spell) || (p1 == 0 && p2 == 0) {
		return p1 + p2
	}
	if memo[i][p1][p2] != -1 {
		return memo[i][p1][p2]
	}
	x1, x2 := 0, 0
	if spell[i] == '+' {
		// on 1
		if p1 > 0 {
			x1 = dp(i+1, p1-1, p2+1, spell, memo)
		}
		// on 2
		if p2 > 0 {
			x2 = dp(i+1, p1, p2+1, spell, memo)
		}
	} else if spell[i] == '-' {
		if p1 > 0 {
			x1 = dp(i+1, p1-1, p2, spell, memo)
		}
		if p2 > 0 {
			x2 = dp(i+1, p1+1, p2-1, spell, memo)
		}
	}
	memo[i][p1][p2] = max(x1, x2, p1+p2)
	return memo[i][p1][p2]
}

func solve() {
	p1, p2, m := readInt(), readInt(), readInt()
	spell := []byte(readStr())
	memo := make([][105][105]int, m+1)
	for i := range memo {
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	println(dp(0, p1, p2, spell, memo))
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
