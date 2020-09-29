package dp

import (
	"fmt"
	"log"
	"math"
	"rl-algo/model"
	"rl-algo/utils"

	"gonum.org/v1/gonum/mat"
)

// IterativePolicyEvaluation implements the "iterative policy evaluation, for estimating V ≈ v𝛑" (page 75)
//
// It takes as arguments the model definition to be used as a Markov Decision Process (MDP) (including the policy
// approximation function 𝛑 to use for evaluation, the different states (S+) and the probability function for each state)
// as well as the algorithm parameter θ.
// It returns the expected (and potentially optimal) state value map for the evaluated policy.
func IterativePolicyEvaluation(mdp model.Model, theta float64) model.StateValue {
	if theta < 0 {
		log.Fatalln("invalid theta threshold")
	}
	if len(mdp.States) == 0 || len(mdp.Policy.StateActions) == 0 {
		log.Fatalln("invalid empty model")
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
				V[s] = computeStateValue(s, mdp, previous)
			}
			delta = math.Max(delta, math.Abs(float64(v-V[s])))
		}
		monitor.LogStateValues(V, mdp.States, mdp.GridWidth, 1, 2, 3, 10, 100, 1000)
	}

	return V
}

// Compute value function: V(s) ← Σ 𝛑(a|s) Σ p(s',r|s,a)[r + 𝛾V(s')]
func computeStateValue(s model.State, mdp model.Model, previous model.StateValue) float64 {
	var sumPi float64 = 0
	var weightedRewards float64 = 0
	actions := mdp.Policy.StateActions[s]
	for _, a := range actions {
		sPrime, r := a.ValueFunc()(s, a)
		sumPi += mdp.Policy.Pi(a, s)
		prob := mdp.Probability(sPrime, r, s, a)
		weightedRewards += prob * (r.Value() + mdp.Policy.Gamma*previous.Get(sPrime))
	}
	return sumPi * weightedRewards
}

//--- TEST

func TestIterativePolicyEvaluation() {
	log.Println("Testing IterativePolicyEvaluation (gridworld)... (Press Ctrl^C to end)")
	grid := []model.State{
		gridworldState{0, 0}, gridworldState{1, 0}, gridworldState{2, 0}, gridworldState{3, 0},
		gridworldState{0, 1}, gridworldState{1, 1}, gridworldState{2, 1}, gridworldState{3, 1},
		gridworldState{0, 2}, gridworldState{1, 2}, gridworldState{2, 2}, gridworldState{3, 2},
		gridworldState{0, 3}, gridworldState{1, 3}, gridworldState{2, 3}, gridworldState{3, 3},
	}
	actions := []model.Action{
		gridworldAction{"left", -1, 0},
		gridworldAction{"right", 1, 0},
		gridworldAction{"up", 0, -1},
		gridworldAction{"down", 0, 1},
	}
	random := make(model.StateActions, len(grid))
	for _, s := range grid {
		if s.IsTerminal() {
			random[s] = []model.Action{} // NO ACTION
		} else {
			random[s] = actions
		}
	}
	deterministic := model.Policy{
		StateActions: random,
		Gamma:        1,
		Pi:           func(a model.Action, s model.State) float64 { return 0.25 },
	}
	mdp := model.Model{
		Policy: deterministic,
		States: grid,
		Probability: func(sPrime model.State, r model.Reward, s model.State, a model.Action) float64 {
			return 1 / float64(len(actions))
		},
		GridWidth: 4,
	}
	theta := 1e-12
	optimal := IterativePolicyEvaluation(mdp, theta)

	functions := make(map[model.Action]model.ActionFunc, len(actions))
	for _, a := range actions {
		functions[a] = a.ValueFunc()
	}
	_, displayOptimalPolicy := optimal.ToPolicy(functions, grid, mdp.GridWidth)
	fmt.Printf("\nPolicy result at θ = %.1e:\n%s\n\n%s\n", theta, optimal.Print(grid, mdp.GridWidth), displayOptimalPolicy)
}

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

func (a gridworldAction) GetName() string {
	return a.Name
}
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
