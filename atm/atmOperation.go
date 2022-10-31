package atm

import (
	"fmt"
)

var balance float32 = 0
var amount float32
var secretPin string = "1234"

func newLine(numberOfLines int) {
	i := 0
	for i < numberOfLines {
		fmt.Println("\n")
		i++
	}
}

func login() {
	for {

		if !VerifySecretPin(secretPin) {
			continue
		}

		break
	}
}

func ChangeSecretPin() {
	newLine(1)
	fmt.Println("Change Pin")
	for {
		fmt.Printf("Enter your new pin: ")
		var new_pin string
		fmt.Scan(&new_pin)

		if len(new_pin) != 4 {
			fmt.Println("Your Pin should be 4 characters long")
			continue
		}
		if new_pin == secretPin {
			fmt.Println("You have entered your old pin, please entered a different pin")
			continue
		}

		secretPin = new_pin
		fmt.Println("Your Pin has been changed Succesfully")
		break
	}
}

func VerifySecretPin(secretPin string) bool {
	var secret_pin string
	fmt.Printf("ENTER YOUR PIN(Ensure no one is looking): ")
	fmt.Scan(&secret_pin)

	if secretPin != secret_pin {
		fmt.Println("Entered Pin is Incorrect, Please try again")
		return false
	}

	return true
}

func Balance() {
	fmt.Printf("Your account Balance is %.2f USD", balance)

}

func withdraw() {

	fmt.Println("Withdraw money")

	if balance == 0 {
		fmt.Println("Your account is empty, please add money")
		return
	}
	fmt.Printf("Enter amount: ")
	fmt.Scan(&amount)
	if amount < 5 {
		fmt.Println("invalid amount, enter amount greater than 5 to be able to make a withdrawal")
		return
	}

	if amount > balance {
		fmt.Println("There is insufficient funds in your account")
		fmt.Printf("Current balance: %.2f USD", balance)
		return
	}
	old_balance := balance
	balance -= amount

	fmt.Printf("\nOld balance: %.2f \n", old_balance)
	fmt.Printf("Amount withdrawn: %.2f  USD,\nCurrent balance: %.2f  USD", amount, balance)
}

func Deposit() {
	fmt.Println("Please make your deposits")

	fmt.Printf("Enter amount: ")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("invalid amount")
		fmt.Println("You cannot deposit less than $0")
		return
	}
	prev_balance := balance
	balance += amount

	fmt.Printf("Previous Balance: %.2f  USD,\nCurrent Balance: %.2f USD", prev_balance, balance)
}
