package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100
	tries := 0

	fmt.Println("Guess a number between ", low, "and ", high)
	fmt.Println("Press ENTER when you're ready")
	scanner.Scan()

	for {
		tries++
		guess := (high + low) / 2
		fmt.Println("I guess the number is ", guess)
		fmt.Println("Is that :")
		fmt.Println("(a) Too high")
		fmt.Println("(b) Too low")
		fmt.Println("(c) Correct")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("Yay!! I got it right!")
			fmt.Println("It took me ", tries, "tries to get it right")
			break
		} else {
			fmt.Println("Please enter a valid response")
		}

		if tries > int(math.Log2(float64(high))) {
			fmt.Println("Your guess is out of bounds or you missed the answer")
			break
		}

		// switch response {
		// case "a":
		// 	high = guess - 1
		// case "b":
		// 	low = guess + 1
		// case "c":
		// 	fmt.Println("Yay!! I got it right!")
		// 	break  // this does not work https://stackoverflow.com/questions/11104085/in-go-does-a-break-statement-break-from-a-switch-select
		// default:
		// 	fmt.Println("Please enter a valid response")
		// }
	}
}
