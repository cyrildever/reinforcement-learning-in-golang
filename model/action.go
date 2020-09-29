package model

import (
	"math"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
)

// ActionFunc is the actual m.o. for any action granting a reward by transforming a state into another,
// ie. when implementing such a function, it should take the current state and the action to take and
// return the new state and corresponding reward.
type ActionFunc func(State, Action) (State, Reward)

// Action interface must be implemented as struct of n parameter(s)/dimension(s).
type Action interface {
	GetName() string
	ValueFunc() ActionFunc
	Vector() mat.Vector
}

// ActionValues is a map of action -> reward values.
type ActionValues map[Action]Reward

// Argmax returns the greedy action in the passed map of values
// or a random action if multiple actions are greedy.
func Argmax(values map[Action]Reward) (greedy *Action) {
	var highest Reward = 0
	var greedies []*Action
	for a := range values {
		_, v := a.ValueFunc()(nil, a)
		if v > highest {
			greedies = append(greedies, &a)
			highest = v
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
	Name      string
	Parameter float64
}

func (a Action1D) GetName() string {
	return a.Name
}
func (a Action1D) ValueFunc() ActionFunc {
	return func(current State, _ Action) (next State, r Reward) {
		next = current
		r = Reward(a.Parameter)
		return
	}
}
func (a Action1D) Vector() mat.Vector {
	return mat.NewVecDense(1, []float64{a.Parameter})
}

// Action2D is a two-dimension vector implementing the Action interface
type Action2D struct {
	Name string
	X    float64
	Y    float64
}

func (a Action2D) GetName() string {
	return a.Name
}
func (a Action2D) ValueFunc() ActionFunc {
	return func(current State, _ Action) (next State, r Reward) {
		next = current
		r = Reward(math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2)))
		return
	}
}
func (a Action2D) Vector() mat.Vector {
	return mat.NewVecDense(2, []float64{a.X, a.Y})
}

// Action3D is a three-dimension vector implementing the Action interface
type Action3D struct {
	Name string
	X    float64
	Y    float64
	Z    float64
}

func (a Action3D) GetName() string {
	return a.Name
}
func (a Action3D) ValueFunc() ActionFunc {
	return func(current State, _ Action) (next State, r Reward) {
		next = current
		r = Reward(math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2) + math.Pow(a.Z, 2)))
		return
	}
}
func (a Action3D) Vector() mat.Vector {
	return mat.NewVecDense(3, []float64{a.X, a.Y, a.Z})
}

// Action4D is a four-dimension vector implementing the Action interface
type Action4D struct {
	Name string
	X    float64
	Y    float64
	Z    float64
	T    float64
}

func (a Action4D) GetName() string {
	return a.Name
}
func (a Action4D) ValueFunc() ActionFunc {
	return func(current State, _ Action) (next State, r Reward) {
		next = current
		r = Reward(math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2) + math.Pow(a.Z, 2) + math.Pow(a.T, 2)))
		return
	}
}
func (a Action4D) Vector() mat.Vector {
	return mat.NewVecDense(4, []float64{a.X, a.Y, a.Z, a.T})
}
