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
		arr []int
	)
	for {
		fmt.Println("Enter the number to calculate a Fibonacci numbers:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading error!")
			//
			continue
		}
		input = strings.TrimSpace(input)
		num, err = strconv.Atoi(input)
		if err != nil || num < 0 {
			fmt.Println("Enter a valid natural number")
			//reader.Discard(reader.Buffered())
			continue
		}
		break
	}
	if num == 0 {
		arr = []int{0}
	} else if num == 1 {
		arr = []int{0, 1}
	} else {
		arr = make([]int, num+1)
		arr[0] = 0
		arr[1] = 1
		for i := 2; i <= num; i++ {
			arr[i] = arr[i-2] + arr[i-1]
		}
	}
	for i := 0; i <= num; i++ {
		fmt.Printf(" %v", arr[i])
	}
}
