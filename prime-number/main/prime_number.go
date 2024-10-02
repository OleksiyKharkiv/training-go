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
		num int
		res bool = true
	)
	for {
		fmt.Println("Enter the natural number:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input!")
			reader.Discard(reader.Buffered())
			continue
		}
		input = strings.TrimSpace(input)
		num, err = strconv.Atoi(input)
		if err != nil || num < 1 {
			fmt.Println("Incorrect input! Please enter a valid number greater then 0 !")
			reader.Discard(reader.Buffered())
			continue
		}
		break
	}
	if num == 1 || num == 2 || num == 3 {
		res = true
	} else {
		for i := 2; i < num/2+1; i++ {
			if num%i == 0 {
				res = false
			}
		}
	}
	if res == true {
		fmt.Printf("This number (%v) is prime", num)
	} else {
		fmt.Printf("This number (%v) is NOT prime", num)
	}
}
