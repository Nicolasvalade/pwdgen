package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

// how long the generated password should be
const PWD_LEN int = 16

func main() {
	err := clipboardInit()
	if err != nil {
		return
	}

	rand.Seed(time.Now().UnixNano())
	chars := getChars()
	pwd := makePassword(chars)
	clipboard.Write(clipboard.FmtText, pwd)

	fmt.Println("Password was copied to clipboard!")
}

func clipboardInit() error {
	fmt.Println("Initializing clipboard...")
	err := clipboard.Init()
	// clipboard.Init returns error only if it doesn't find the OS clipboard API
	for err != nil {
		fmt.Println("Initialization failed.")
		fmt.Println("Should we try again ? Y/N")

		var input string
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "N" {
			return err
		}

		fmt.Println("Trying again...")
		err = clipboard.Init()
	}
	fmt.Println("Success!")
	return nil
}

func getChars() string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbols := "£$!&+-="
	digits := "0123456789"
	return lower + upper + symbols + digits
}

func makePassword(chars string) []byte {
	nChars := len(chars) - 1

	pwd := make([]byte, PWD_LEN)
	for i := 0; i < PWD_LEN; i++ {
		pwd[i] = chars[rand.Intn(nChars)]
	}

	return pwd
}
