package model

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// State defines the interface for any state implementation.
type State interface {
	IsTerminal() bool
	Vector() mat.Vector
}

// StateValue is a map of expected value for state.
type StateValue map[State]float64

// Get ...
func (sv StateValue) Get(at State) float64 {
	return sv[at]
}

// Clone ...
func (sv StateValue) Clone() StateValue {
	cloned := make(StateValue, len(sv))
	for s, v := range sv {
		cloned[s] = v
	}
	return cloned
}

// Print mimics Python's numpy matrix printing
func (sv StateValue) Print(states []State, width int) (str string) {
	i := 0
	str = "[[\t"
	for _, s := range states {
		v := sv[s]
		if v >= 0 {
			str += " "
		} else {
			str += "-"
		}
		if v == 0 {
			str += "0.\t\t"
		} else {
			str += fmt.Sprintf("%.6f\t", math.Abs(v))
		}
		i++
		if i == len(sv) {
			str += "]"
		} else if i%width == 0 {
			str += "]\n [\t"
		}
	}
	str += "]"
	return
}
