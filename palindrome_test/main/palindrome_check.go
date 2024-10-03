package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	runes []rune
	check bool = true
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter a string to check for palindrome (more than 1 symbol): \n")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			reader.Discard(reader.Buffered())
			continue
		}
		input = strings.TrimSpace(input)
		runes = []rune(input)
		if len(runes) < 2 {
			continue
		}
		break
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			check = false
		}
	}
	if check == true {
		fmt.Println("Cogratulation! The entered string is indeed a palindrome!  \n")
	} else {
		fmt.Println("The entered string is NOT a palindrome!  \n")
	}
}
