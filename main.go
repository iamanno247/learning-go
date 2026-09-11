package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"bufio"
)

func main()  {
	r := bufio.NewReader(os.Stdin)
	numStr, _ := r.ReadString('\n')
	numStr = strings.TrimRight(numStr, "\r\n")
	num, _ := strconv.Atoi(numStr)
	squared := square(num)
	fmt.Println(squared)
}

func square(n int) int  {
	return n * n
}
