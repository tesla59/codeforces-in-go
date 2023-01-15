package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		t int
		n int
		s string
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		fmt.Fscanf(reader, "%d\n", &n)
		fmt.Fscanf(reader, "%s\n", &s)
		answer := "NO"

		for i := 0; i < n - 1; i++ {
			text := string(s[i]) + string(s[i+1])
			if strings.Count(s, text) > 1 {
				answer = "YES"
				break
			}
		}
		fmt.Println(answer)
	}
}
