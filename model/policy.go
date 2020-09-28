package model

// Policy is the policy to evaluate/use
type Policy struct {
	// Actions is the list of possible actions to evaluate
	Actions []Action

	// Gamma is the discounting rate to use
	Gamma float64

	// Pi is a function that defines the probability of action a when being in state s: 𝛑(a|s)
	Pi func(Action, State) float64
}
