package main

import "fmt"

//Remove the character from left side
func removeLeftChars(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return ""
}

//Remove the character from right side
func removeRightChars(s string) []string {
	var getSubString []string
	sz := len(s)
	if sz > 0 {
		for i := range s {
			s = s[:sz-i]
			fmt.Println(s)
			getSubString = append(getSubString, s)

		}
	}
	return getSubString
}

func main() {
	for {
		var input string
		var getSubstring []string
		//Read the string from CLI
		fmt.Print("Please Enter string : ")
		fmt.Scanln(&input)
		if input != "" {
			if input == "exit" {
				break
			}
			length := len(input)
			for i := 0; i <= length; i++ {
				getstring := removeRightChars(input)
				getSubstring = append(getSubstring, getstring...)
				input = removeLeftChars(input)
			}
			fmt.Println("The substring is ", getSubstring)
		}
	}
}
