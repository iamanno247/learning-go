package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
	"bufio"
)

func main()  {
	r := bufio.NewReader(os.Stdin)
	numstr, _ := r.ReadString('\n')
	numstr = strings.TrimRight(numstr, "\r\n")
	num, _ := strconv.Atoi(numstr)
	switch {
		case num % 3 == 0 && num % 5 == 0:
			fmt.Println("FizzBuzz")
		case num % 3 == 0:
			fmt.Println("Fizz")
		case num % 5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(num)
	}
}
