package model

// Value is the value of an action as a float number
type Value float64

// Values is the map
type Values map[Action]Value

// Inverse computes the inverse of the value, ie. `inv := 1/v`
func (v Value) Inverse() Value {
	return Value(1 / float64(v))
}

// UintToValue ...
func UintToValue(n uint) Value {
	return Value(float64(n))
}
