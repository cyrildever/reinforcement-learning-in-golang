package model

// Reward is the value of the reward
type Reward float64

// Float64 ...
func (r Reward) Value() Value {
	return Value(r)
}
