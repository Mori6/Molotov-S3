package main

import (
	"fmt"
	"strings"
)

//given user|pass returns  user
func getUser(user_pass string) string { 
	user := strings.Split(user_pass, "|")
	return user[0]
}

//given user|pass returns  Pass
func getPass(user_pass string) string { 

	index := strings.Index(user_pass, "|") + 1
	if index == -1 { 
		fmt.Println("| not found")
		return ""
	}

	return user_pass[index:]
}


