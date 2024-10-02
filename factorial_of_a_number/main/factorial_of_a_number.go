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
	var num, fact int
	for {
		fmt.Println("Enter the natural number to calculate the factorial of this number:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input!")
			reader.Discard(reader.Buffered())
			continue
		}
		input = strings.TrimSpace(input)
		num, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Incorrect input! Please enter a valid number! ")
			reader.Discard(reader.Buffered())
			continue
		}
		if num < 0 {
			fmt.Println("Incorrect input! Please enter a natural number ( 0 or more) ! ")
			reader.Discard(reader.Buffered())
			continue
		}
		break

	}
	if num == 0 || num == 1 {
		fact = 1
	} else {
		fact = 1
		for i := 2; i <= num; i++ {
			fact *= i
		}
	}
	fmt.Printf("The factorial of a number %v is : %v\n", num, fact)
}
