package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
)

func main()  {
	r := bufio.NewReader(os.Stdin)
	numStr, _ := r.ReadString('\n')
	numStr = strings.TrimRight(numStr, "\r\n")
	num, _ := strconv.Atoi(numStr)
	total := 0
	for i := 1; i <= num; i++ {
		total += i
	}
	fmt.Println(total)
}
