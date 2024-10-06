//This program is a workspace for practicing and getting used to the basics of Golang
package main

import "fmt"
import "log"
import "time"

//Variable outside of any scope?
var s = "seven"

//Structs!
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	fmt.Println("Hello Go Learner")

	var whatToSay string
	var i int

	whatToSay = "I don't know what to say"

	fmt.Println(whatToSay)

	i = 42

	fmt.Println("i is set to: ", i)

	//Functions in go can return more than one variable/thing!
	one, two := returnMultiple()
	fmt.Println("Two variables from one function: ", one, two)

	//Using := is a short hand for declaring a variable, the variable will be of the type on the right hand side
	messageInABottle := sendMessage()

	fmt.Println(messageInABottle)

	//Pointers in Go	

	myString := "Green"
	log.Println("myString is set to: ", myString)	
	changeUsingPointer(&myString)
	log.Println("after calling changeUsingPointer myString is: ", myString)

	//Print the global/package level variable s
	log.Println("Global variable s is: ", s)

	//Using the user struct
	user := User {
		FirstName: "Johnny",
		LastName: "Knoxville",
	}

	log.Println(user.FirstName, user.LastName)

	//Didnt assign birthdate to user, what does it print out? Let see
	log.Println(user.FirstName, user.LastName, user.BirthDate)

}

func sendMessage() string {
	var returnMessage string

	returnMessage = "This is a Message in a bottle"

	return returnMessage
}

func returnMultiple() (string, string) {
	return "thing one", "thing two"
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	log.Println("s is set to:", s)
	*s = newValue
}