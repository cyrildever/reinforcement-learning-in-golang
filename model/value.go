package model

// Value is the value of an action as a float number
type Value float64

// Float64 ...
func (v Value) Float64() float64 {
	return float64(v)
}

// Inverse computes the inverse of the value, ie. `inv := 1/v`
func (v Value) Inverse() Value {
	return Value(1 / float64(v))
}

// UintToValue ...
func UintToValue(n uint) Value {
	return Value(float64(n))
}
