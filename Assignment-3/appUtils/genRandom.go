package apputils

import (
	"log"
	"math/rand"
)

func randNumber(min int, max int) (int, int) {
	water := rand.Intn(max-min+1) + min
	wind := rand.Intn(max-min+1) + min

	log.Printf("Generate : [ Water : %d | Wind : %d ]\n", water, wind)
	return water, wind
}

func GenerateRandom() *BodyCuaca {
	min := 1
	max := 100

	water, wind := randNumber(min, max)

	status := &StatusCuaca{
		Water: water,
		Wind:  wind,
	}

	body := BodyCuaca{
		Status: *status,
	}
	return &body
}
