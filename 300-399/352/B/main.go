// https://codeforces.com/problemset/problem/352/B

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

type IntPair struct {
	First  int
	Second int
}

func solve() {
	n := readInt()
	numbers := readSlice(n)
	pos := make(map[int][]int)
	for i := range numbers {
		pos[numbers[i]] = append(pos[numbers[i]], i)
	}

	slices.Sort(numbers)
	numbers = slices.Compact(numbers)

	answer := make([]IntPair, 0)
	for _, num := range numbers {
		diff, ok := 0, true
		if len(pos[num]) == 1 {
			answer = append(answer, IntPair{First: num, Second: 0})
		} else if len(pos[num]) == 2 {
			answer = append(answer, IntPair{First: num, Second: pos[num][1] - pos[num][0]})
		} else {
			diff = pos[num][1] - pos[num][0]
			for j := 1; j < len(pos[num]); j++ {
				if pos[num][j]-pos[num][j-1] != diff {
					ok = false
					break
				}
			}
			if ok {
				answer = append(answer, IntPair{First: num, Second: diff})
			}
		}
	}
	println(len(answer))
	for i := range answer {
		println(answer[i].First, answer[i].Second)
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
