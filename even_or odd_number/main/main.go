package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var num int

	for {
		fmt.Println("Enter number:")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			reader.Discard(reader.Buffered())
			continue
		}
		input = strings.TrimSpace(input)
		num, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Incorrect input! Please enter a valid  number!")
			reader.Discard(reader.Buffered())
			continue
		}
		break
	}
	if num%2 != 0 {
		fmt.Print("The number is odd")
	} else {
		fmt.Print("The number is even")
	}
}
