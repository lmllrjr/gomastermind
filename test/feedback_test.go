package get

import (
	"fmt"
	"testing"

	"github.com/lmllr/gomastermind/packages/get"
)

func TestFeedback(t *testing.T) {
	codemaker := []byte{49, 50, 51, 52}
	codebreaker := []byte{110, 200, 94, 94}

	want := []string{".", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{49, 50, 51, 52}
	codebreaker = []byte{49, 200, 51, 200}

	want = []string{"*", "*", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{49, 50, 51, 52}
	codebreaker = []byte{52, 51, 50, 49}

	want = []string{"+", "+", "+", "+"}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{49, 50, 51, 52}
	codebreaker = []byte{23, 50, 50, 50}

	want = []string{"*", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{49, 50, 51, 52}
	codebreaker = []byte{23, 49, 49, 49}

	want = []string{"+", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{49, 50, 51, 52}
	codebreaker = []byte{23, 49, 51, 49}

	want = []string{"*", "+", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{51, 50, 54, 49}
	codebreaker = []byte{49, 97, 97, 97}

	want = []string{"+", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{51, 54, 52, 52}
	codebreaker = []byte{52, 51, 51, 51}

	want = []string{"+", "+", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{54, 51, 51, 51}
	codebreaker = []byte{51, 49, 49, 49}

	want = []string{"+", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	// ...
	codemaker = []byte{51, 52, 53, 49}
	codebreaker = []byte{49, 49, 49, 50}

	want = []string{"+", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{51, 54, 54, 50}
	codebreaker = []byte{54, 54, 54, 54}

	want = []string{"*", "*", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{52, 54, 51, 51}
	codebreaker = []byte{51, 51, 49, 50}

	want = []string{"+", "+", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{49, 49, 53, 53}
	codebreaker = []byte{53, 53, 49, 49}

	want = []string{"+", "+", "+", "+"}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []byte{50, 53, 50, 51}
	codebreaker = []byte{49, 50, 49, 50}

	want = []string{"+", "+", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !contains(got, want...) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}
}

// helper func to compare 2 slices
// check if a and b contain the same elements
// the order does not matter
func contains(slice []string, values ...string) bool {
	for _, v := range values {
		found := false
		for _, s := range slice {
			if s == v {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
