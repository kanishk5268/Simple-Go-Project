package main

import (
	"fmt"

	"example.com/control-structure/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24*7 on :", randomdata.PhoneNumber())

	presentOptions()
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("------------")
		//panic("can't continue sorry")
	}

	for {

		//fmt.Print("hello world")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		fmt.Println("Your choice: ", choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Println("Deposit Value: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("New Account Balance:", accountBalance)
			fileops.WritefloatTofile(accountBalance, accountBalanceFile)

		case 3:
			fmt.Println("Withdraw Value: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("New Account Balance:", accountBalance)

			fileops.WritefloatTofile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}

		// wantsCheckBalance := choice == 1

		// if choice == 1 {
		// 	fmt.Println("Your balance is: ", accountBalance)
		// } else if choice == 2 {

		// } else if choice == 3 {
		// 	fmt.Println("Withdraw Value: ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}

		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Invalid amount. You can't withdraw more than you have")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("New Account Balance:", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
	}

}
