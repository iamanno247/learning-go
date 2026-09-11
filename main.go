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
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("bad")
	} else {
		fmt.Printf("ok %d", num)
	}
}
