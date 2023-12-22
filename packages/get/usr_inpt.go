package get

import (
	"bufio"
	"os"
)

// scan terminal for user input
// only valid inputs are four single characters
// examples:
//	1234
//	5314
//	1111
func UsrInpt() []byte {
	scanner := bufio.NewScanner(os.Stdin)

	var usrInpt []byte
	for {
		scanner.Scan()
		usrInpt = []byte(scanner.Text())

		if len(usrInpt) == 4 {
			break
		}
	}
	return usrInpt
}
