package model

import (
	"math"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
)

// Action interface must be implemented as some kind of vector of n parameter(s)/dimension(s)
type Action interface {
	Value() Value
	Vector() mat.Vector
}

// ActionFunc is the actual m.o. for any action granting a reward by transforming a state into another,
// ie. when implementing such a function, it should take the current state and the action to take and
// return the new state and corresponding reward.
type ActionFunc func(*State, Action) (*State, Reward)

// Argmax returns the greedy action in the passed map of values
// or a random action if multiple actions are greedy
func Argmax(values map[Action]Value) (greedy *Action) {
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

func (a Action1D) Vector() mat.Vector {
	return mat.NewVecDense(1, []float64{a.Parameter})
}

// Action2D is a two-dimension vector implementing the Action interface
type Action2D struct {
	X float64
	Y float64
}

func (a Action2D) Value() Value {
	return Value(math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2)))
}

func (a Action2D) Vector() mat.Vector {
	return mat.NewVecDense(2, []float64{a.X, a.Y})
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

func (a Action3D) Vector() mat.Vector {
	return mat.NewVecDense(3, []float64{a.X, a.Y, a.Z})
}

// Action4D is a four-dimension vector implementing the Action interface
type Action4D struct {
	X float64
	Y float64
	Z float64
	T float64
}

func (a Action4D) Value() Value {
	return Value(math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2) + math.Pow(a.Z, 2) + math.Pow(a.T, 2)))
}

func (a Action4D) Vector() mat.Vector {
	return mat.NewVecDense(4, []float64{a.X, a.Y, a.Z, a.T})
}
