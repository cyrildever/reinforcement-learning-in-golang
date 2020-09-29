package model

// Policy is the policy to evaluate/use
type Policy struct {
	// StateActions is the actual policy per state, ie. the possible action(s) for each state
	StateActions

	// Gamma is the discounting rate to use
	Gamma float64

	// Pi is a function that defines the probability of action a when being in state s: 𝛑(a|s)
	Pi func(Action, State) float64
}

// StateActions describe the possible actions per state, ie. the actual policy
type StateActions map[State][]Action
