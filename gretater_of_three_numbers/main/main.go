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
	var (
		arr       [3]int
		maxNumber int
	)

	fmt.Println("Enter three number one by one , pressing Rnter after each number.")
	for i := 0; i < 3; i++ {
		fmt.Printf("Enter the %v number:\n", i+1)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input", err)
			reader.Discard(reader.Buffered())
			i--
			continue
		}
		input = strings.TrimSpace(input)
		arr[i], err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Incorrect input! Please enter a valid numner!")
			reader.Discard(reader.Buffered())
			i--
			continue
		}
		if i == 0 {
			maxNumber = arr[0]
		}
		if arr[i] > maxNumber {
			maxNumber = arr[i]
		}

	}
	fmt.Printf("The grearer of three number is : %v", maxNumber)
}
