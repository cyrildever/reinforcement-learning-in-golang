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
// It takes the function `bandit()` assuming to take an action and return a corresponding reward (w/o using any state),
// plus an array of the `k` possible actions to take and the ɛ probability of taking a random action rather than the greedy one.
func SimpleBandit(bandit model.ActionFunc, actions []model.Action, epsilon float64) {
	if len(actions) == 0 {
		log.Fatalln("no action provided")
	}
	if epsilon < 0 || epsilon > 1 {
		log.Fatalln("epsilon is not within the right probability range")
	}
	k := len(actions)
	src := rand.NewSource(uint64(time.Now().UnixNano()))

	// Initialize
	Q := make(model.ActionValues, k)
	N := make(map[model.Action]uint, k)
	for _, a := range actions {
		Q[a] = 0
		N[a] = 0
	}

	monitor := utils.LiveMonitor{Output: "reward"}
	for {
		var A model.Action
		// Random action with probability ɛ
		if ok, err := utils.Pick(epsilon, src); err == nil && ok {
			A = model.Random(actions)
		} else {
			A = *model.Argmax(Q)
		}
		_, R := bandit(nil, A)
		N[A] = N[A] + 1
		Q[A] = Q[A] + model.Reward(1/N[A])*(R-Q[A])

		monitor.ComputeAndLog(R.Value())
	}
}

//--- TEST

func TestSimpleBandit() {
	log.Println("Testing SimpleBandit (10-armed testbed agent)... (Press Ctrl^C to end)")
	const (
		FIRST float64 = 1 + iota
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
	actions := []model.Action{
		model.Action1D{Parameter: FIRST},
		model.Action1D{Parameter: SECOND},
		model.Action1D{Parameter: THIRD},
		model.Action1D{Parameter: FOURTH},
		model.Action1D{Parameter: FIFTH},
		model.Action1D{Parameter: SIXTH},
		model.Action1D{Parameter: SEVENTH},
		model.Action1D{Parameter: EIGHTH},
		model.Action1D{Parameter: NINTH},
		model.Action1D{Parameter: TENTH},
	}
	bandit := func(_ model.State, action model.Action) (model.State, model.Reward) {
		log.Println("Action taken:", action)
		time.Sleep(1 * time.Second)
		_, r := action.ValueFunc()(nil, action)
		return nil, model.Reward(math.Pow(r.Value(), 1/r.Value()))
	}
	SimpleBandit(bandit, actions, 0.05)
}
