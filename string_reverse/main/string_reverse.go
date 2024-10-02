package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var runes []rune
	for {
		fmt.Println("Enter the string to reverse:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input!")
			reader.Discard(reader.Buffered())
			continue
		}
		input = strings.TrimSpace(input)
		runes = []rune(input)
		break
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversed := string(runes)
	fmt.Println("Revesed string: ", reversed)
}
