package main

import (
	"fmt"
	"regexp"
)

func main() {

	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+@gmail.com$`)
	emailID := []string{
		"vikasbahndekar@gmail.com",
		"PraticRamelwar@gmal.com",
		"SanketBezalwar#kishor@gmail.com",
		"nitinFulzale@gmail.com",
		"ritik!1Bhandekar2034@gmail.com"}

	for _, email := range emailID {
		if pattern.MatchString(email) {
			fmt.Println(email, " = EmailID is correct")
		} else {
			fmt.Println(email, " = EmailID is Incorrect")
		}
	}

}
