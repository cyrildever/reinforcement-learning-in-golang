package model

// Reward is the value of the reward
type Reward float64

// Value ...
func (r Reward) Value() float64 {
	return float64(r)
}
