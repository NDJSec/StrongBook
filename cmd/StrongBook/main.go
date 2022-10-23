package main

import (
	"fmt"
	"os"

	"github.com/NDJSec/StrongBook/cmd/StrongBook/aescrypto"
)

func main() {

	if os.Args[1] != string('d') {
		if os.Args[1] != string('e') {
			usageMenu()
		}
		if os.Args[1] == string('e') {
			if len(os.Args[1:]) != 4 {
				usageMenu()
			}
			aescrypto.EncryptFile(os.Args[2], os.Args[3], os.Args[4])
		}
	} else if os.Args[1] == string('d') {
		fmt.Println("Running Decryption")
	} else {
		usageMenu()
	}
}

func usageMenu() {
	fmt.Println("Encryption Usage: StrongBook.exe e [password] [inputfile] [outfile]")
	fmt.Println("StrongBook.exe e 0123456789 test.txt test.enc")
	fmt.Printf("Note: Password must be at least 10 characters and the original inputfile will get deleted.\n\n")

	fmt.Println("Decryption Usage: StrongBook.exe d [password] [inputfile] [outputfile]")
	fmt.Println("StrongBook.exe e 0123456789 test.enc test.txt")
	os.Exit(1)
}
