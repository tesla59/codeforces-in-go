// https://codeforces.com/problemset/problem/431/B

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

func dp(i int, arr [5]int, ans [][5]int) [][5]int {
	if i == 5 {
		ans = append(ans, arr)
		return ans
	}

	for j := i; j < len(arr); j++ {
		arr[j], arr[i] = arr[i], arr[j]
		ans = dp(i+1, arr, ans)
		arr[j], arr[i] = arr[i], arr[j]
	}
	return ans
}

func solve() {
	matrix := [5][5]int{}
	for i := range 5 {
		for j := range 5 {
			matrix[i][j] = readInt()
		}
	}
	permutations := [][5]int{}
	permutations = dp(0, [5]int{1, 2, 3, 4, 5}, permutations)

	ans := MinInt
	for _, permute := range permutations {
		happiness := matrix[permute[0]-1][permute[1]-1] + matrix[permute[1]-1][permute[0]-1] +
			matrix[permute[1]-1][permute[2]-1] + matrix[permute[2]-1][permute[1]-1] +
			2*matrix[permute[2]-1][permute[3]-1] + 2*matrix[permute[3]-1][permute[2]-1] +
			2*matrix[permute[3]-1][permute[4]-1] + 2*matrix[permute[4]-1][permute[3]-1]
		ans = max(happiness, ans)
	}
	println(ans)
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
