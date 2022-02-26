package main

import "fmt"

func main() {
	fmt.Println("Welcome to QUIZRY!")

	fmt.Println("Enter your name: ")
	var name string 
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Println("Enter your age: ")
	var age int 
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yeah, you can play!")
	}else {
		fmt.Println("You can not play!")
		return
	}

	score := 0

	fmt.Printf("Who developed Go language, Google or Microsoft?\n")

	var answer string
	fmt.Scan(&answer)

	if answer == "Google" {
		fmt.Println("Correct!")
		score++
	}else {
		fmt.Println("Incorect!")
	}


	fmt.Println("In which year was GO Lang introduced?")

	var year uint 
	fmt.Scan(&year)

	if year == 2012 {
		fmt.Println("Correct!")
		score++
	}else {
		fmt.Println("Incorect!")
	}

	fmt.Printf("\nScore: %v", score)
}