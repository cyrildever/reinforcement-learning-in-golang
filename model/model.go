package model

// Model sums up the characteristics of a model
type Model struct {
	Policy      Policy
	States      []State
	Probability ProbabilityFunc
	GridWidth   int // optional value for model representing some kind of grid/2D matrix
}

// ProbabilityFunc is a function to return the probability to end up in state s' and receive reward r starting in state s and picking action a:
// Pr = p(s',r|s,a)
type ProbabilityFunc func(sPrime State, r Reward, s State, a Action) float64
