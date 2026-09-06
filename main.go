package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    // Print a + b.
	sum := float64(a) + float64(b)
	fmt.Print(sum)
 }

