package main 

import (
	"fmt"
	"regexp"
)


func main() {
	emails := []string{
		"julia@hackerrank.com",
		"julia_@hackerrank.com",
		"julia_0@hackerrank.com",
		"julia0_@hackerrank.com",
		"julia@gmail.com",
		"mhas@hotmail.com",
	}

	re := regexp.MustCompile(`^[a-z]{1,6}_?[0-9]{0,4}\@hackerrank\.com$`)

	for _, v := range emails {
		match := re.FindString(v)
		if match == "" {
			fmt.Println("False", v)
		} else {
			fmt.Println("True", v)
		}
	}
}





