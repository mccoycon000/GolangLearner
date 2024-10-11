package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/mccoycon000/mygoprogram/helpers"
)


type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType //made via package
	myVar.TypeName = "some name"
	log.Println(myVar.TypeName)

	//Channels

	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)


	//JSON


	myJson := `
[
    {
        "first_name": "Clark",
        "last_name": "Kent",
        "hair_color": "black",
        "has_dog": true
    },
    {
        "first_name": "Bruce",
        "last_name": "Wayne",
        "hair_color": "black",
        "has_dog": false
    }
]`


	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("error was encountered", err)
	}

	log.Printf("unmarshalled %v", unmarshalled)

	//write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Johnny"
	m1.LastName = "Knoxville"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Mary"
	m2.LastName = "Knoxville"
	m2.HairColor = "blue"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))


	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("results of division is", result)
}

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(10)
	intChan <- randomNumber
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y==0 {
		return result, errors.New("cannot divide by 0")
	}

	result = x / y

	return result, nil
}