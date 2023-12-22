package get

import (
	"math/rand"
	"time"
)

func RandomCode() []byte {
	pins := []byte{49, 50, 51, 52, 53, 54}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(pins), func(i, j int) { pins[i], pins[j] = pins[j], pins[i] })
	randomPins := []byte{}
	for i := 0; i < 4; i++ {
		randomPins = append(randomPins, pins[rand.Intn(len(pins))])
	}
	return randomPins
}
