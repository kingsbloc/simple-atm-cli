package internals

import (
	"fmt"
)

var balance float64 = 1000000

// Deposit Funds
func Deposit() {
	var input float64
	fmt.Print(Yellow + "Enter amount to deposit: " + Reset)
	fmt.Scan(&input)
	balance = balance + input
	fmt.Println(Blue + "Deposit Successful" + Reset)
	NewLine(2)
	Menu()
}

// Withraw Funds
func Withdraw() {
	var input float64
	fmt.Print(Yellow + "Enter amount to withdraw: " + Reset)
	fmt.Scan(&input)
	if input > balance {
		fmt.Println(Red + "Insufficent funds" + Reset)
		Menu()
		return
	}
	balance = balance - input
	fmt.Println(Blue + "Withdrawal Successful" + Reset)
	NewLine(2)
	Menu()
}

// Check Current Balance
func CheckBalance() {
	fmt.Printf(Blue+"Current Balance is %.2f \n"+Reset, balance)
	NewLine(2)
	Menu()
}
