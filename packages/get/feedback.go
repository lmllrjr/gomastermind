package get

import (
	"fmt"
	"math/rand"
	"time"
)

// check users guess on the code
// provides feedback by placing from zero to four chars to
// determine the correctness of position and color

// a "*" is placed for each char from the guess which
// is correct in both char and position.
// a "+" indicates the existence of a correct char placed in the wrong position.
func Fdbk(c []byte, g []byte) []string {
	b := map[int][2]bool{}

	for i, v := range c {
		if v == g[i] {
			b[i] = [2]bool{true, true}
		} else {
			for j, v2 := range g {
				if v2 == v {
					b[i] = [2]bool{false, true}
					g[j] = 'X'
					break
				} else {
					b[i] = [2]bool{false, false}
				}
			}
		}
	}

	// fmt.Println(b)

	bytes := []byte{}
	for _, v := range b {
		if v[0] && v[1] {
			bytes = append(bytes, '*')
		} else if v[1] {
			bytes = append(bytes, '+')
		} else {
			bytes = append(bytes, '.')
		}
	}

	return randomizeString(bytes)
}

// check two slices for equality
func Compare(a, b []byte) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func randomizeString(bytes []byte) []string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Shuffle(len(bytes), func(i, j int) {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	})

	s := make([]string, len(bytes))
	for i, byte := range bytes {
		s[i] = fmt.Sprintf("%c", byte)
	}

	return s
}
