package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"bufio"
)

func main()  {
	r := bufio.NewReader(os.Stdin) // taking input
	values, _ := r.ReadString('\n')  // reading input to /n
	// trims space and separates the nums with space inbetween
	parts := strings.Fields(strings.TrimSpace(values))
	nums := make([]int, 0, len(parts)) // making an empty slice(string)
	// for loop to convert the input from strings to int
	for _, p := range parts{
		n, _ := strconv.Atoi(p)
		nums = append(nums, n)
	}
  maximum := nums[0] // seeds the first num from nums this fixs neg nums
	for _, n := range nums{
		if n > maximum {
			maximum = n
		}
	}
	fmt.Println(maximum)
}
