package main

import "fmt"

func guessGame() {
	magicNum := 34
	chance := 3
	var guess int
	var try int
	for guess != magicNum && try != chance {
		fmt.Println("Write your guess:")
		fmt.Scan(&guess)
		try++
	}
	if magicNum != guess {
		fmt.Printf("My guess was %v\n", magicNum)
	} else {
		fmt.Println("Congralutations!")
	}

}

func main() {
	guessGame()
}
