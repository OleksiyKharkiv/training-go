package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	n, num int
	check  bool
)

func generateRandomArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(101)
	}
	return arr
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter the number from 1 to 100 of elements of array")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			reader.Discard(reader.Buffered())
			continue
		}
		input = strings.TrimSpace(input)
		n, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter a valid numer of elements of array from 1 to 100")
			continue
		}
		break
	}
	randomArr := generateRandomArray(n)
	for {
		fmt.Println("Enter the number you want to find in this array")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			reader.Discard(reader.Buffered())
			continue
		}
		input = strings.TrimSpace(input)
		num, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter a valid numer of elements of array from 1 to 100")
			continue
		}
		break
	}
	for i := 0; i < len(randomArr); i++ {
		if randomArr[i] == num {
			fmt.Printf(" The number %v (you selected) is present in the %v - position in the array \n", num, i)
			check = true
			break
		}
	}
	if check == false {
		fmt.Printf("The number %v is absent in the array \n", num)
	}
	fmt.Println(" The array is : ")
	for i := 0; i < len(randomArr); i++ {
		fmt.Printf("%v ", randomArr[i])
	}
}
