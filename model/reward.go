package model

// Reward is the value of the reward
type Reward float64

// Float64 ...
func (r Reward) Float64() float64 {
	return float64(r)
}

// ToValue ...
func (r Reward) ToValue() Value {
	return Value(float64(r))
}
