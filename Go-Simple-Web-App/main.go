package main

//This is a simple web app in golang, stuff that is commented out is something I used to test then wanted to remove but still want to show how I did it for reference

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// Home is the home page value
func Home(w http.ResponseWriter, r *http.Request) {
	/*
		fmt.Fprintf(w, "This is the home page")
	*/
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	/*
		sum := addValues(2, 2)
		_, _ = fmt.Fprintf(w, fmt.Sprintf("this is the about page and 2 + 2 is %d", sum))
	*/
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parseing template:", err)
		return
	}
}

/*
// addValues adds two integers and returns the sum
func addValues(x,y int) int {
	return x + y
}

//Divide landing page for divide, used to test err
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 0.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
*/

// main is the main function starts the server on localhost
func main() {

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			n, err := fmt.Fprintf(w, "Hello, world!")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
		})
	*/

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	//http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s \n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
