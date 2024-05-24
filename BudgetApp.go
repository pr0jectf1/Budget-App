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

func main() {
	u1 := createNewUser(1, "Johnny", "Garcia")
	fmt.Println(u1.FirstName)
}