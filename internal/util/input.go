package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Get input from the user
func GetInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("$ %s: ", prompt)

	result, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Could not read input: ", err)
	}

	result = strings.TrimSuffix(result, "\n")
	result = cleanSpaces(result)
	return result
}
func MustGetInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("$ %s: ", prompt)

	result, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Could not read input: ", err)
	}

	result = strings.TrimSuffix(result, "\n")
	result = cleanSpaces(result)
	if len(result) < 1 {
		log.Fatal("Expected a value")
	}
	return result
}

func cleanSpaces(s string) string {
	cleanStr := strings.ReplaceAll(s, " ", "_")
	return cleanStr
}
