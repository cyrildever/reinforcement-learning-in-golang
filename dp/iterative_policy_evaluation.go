package dp

import (
	"log"
	"math"
	"rl-algo/model"
	"rl-algo/utils"

	"gonum.org/v1/gonum/mat"
)

// IterativePolicyEvaluation implements the "iterative policy evaluation, for estimating V ≈ v𝛑" (page 75)
//
// It takes as arguments the input policy approximation function 𝛑 to use for evaluation and the algorithm parameter θ
// as well as the model definition to be used as a Markov Decision Process (MDP), ie. the different states of the
// environment including the terminal one (S+), the possible actions at each step and the corresponding probability function.
// It returns the expected (and potentially optimal) state value map for the evaluated policy.
func IterativePolicyEvaluation(pi model.Policy, theta float64, mdp model.Model) model.StateValue {
	if theta < 0 {
		log.Fatalln("invalid theta threshold")
	}
	if len(mdp.States) == 0 || len(mdp.Actions) == 0 {
		log.Fatalln("invalid empty model parameter(s)")
	}

	// Initialize
	V := make(model.StateValue, len(mdp.States))
	for _, s := range mdp.States {
		V[s] = 0
	}

	var delta float64 = math.MaxFloat64
	monitor := utils.LiveMonitor{Output: "state-value"}
	for delta >= theta {
		delta = 0
		previous := V.Clone()
		for _, s := range mdp.States {
			v := V[s]
			if s.IsTerminal() {
				V[s] = 0
			} else {
				V[s] = computeStateValue(s, pi, mdp, previous)
			}
			delta = math.Max(delta, math.Abs(float64(v-V[s])))
		}
		monitor.LogStateValues(V, mdp.States, mdp.GridWidth, 1, 2, 3, 10, 100, 500)
	}

	return V
}

func computeStateValue(s model.State, p model.Policy, mdp model.Model, sv model.StateValue) float64 {
	var sumPi float64 = 0
	var weightedRewards float64 = 0
	for _, a := range mdp.Actions {
		sPrime, r := a.ValueFunc()(s, a)
		sumPi += p.Pi(a, s)
		prob := mdp.Probabilities(sPrime, r, s, a)
		weightedRewards += prob * (r.Value() + p.Gamma*sv.Get(sPrime))
	}
	return sumPi * weightedRewards
}

//--- TEST

// State implementation for the gridworld
type gridworldState struct {
	X float64
	Y float64
}

func (s gridworldState) IsTerminal() bool {
	return (s.X == 0 && s.Y == 0) || (s.X == 3 && s.Y == 3)
}
func (s gridworldState) Vector() mat.Vector {
	return mat.NewVecDense(2, []float64{s.X, s.Y})
}

// Action implementation for the gridworld
type gridworldAction model.Action2D

func (a gridworldAction) ValueFunc() model.ActionFunc {
	return func(current model.State, _ model.Action) (next model.State, r model.Reward) {
		next = gridworldState{
			X: current.Vector().AtVec(0) + a.X,
			Y: current.Vector().AtVec(1) + a.Y,
		}
		if next.(gridworldState).X < 0 || next.(gridworldState).Y < 0 || next.(gridworldState).X > 3 || next.(gridworldState).Y > 3 {
			next = current
		}
		r = model.Reward(-1)
		return
	}
}
func (a gridworldAction) Vector() mat.Vector {
	return mat.NewVecDense(2, []float64{a.X, a.Y})
}

func TestIterativePolicyEvaluation() {
	log.Println("Testing IterativePolicyEvaluation (gridworld)... (Press Ctrl^C to end)")
	grid := []model.State{
		gridworldState{0, 0}, gridworldState{1, 0}, gridworldState{2, 0}, gridworldState{3, 0},
		gridworldState{0, 1}, gridworldState{1, 1}, gridworldState{2, 1}, gridworldState{3, 1},
		gridworldState{0, 2}, gridworldState{1, 2}, gridworldState{2, 2}, gridworldState{3, 2},
		gridworldState{0, 3}, gridworldState{1, 3}, gridworldState{2, 3}, gridworldState{3, 3},
	}
	var (
		LEFT  = gridworldAction{-1, 0}
		RIGHT = gridworldAction{1, 0}
		UP    = gridworldAction{0, 1}
		DOWN  = gridworldAction{0, -1}
	)
	actions := []model.Action{LEFT, RIGHT, UP, DOWN}
	policy := model.Policy{
		Actions: actions,
		Gamma:   1,
		Pi:      func(a model.Action, s model.State) float64 { return 0.25 },
	}
	mdp := model.Model{
		Actions: actions,
		States:  grid,
		Probabilities: func(sPrime model.State, r model.Reward, s model.State, a model.Action) float64 {
			return 1 / float64(len(actions))
		},
		GridWidth: 4,
	}
	_ = IterativePolicyEvaluation(policy, 1e-12, mdp)
}
