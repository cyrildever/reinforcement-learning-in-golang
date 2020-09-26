package utils

import (
	"errors"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// Pick returns `true` if one should take the action, ie. considering the wished probability ɛ,
// if the random number returned by the uniform distribution is equal or below ɛ, than it's considered a pick.
func Pick(epsilon float64, src rand.Source) (bool, error) {
	if epsilon < 0 || epsilon > 1 {
		return false, errors.New("invalid probability")
	}
	dist := distuv.Uniform{
		Min: 0,
		Max: 1,
		Src: src,
	}
	rnd := dist.Rand()
	return rnd <= epsilon, nil
}
