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
	return abs(a/gcd(a, b) * b)
}

func reverseSlice[T any](s []T) {
	slices.Reverse(s)
}

func sortSlice[T cmp.Ordered](s []T) {
	slices.Sort(s)
}

func binarySearch[T cmp.Ordered](s []T, target T) int {
	idx, found := slices.BinarySearch(s, target)
	if found {
		return idx
	}
	return -1
}

type intPair struct {
	first int
	second int
}

func solve() {
	s, n := readInt(), readInt()
	var stages []intPair
	for range n {
		x, y := readInt(), readInt()
		stages = append(stages, intPair{first: x, second: y})
	}

	slices.SortFunc(stages, func(a, b intPair) int {
		n := cmp.Compare(a.first, b.first)
		if n != 0 {
			return n
		}
		return cmp.Compare(a.second, b.second)
	})
	for _, v := range stages {
		if s <= v.first {
			print("NO")
			return
		}
		s += v.second
	}
	print("YES")
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
