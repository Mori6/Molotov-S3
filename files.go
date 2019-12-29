package main

import (
	"fmt"

	"strings"
	"io/ioutil"
)

func countNewLines(filename string) int { 
	data, _ := ioutil.ReadFile(filename)
	var characterCount int = 0
	for _, character := range data { 
		if character == '\n' { 
			characterCount++
		}
	}
	return characterCount
}


func getUsernames(decryptedString string) []string { 
	fmt.Println("beginning scan...")
	var user_pass []string = strings.Split(decryptedString, "\n")
	numberOfSpaces := countNewLines("creds.txt")
	
	var users  = make([]string, numberOfSpaces)

	for _, data := range user_pass { 
		users = append(users, data)
	}
	return users
}

func getPasswords(decryptedString string) []string{ 
	var user_pass []string = strings.Split(decryptedString, "|")
	return user_pass
}