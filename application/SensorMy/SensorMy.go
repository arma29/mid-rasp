package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Seed with a constantly changing number
	rand.Seed(time.Now().UnixNano())

	// Differentiating Radiatin Type and Origin
	radTypes := [3]string{"Alfa", "Beta", "Gamma"}
	radOrigins := [2]string{"RadCorps", "ActiveX"}
	
	radType := radTypes[rand.Intn(len(radTypes))]
	radOrigin := radOrigins[rand.Intn(len(radOrigins))]
	radValue := rand.Float64()

	fmt.Printf("%s - %s %f", radOrigin, radType, radValue)
}
