package agent

import (
	"log"
	"math"
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

	monitor := utils.LiveMonitor{Value: "reward"}
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

		monitor.ComputeAndLog(float64(R))
	}
}

func TestSimpleBandit() {
	log.Println("Testing SimpleBandit (10-armed testbed agent)... (Press Ctrl^C to end)")
	const (
		FIRST model.Action = 1 + iota
		SECOND
		THIRD
		FOURTH
		FIFTH
		SIXTH
		SEVENTH
		EIGHTH
		NINTH
		TENTH
	)
	bandit := func(action model.Action) model.Reward {
		log.Println("Action taken:", action)
		time.Sleep(1 * time.Second)
		return model.Reward(math.Pow(float64(action), 1/float64(action)))
	}
	SimpleBandit(bandit, []model.Action{FIRST, SECOND, THIRD, FOURTH, FIFTH, SIXTH, SEVENTH, EIGHTH, NINTH, TENTH}, .05)
}
