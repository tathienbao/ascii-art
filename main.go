package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := os.Args
	if len(input) <= 1 {
		fmt.Println("Please enter some text")
		return
	}
	//read file
	standard := "standard.txt"
	lines := readfile(standard)

	if input[1] == "" {
		os.Exit(0)
	}
	for _, value := range input[1] {
		if value > 128 || value < 0 {
			fmt.Println("You can not use non ascii characters!")
			os.Exit(0)
		}
	}
	//convert input to runes
	if input[1] == "\\n" {
		fmt.Println("")
		os.Exit(0)
	}
	runes := ConvertToRune(input[1])
	// nested loop to print line by line depending on input.
	splittwo := "\\n"
	words := strings.Split(string(runes), splittwo)
	for _, word := range words {
		for h := 1; h < 9; h++ {
			if word == "" {
				fmt.Println("")
				break
			}
			for _, l := range []byte(word) {
				for i, line := range lines {
					if i == (int(l)-32)*9+h {
						fmt.Print(line)
					}
				}
			}
			fmt.Println()
		}
	}
}

// function to read file
func readfile(name string) []string {
	//read full file
	file, err := os.Open(name)
	//error check
	if err != nil {
		fmt.Printf("Error message: %s:\n", err)
		os.Exit(2)
	}
	//close file
	defer file.Close()
	//reads full file content to data
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error!")
		os.Exit(2)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

// function to convert string to rune
func ConvertToRune(word string) []rune {
	s := word
	r := []rune(s)
	return r
}
