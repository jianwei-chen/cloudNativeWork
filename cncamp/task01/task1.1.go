package main

import "fmt"

func main() {
	replace()
}

func replace() {
	oriString := [...]string{"i", "am", "stupid", "and", "work"}

	for index, value := range oriString {
		fmt.Print(value, " ")
		if index == 2 {
			oriString[index] = "smart"
		}
		if index == 4 {
			oriString[index] = "strong"
		}
	}
	fmt.Println()
	for index := range oriString {
		fmt.Print(oriString[index], " ")
	}
}
