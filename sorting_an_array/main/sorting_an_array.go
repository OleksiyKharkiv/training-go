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
		arr  [5]int
		temp int
	)
	for i := 0; i < 5; i++ {
		fmt.Printf(" Enter the %v - element of array:", i+1)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(" Error reading input!")
			reader.Discard(reader.Buffered())
			i--
			continue
		}
		input = strings.TrimSpace(input)
		arr[i], err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter a valid integer number! ")
			i--
			continue
		}
		if i > 0 {
			if arr[i] < arr[i-1] {
				temp = arr[i]
				arr[i] = arr[i-1]
				arr[i-1] = temp
			}
		}
	}
	for i := 0; i < 5; i++ {
		fmt.Print("[ %v ],", arr[i])
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		for j := 1; j < 5-i; j++ {
			if arr[j] < arr[j-1] {
				temp = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
	for i := 0; i < 5; i++ {
		fmt.Print("[ %v ],", arr[i])
	}
}
