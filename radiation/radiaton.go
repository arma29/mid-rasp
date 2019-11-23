package radiation

import (
	"math/rand"
	"time"
)

// GetRadiation returns random float32 value
func GetRadiation() float32 {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return myRand.Float32()
}
