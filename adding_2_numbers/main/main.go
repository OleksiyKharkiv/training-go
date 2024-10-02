package main

import (
	"bufio"
	"fmt"
	"os"
)
import _ "fmt"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert firs number:")
	input1, _ := reader.ReadString('\n')
}
