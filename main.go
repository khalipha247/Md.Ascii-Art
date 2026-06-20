package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("use: go run . [string] [banner]")
		return
	}
	input := os.Args[1]
	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}
	banner := "standard.txt"

	if len(os.Args) == 3 {
		banner = os.Args[2]
	}
	if !strings.HasSuffix(banner, ".txt") {
		banner += ".txt"
	}

	data, err := LoadBanner(banner)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := strings.Split(input, "\\n")
	for _, char := range text {
		if char == "" {
			fmt.Println()
		} else {
			GenerateArt(char, data)
		}
	}
}
