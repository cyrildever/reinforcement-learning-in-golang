package model

import (
	"gonum.org/v1/gonum/mat"
)

// State defines the interface for any state implementation.
type State interface {
	IsTerminal() bool
	Vector() mat.Vector
}
