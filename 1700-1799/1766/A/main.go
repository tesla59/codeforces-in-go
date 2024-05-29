package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		t int
		n int
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		fmt.Fscanf(reader, "%d\n", &n)
		str :=  strconv.Itoa(n)
		temp := strings.Split(str, "")
		x, _ := strconv.Atoi(temp[0])
		fmt.Println(x + (len(temp)-1) * 9)
	}
}
