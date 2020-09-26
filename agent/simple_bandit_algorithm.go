package agent

import (
	"log"
	"rl-algo/model"
	"rl-algo/utils"
	"time"

	"golang.org/x/exp/rand"
)

// SimpleBandit implements "a simple bandit algorithm" (page 32)
//
// It takes the function `bandit()`` assuming to take an action and return a corresponding reward,
// plus an array of the `k` possible actions to take and the ɛ probability of taking a random action rather than the greedy one.
func SimpleBandit(bandit model.SimpleActionFunc, actions []model.Action, epsilon float64) {
	if len(actions) == 0 {
		log.Fatalln("no action provided")
	}
	if epsilon < 0 || epsilon > 1 {
		log.Fatalln("epsilon is not within the right probability range")
	}
	k := len(actions)
	src := rand.NewSource(uint64(time.Now().UnixNano()))

	// Initialize
	Q := make(model.Values, k)
	N := make(map[model.Action]uint, k)
	for _, a := range actions {
		Q[a] = 0
		N[a] = 0
	}

	for {
		var A model.Action
		// Random action with probability ɛ
		if ok, err := utils.Pick(epsilon, src); err == nil && ok {
			A = model.Random(actions)
		} else {
			A = *bandit.Argmax(Q)
		}
		R := bandit(A)
		N[A] = N[A] + 1
		Q[A] = Q[A] + model.UintToValue(1/N[A])*(R.Value()-Q[A])
	}
}

func TestSimpleBandit() {
	log.Println("Testing SimpleBandit agent... (Press Ctrl^C to end)")
	const FIRST_ACTION model.Action = 1
	const SECOND_ACTION model.Action = 2
	bandit := func(action model.Action) model.Reward {
		log.Println("Action taken", action)
		time.Sleep(1 * time.Second)
		return model.Reward(float64(action) * rand.Float64())
	}
	SimpleBandit(bandit, []model.Action{FIRST_ACTION, SECOND_ACTION}, .05)
}
