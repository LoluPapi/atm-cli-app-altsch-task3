package main

import (
	"fmt"
	"os"

	"github.com/lolupapi/atm-cli-app-altsch-task3/atm"
)

func welcome() {
	fmt.Println("**-**{ Welcome to the ATM CLI app!}**-**")
}

func newLine(numberOfLines int) {
	i := 0
	for i < numberOfLines {
		fmt.Println("\n")
		i++
	}
}

func anotherTransaction() {
	var yesOrNo string
	newLine(2)
	fmt.Printf("\n\nDo you want another transaction?\nPress y to proceed and n to exit\n\n")
	fmt.Scan(&yesOrNo)
	if yesOrNo == "y" {
		transaction()
	}
	endSession()
}

func transaction() {
	newLine(1)
	fmt.Printf("\nEnter any option to be served!\n\n")
	fmt.Printf("1. Change Pin\n")
	fmt.Printf("2. Balance\n")
	fmt.Printf("3. Withdraw\n")
	fmt.Printf("4. Deposit\n\n")
	newLine(1)

	var transactionNumber int
	fmt.Printf("Enter Transaction number: ")
	_, err := fmt.Scan(&transactionNumber)
	if err != nil {
		fmt.Println("Error:, ")
	}

	switch transactionNumber {
	case 1:
		atm.ChangeSecretPin()
		anotherTransaction()
	case 2:
		atm.Balance()
		anotherTransaction()
	case 3:
		atm.withdraw()
		anotherTransaction()
	case 4:
		atm.Deposit()
		anotherTransaction()
	case 0:
		endSession()
	default:
		fmt.Println("Invalid input ")
		anotherTransaction()
	}
}

func main() {
	atm.login()
	welcome()
	transaction()
}

func endSession() {
	fmt.Println("\nThanks for using our service!!! \nHave a nice day👋")
	os.Exit(0)
}
