package imgsim

import (
	"fmt"
	"math/bits"
)

// Hash contains the perceptual hash stored as bits in a 64 bit unsigned integer.
type Hash uint64

func (h Hash) String() string {
	return fmt.Sprintf("%064b", h)
}

// Distance calculates the number of different bits in the hash
func Distance(a, b Hash) int {
	return bits.OnesCount64(uint64(a ^ b))
}
