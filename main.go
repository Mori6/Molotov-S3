/*
Amazonrake by Mori
this program uses amazon accounts to make Amazon S3 servers
Don't run if you don't know how to protect yourself from defenses
*/
 
package main


import (
	"fmt"
	"io/ioutil"
)


func main() { 
	fmt.Println("looking for encrypted file...\n")
	var file string = "creds.txt"
	if (isEncrypted(file)) { 
		fmt.Println("file is encrypted...")
		//fmt.Println(string(decryptFile(file, "password")))
		var decrypted string = (string(decryptFile(file, "password")))
		var sliceUsernames []string = getUsernames(decrypted)
		for _, data := range sliceUsernames { 
			fmt.Println("trying user: ", getUser(data), "with pass: ", getPass(data))
		}
		//fmt.Println("%q\n", sliceUsernames)
		fmt.Println(sliceUsernames[126])
	} else { 
		fmt.Println("encrypting file...")
		data, _ := ioutil.ReadFile(file)
		encryptFile(file, []byte(data), "password")
		var decrypted string = (string(decryptFile(file, "password")))
		//fmt.Println(decrypted)
		var sliceUsernames = getUsernames(decrypted)
		for _, data := range sliceUsernames { 
			fmt.Println(getUser(data))
		}
		//fmt.Println("%q\n", sliceUsernames)
	}
	
	
}