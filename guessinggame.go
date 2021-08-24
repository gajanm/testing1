package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	attempts := 0

	fmt.Println(" Welcome to Guess to Win ")
	fmt.Println("-------------------------")
	fmt.Println("You will try to guess a number between 1 and 10. I will select this special number. You shall start by guessing a number number and I’ll say if your said number was the number that I selected, higher than the number that I selected, or lower than the number that I selected. You shall get 5 tries to guess this number")
	fmt.Println("Type in your guess between 1 and 10")
	min, max := 1, 10
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(max-min) + min
	reader := bufio.NewReader(os.Stdin)

	for {
		attempts++
		if attempts > 5 {
			fmt.Println("You've used over 5 turns to guess the secret number. The secret number is", secretNumber)
		}
		input, err :=w reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			return
		}

		fmt.Println("Your guess is", guess)

		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Try again")
		} else {
			fmt.Println("Correct, you Legend! You guessed right after", attempts, "attempts. To play again, type in 'go run guessinggame.go'")
			break
		}
	}

}
