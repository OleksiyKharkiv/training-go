package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"training-go/sorting_structures/internal/domain/entity"
	"training-go/sorting_structures/internal/domain/service"
)

// CLIHandler handles command-line interface interactions.
type CLIHandler struct {
	sorterService *service.SorterServiceImpl
}

// NewCLIHandler creates a new instance of CLIHandler.
func NewCLIHandler() *CLIHandler {
	return &CLIHandler{
		sorterService: &service.SorterServiceImpl{},
	}
}

// Run starts the CLI Handler for sorting people.
func (h *CLIHandler) Run() error {
	reader := bufio.NewReader(os.Stdin)
	var people []entity.Person
	numbUsers := 0

	// Input number of users
	for {
		fmt.Println("Enter the number of users (from 1 to 10):")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			continue
		}
		input = strings.TrimSpace(input)
		numbUsers, err = strconv.Atoi(input)
		if err != nil || numbUsers < 1 || numbUsers > 10 {
			fmt.Println("Please enter a valid number of users (between 1 and 10)!")
			continue
		}
		break
	}

	// Input name and age for each user
	for i := 0; i < numbUsers; i++ {
		var name string
		var age int

		// Input name
		for {
			fmt.Printf("Please enter name of user %d (not empty): ", i+1)
			name, _ = reader.ReadString('\n')
			name = strings.TrimSpace(name)
			if len(name) != 0 {
				break
			}
			fmt.Println("Name cannot be empty!")
		}

		// Input age
		for {
			fmt.Printf("Please enter age for user %d: ", i+1)
			ageInput, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input")
				continue
			}
			ageInput = strings.TrimSpace(ageInput)
			age, err = strconv.Atoi(ageInput)
			if err != nil || age < 1 || age > 120 {
				fmt.Println("Please enter a valid age!")
				continue
			}
			break
		}

		// Add person to the list
		people = append(people, entity.Person{Name: name, Age: int8(age)})
	}

	// Sort people by age
	sortedPeople := h.sorterService.Sort(people)
	fmt.Println("\nSorted by age:")
	for _, person := range sortedPeople {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}

	return nil
}
