package main

import "log"
import "sort"

type myStruct struct {
	FirstName string
}

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

// m is a reciever
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Johnny"

	myVar2 := myStruct{
		FirstName: "Sally",
	}

	//using the function with a reciever
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())

	//Maps

	//var myOtherMap map[string]string // you can't do any thing with this map

	myMap := make(map[string]string) //First string is the type for the key second is the type for the value

	myMap["dog"] = "Teddy"

	log.Println(myMap["dog"])

	myMap["otherDog"] = "Georgie"

	log.Println(myMap["otherDog"])

	//Kinda like dictionarys or json

	myMap["dog"] = "Teddy is the best boy"

	log.Println(myMap["dog"])

	myMapTwo := make(map[string]int)

	myMapTwo["First"] = 1

	log.Println(myMapTwo["First"])

	//Use it with structs

	myMapStruct := make(map[string]myStruct)

	testStruct := myStruct {
		FirstName: "Connor",
	}

	myMapStruct["me"] = testStruct

	log.Println(myMapStruct["me"].FirstName)

	//Map are the same no matter where they are in the program
	//Meaning you don't need to use pointers with maps

	//Slices

	//Used instead of arrays

	var mySlice []string

	mySlice = append(mySlice, "Connor")
	mySlice = append(mySlice, "Abby")
	mySlice = append(mySlice, "Johnny")

	log.Println(mySlice)

	sort.Strings(mySlice)

	log.Println(mySlice)

	numbers := []int{88, 1, 2, 3, 4, 5, 6, 2, 4 ,5, 9, 10}

	log.Println(numbers)
	log.Println(numbers[0:2])


	//Decisions

	isTrue := true

	//If is the same as most languages
	if isTrue == true {
		log.Println("isTrue is ", isTrue)
	} else {
		log.Println("isTrue is ", isTrue)
	}

	myNum := 100

	if myNum > 99 && isTrue {
		log.Println("myNum is > 99 and isTrue is true")
	}

	//Switch statement break out once its reaches its first match, doesn't need to specify break
	switch myNum {
	case 99:
		log.Println(myNum)

	case 100:
		log.Println(myNum)

	case 101:
		log.Println(myNum)

	default:
		log.Println("What")
	}

	//Loop over data/ range of data

	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	//Iterate over a slice
	for i, people := range mySlice {
		log.Println(i, people)
	}
	//I case you don't want to use i, sense it will throw error if you don't
	for _, people := range mySlice {
		log.Println(people)
	}

	for key, value := range myMap {
		log.Println(key, value)
	}

	var aString string = "this is a string"

	for i, l := range aString {
		log.Println(i, l) //A string is a slice of bytes, so you get numbers output instead of chars
	}

	type User struct {
		FirstName string
		LastName string
		Age int
	}

	var users []User
	users = append(users, User{"Johnny", "Knoxville", 45})
	users = append(users, User{"April", "Knowles", 88})
	users = append(users, User{"Terry", "Berry", 33})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Age)
	}

	//Interfaces

	dog := Dog{
		Name: "Teddy",
		Breed: "Shiba",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name: "Johnny",
		Color: "Gray",
		NumberOfTeeth: 100,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "Legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Grrr"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}