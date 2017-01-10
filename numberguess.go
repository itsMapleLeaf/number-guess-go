package numberguess

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
	rand.Seed(time.Now().UnixNano())

	number := rand.Uint32()%100 + 1
	reader := bufio.NewReader(os.Stdin)
	won := false
	guesses := 0

	fmt.Println("Guess the number! (1-100)")

	for !won {
		fmt.Print("Your guess: ")
		input, readError := reader.ReadString('\n')
		guess, numConvertError := strconv.ParseUint(strings.Trim(input, "\r\n"), 10, 32)

		if readError != nil || numConvertError != nil {
			fmt.Println("Try guessing an actual number...")
			continue
		}

		switch {
		case uint32(guess) < number:
			fmt.Println("Too low")
		case uint32(guess) > number:
			fmt.Println("Too high")
		default:
			fmt.Println("You win!")
			won = true
		}

		guesses++
	}

	fmt.Printf("Took %d guesses", guesses)
}
