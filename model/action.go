package model

import (
	"math"
	"time"

	"golang.org/x/exp/rand"
)

// Action interface must be implemented as a vector of n parameter(s)/dimension(s)
type Action interface {
	Value() Value
}

// ActionFunc is the "interface" for any action giving a reward, ie.
// when implementing such a function, it should take the action and return the corresponding reward amount.
type ActionFunc func(Action) Reward

// Argmax returns the greedy action in the passed map of values
// or a random action if multiple actions are greedy
func (f ActionFunc) Argmax(values map[Action]Value) (greedy *Action) {
	var highest Value = 0
	var greedies []*Action
	for a := range values {
		if a.Value() > highest {
			greedies = append(greedies, &a)
			highest = a.Value()
		}
	}
	if len(greedies) != 1 {
		rand.Seed(uint64(time.Now().UnixNano()))
		var random int
		if len(greedies) == 0 {
			random = rand.Intn(len(values))
			actions := []Action{}
			for key := range values {
				actions = append(actions, key)
			}
			greedy = &actions[random]
			return
		}
		random = rand.Intn(len(greedies))
		actions := []*Action{}
		actions = append(actions, greedies...)
		greedy = actions[random]
		return
	}
	greedy = greedies[0]
	return greedy
}

// RandomFrom returns a random item from the passed slice
func Random(items []Action) Action {
	rand.Seed(uint64(time.Now().UnixNano()))
	random := rand.Intn(len(items))
	return items[random]
}

//--- "Vectorized" standard implementations

// Action1D is a one-dimension vector implementing the Action interface
type Action1D struct {
	Parameter float64
}

func (a Action1D) Value() Value {
	return Value(a.Parameter)
}

// Action2D is a two-dimension vector implementing the Action interface
type Action2D struct {
	X float64
	Y float64
}

func (a Action2D) Value() Value {
	return Value(math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2)))
}

// Action3D is a three-dimension vector implementing the Action interface
type Action3D struct {
	X float64
	Y float64
	Z float64
}

func (a Action3D) Value() Value {
	return Value(math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2) + math.Pow(a.Z, 2)))
}

// Action4D is a three-dimension vector implementing the Action interface
type Action4D struct {
	X float64
	Y float64
	Z float64
	T float64
}

func (a Action4D) Value() Value {
	return Value(math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2) + math.Pow(a.Z, 2) + math.Pow(a.T, 2)))
}
