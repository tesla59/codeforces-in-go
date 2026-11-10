// https://codeforces.com/problemset/problem/254/A

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
	in  *bufio.Scanner
	out *bufio.Writer
	_   = slices.Sort[[]int]
	_   = cmp.Compare[int]
)

func init() {
	input, _ := os.Open("input.txt")
	in = bufio.NewScanner(input)
	output, _ := os.Create("output.txt")
	out = bufio.NewWriter(output)
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

type IntPair struct {
	First, Second int
}

func solve() {
	n := readInt()
	cards := make([]IntPair, 2*n)
	for i := range cards {
		num := readInt()
		cards[i] = IntPair{
			First:  num,
			Second: i + 1,
		}
	}
	slices.SortFunc(cards, func(a, b IntPair) int {
		i := cmp.Compare(a.First, b.First)
		if i != 0 {
			return i
		}
		return cmp.Compare(a.Second, b.Second)
	})
	answers := make([]IntPair, 0)
	for i := 0; i < 2*n-1; i += 2 {
		if cards[i].First == cards[i+1].First {
			answers = append(answers, IntPair{cards[i].Second, cards[i+1].Second})
		} else {
			println(-1)
			return
		}
	}
	for i := range answers {
		println(answers[i].First, answers[i].Second)
	}
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
