package main

import (
	"fmt"
	"os"

	"github.com/lmllr/gomastermind/packages/get"
)

func main() {
	code := get.RandomCode()[:4]

	// fmt.Println(code)
	fmt.Println("Crack the secret code 🔒")

	for i := 0; i < 12; i++ {
		ui := get.UsrInpt()
		// fmt.Println(ui)

		fdbk := get.Fdbk(code, ui)
		fmt.Println(fdbk)

		if get.Compare(code, ui) {
			fmt.Println("Code cracked 🔓")
			os.Exit(0)
		}
	}
}

// OBJECT OF THE GAME
// The object of MASTERMIND (r) is to guess a secret code consisting of a series of 4 numbers/colors.
// Each guess results in feedback narrowing down the possibilities of the code.
// You win when you can crack the code within 12 guesses.

// (1) A "*" char indicates a Code Character of the right number and
// in the right position (without indication of which position it corresponds to).

// (2) A "+" indivates a Character of the right number but in the wrong position.

// (3) No Character to indicate a wrong number that does not appear in the secret code.
