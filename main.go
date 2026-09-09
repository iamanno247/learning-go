package main 

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main()  {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	// Print the upper case version
  line = strings.ToUpper(line)
	fmt.Println(line)
}
