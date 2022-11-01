package internals

import (
	"fmt"
	"os"
)

func NewLine(n int) {
	for n < 0 {
		fmt.Println("")
		n--
	}
}

func Exit() {
	fmt.Println(Blue + "Thank you. Do have a great day!" + Reset)
	os.Exit(0)
}

func Welcome() {
	fmt.Println(Yellow + "** Welcome to the simple ATM **" + Reset)
	NewLine(1)
	fmt.Println(Yellow + "Please select an action:" + Reset)
	fmt.Println("1. Enter Pin")
	fmt.Println("0. Exit")
	var actionCode int
	fmt.Printf(Yellow + "Enter desired action code: " + Reset)
	_, err := fmt.Scan(&actionCode)
	if err != nil {
		fmt.Println(Red + "Invalid action code. Please enter a valid action code ==" + Reset)
	}
	switch actionCode {
	case 0:
		Exit()
	case 1:
		Login()
	default:
		Welcome()
	}
}

func InitATM() {
	Welcome()
	Menu()
}

func Menu() {
	NewLine(1)
	fmt.Println(Yellow + "Please select an action:" + Reset)
	fmt.Println("1. Change Pin")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Check Balance")
	fmt.Println("0. Exit")
	NewLine(1)

	var actionCode int
	fmt.Printf(Yellow + "Enter desired action code: " + Reset)
	_, err := fmt.Scan(&actionCode)
	if err != nil {
		fmt.Println(Red + "Invalid action code. Please enter a valid action code ==" + Reset)
	}
	switch actionCode {
	case 0:
		Exit()
	case 1:
		ChangePin()
	case 2:
		Deposit()
	case 3:
		Withdraw()
	case 4:
		CheckBalance()
	default:
		fmt.Println(Red + "Invalid action code" + Reset)
		Menu()
	}
}
