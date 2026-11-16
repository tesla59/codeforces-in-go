// https://codeforces.com/problemset/problem/2285/C

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

func solve() {
	n := readInt()
	array := make([]int, n)
	maxElement := MinInt
	for i := range n {
		a := readInt()
		array[i] = a
		maxElement = max(maxElement, a)
	}

	count := 0
	for i := range array {
		if array[i] == maxElement {
			count++
		}
		if count > 1 {
			println(1)
			return
		}
	}
	slices.Sort(array)
	MAX := 1_000_000_007
	ans := array[n-1] % MAX
	power := 2
	for i := n - 2; i >= 0; i-- {
		if i > 0 && array[i] == array[i-1] {
			break
		}
		term, b, e := 1, (array[i]-1)%MAX, power
		for e > 0 {
			if e&1 == 1 {
				term = (term * b) % MAX
			}
			b = (b * b) % MAX
			e /= 2
		}
		ans = (ans + term) % MAX
		power++
	}
	println(ans)
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
