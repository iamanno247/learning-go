package main

import "fmt"

func increment(n *int) {
    // dereference and increment
		*n++
}

func main() {
    var x int
    fmt.Scan(&x)
    increment(&x)
    fmt.Println(x)
}
