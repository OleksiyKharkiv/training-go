package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var (
		num1, num2, res float64
		symb            rune
	)
	fmt.Println("Welcome to Troll-Calculator!")
	time.Sleep(2000 * time.Millisecond)

	for {
		// Entering the first argument
		fmt.Println("Please enter the first number and press enter:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			reader.Discard(reader.Buffered())
			continue
		}
		input = strings.TrimSpace(input)
		num1, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Please enter a valid number!")
			continue
		}
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("Oh!!! So beatiful number....!!!... ")
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("Please wait... Let me enjoy this number!")
		time.Sleep(2000 * time.Millisecond)
		break
	}
	for {
		// Entering the action
		fmt.Println("Ok... Now... Please enter the mathematical operation : '+' or '-' or '*' or '/' and press enter")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			reader.Discard(reader.Buffered())
			continue
		}
		input = strings.TrimSpace(input)
		if len(input) != 1 {
			fmt.Println("Please enter a valid symbol operation !")
			continue
		}
		symb = rune(input[0])
		if symb != '*' && symb != '+' && symb != '-' && symb != '/' {
			fmt.Println("Please enter a valid symbol operation !")
			continue
		}
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(".. and you are such an inventor!...   ;)      ")
		time.Sleep(2000 * time.Millisecond)
		fmt.Printf("Please wait... Let me enjoy this beaty symbol... wow!  %c !!!   are you kidding me?!...", symb)
		fmt.Println()
		time.Sleep(2000 * time.Millisecond)
		break
	}
	for {
		// Entering the second argument
		fmt.Println("Okay-Okay I'll try to do this, I just need to take counting sticks...  Please enter the second number and press enter:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			reader.Discard(reader.Buffered())
			continue
		}
		input = strings.TrimSpace(input)
		num2, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Please enter a valid number!")
			continue
		}
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(".. are you by any chance the grandson of Pythagoras?!...    ")
		time.Sleep(2000 * time.Millisecond)
		fmt.Printf("Wait, I need to take some sedatives.  %v !!!   are you kidding me?!...", num2)
		fmt.Println()
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(".. I not sure ... but maybe this...    ")
		time.Sleep(2000 * time.Millisecond)
		fmt.Println()
		break
	}
	switch symb {
	case '*':
		res = num1 * num2
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '/':
		res = num1 / num2
	}
	fmt.Printf("The result of your operation is : %v %c %v = %v\n", num1, symb, num2, res)
	time.Sleep(5000 * time.Millisecond)
}
