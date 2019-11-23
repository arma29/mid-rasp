package radiation

import (
	"math/rand"
	"time"
)

// Radiation is
type Radiation struct {
	Value     float32
	Timestamp int64 //nanosegundos
}

// Validator is
type Validator struct {
	IsDangerous bool
}

// GenerateRadiationValue returns random float32 value
func GenerateRadiationValue() float32 {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return myRand.Float32()
}

// IsRadiationDangerous is
func IsRadiationDangerous(r float32) bool {
	if r > 0.5 {
		return true
	}
	return false
}
