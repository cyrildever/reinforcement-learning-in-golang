package model

import (
	"time"

	"golang.org/x/exp/rand"
)

// Action identifies an action with an integer
type Action int

// SimpleActionFunc is the "interface" for a simple action giving a reward, ie.
// when implementing such a function, it should take the action and return the corresponding reward amount.
type SimpleActionFunc func(action Action) Reward

// Argmax returns the greedy action in the passed map of values
// or a random action if multiple actions are greedy
func (f SimpleActionFunc) Argmax(qs Values) (greedy *Action) {
	var highest Value = 0
	var greedies []*Action
	for a, r := range qs {
		if r > highest {
			greedies = append(greedies, &a)
			highest = r
		}
	}
	if len(greedies) != 1 {
		rand.Seed(uint64(time.Now().UnixNano()))
		var random int
		if len(greedies) == 0 {
			random = rand.Intn(len(qs))
			actions := []Action{}
			for key := range qs {
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
