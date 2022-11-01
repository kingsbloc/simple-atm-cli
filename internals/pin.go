package internals

import "fmt"

var pin string = "0000"

func Login() {
	for {
		if ValidatePin() {
			break
		}
		continue
	}
}

func ValidatePin() bool {
	var input string
	fmt.Printf(Yellow + "Enter Pin: " + Reset)
	fmt.Scan(&input)
	if pin != input {
		fmt.Println(Red + "Invalid pin" + Reset)
		return false
	}
	return true
}

// Change Pin
func ChangePin() {
	var input string
	var cinput string
	for {
		fmt.Printf(Yellow + "Please enter current pin:" + Reset)
		fmt.Scan(&input)
		if pin != input {
			fmt.Println(Red + "Invalid pin" + Reset)
			continue
		} else {
			break
		}
	}
	fmt.Printf(Yellow + "Please enter new pin: " + Reset)
	fmt.Scan(&input)
	fmt.Printf(Yellow + "Please verify new pin: " + Reset)
	fmt.Scan(&cinput)
	if input != cinput {
		fmt.Println(Red + "Invalid pin" + Reset)
		Menu()
	} else {
		pin = input
		fmt.Println(Blue + "Please changed Successfully" + Reset)
		InitATM()
	}
}
