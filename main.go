package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)


func main()  {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	words := strings.Fields(strings.TrimSpace(line))
	mappedWords := make(map[string]bool)
	for _, words := range words{
		mappedWords[words] = true
	}
	fmt.Println(len(mappedWords))
}
