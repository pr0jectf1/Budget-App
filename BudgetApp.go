package main

import "fmt"

type CreditDebt struct {
	CreditorName string
	Amount	     float64
}

type User struct {
	
	Id int
	FirstName string
	LastName string
	CombinedCreditDebt []CreditDebt
	
}

func createNewUser(Id int, firstName, lastName string) User {
	return User{Id: Id, FirstName: firstName, LastName: lastName}
}

func (u User) addCreditDebt(creditDebt CreditDebt) {
	
	u.CombinedCreditDebt = append(u.CombinedCreditDebt, creditDebt)
}

func (u User) addUserDebt() {
	var input string
	var numInput int
	var debtAmount float64

	for{
		fmt.Println("Hello", u.FirstName, "would you like to add any new debt? (y/n)", )
		fmt.Scan(&input)

		if input == "y" {
			fmt.Println("How many different debts would you like to add? (Please enter a number)")
			fmt.Scan(&numInput)

			for i:= 0; i < numInput; i++ {

				fmt.Print("Creditors name: ")
				fmt.Scan(&input)
				fmt.Print("Credit amount: ")
				fmt.Scan(&debtAmount)
				fmt.Print()

				//check if debt exists in users list
				// if the debt exists ask user if this is a new debt to the same creditor.
				//if so, create new debt with different debtId
				//if not, go on to next debt 
		
				creditDebt := CreditDebt{
					CreditorName: input,
					Amount: debtAmount,
				}
				u.addCreditDebt(creditDebt)
			}
			break
		} else if input == "n" {
			fmt.Printf("Wow %v, congrats on having no debt!\n\n", u.FirstName)
			break
		} else {
			fmt.Println("Sorry, that is not a valid input. Please try again.")
		}
	}
}

func getNewUserDetails() User{
	var firstName string
	var lastName string

	fmt.Println("Hello what's your first name?")
	fmt.Scan(&firstName)
	fmt.Println("What's your last name?")
	fmt.Scan(&lastName)
	newUser := createNewUser(1, firstName, lastName)

	return newUser
}

func main() {

	u1 := getNewUserDetails()
	u1.addUserDebt()

	fmt.Printf("What would you like to do \n")
	
}