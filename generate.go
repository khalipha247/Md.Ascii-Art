package main

import (
	"fmt"
)

func GenerateArt(text string, banner map[rune][]string) {
	for row := range 8 {
		for _, r := range text {
			art := GetChar(r, banner)
			fmt.Print(art[row])
		}
		fmt.Println()
	}
}
