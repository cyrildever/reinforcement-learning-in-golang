package model

import (
	"fmt"
	"math"
	"strings"

	"gonum.org/v1/gonum/mat"
)

// State defines the interface for any state implementation.
type State interface {
	IsTerminal() bool
	Vector() mat.Vector
}

// StateValue is a map of expected value for state.
type StateValue map[State]float64

// Get ...
func (sv StateValue) Get(at State) float64 {
	return sv[at]
}

// Clone ...
func (sv StateValue) Clone() StateValue {
	cloned := make(StateValue, len(sv))
	for s, v := range sv {
		cloned[s] = v
	}
	return cloned
}

// Print mimics Python's numpy matrix printing
func (sv StateValue) Print(states []State, width int) (str string) {
	i := 0
	str = "[[\t"
	for _, s := range states {
		v := sv[s]
		if v >= 0 {
			str += " "
		} else {
			str += "-"
		}
		if v == 0 {
			str += "0.\t\t"
		} else {
			str += fmt.Sprintf("%.6f\t", math.Abs(v))
		}
		i++
		if i == len(sv) {
			str += "]"
		} else if i%width == 0 {
			str += "]\n [\t"
		}
	}
	str += "]"
	return
}

// ToPolicy transforms the state value to policy for action
func (sv StateValue) ToPolicy(actions map[Action]ActionFunc, states []State, width int) (map[State][]Action, string) {
	// Build policy
	policy := make(map[State][]Action, len(states))
	for _, s := range states {
		if s.IsTerminal() {
			policy[s] = []Action{}
		} else {
			values := make(map[Action]float64, len(actions))
			for a, f := range actions {
				next, _ := f(s, a)
				if next == s {
					values[a] = -math.MaxFloat64
				} else {
					values[a] = sv[next]
				}
			}
			var highest float64 = -math.MaxFloat64
			taken := make(map[float64][]Action)
			for a, v := range values {
				if v >= highest {
					highest = v
					if len(taken[highest]) > 0 {
						existing := taken[highest]
						existing = append(existing, a)
						taken[highest] = existing
					} else {
						taken[highest] = []Action{a}
					}
				}
			}
			policy[s] = taken[highest]
		}
	}

	// Create display
	i := 0
	str := "[[\t"
	for _, s := range states {
		as := policy[s]
		if len(as) == 0 {
			str += fmt.Sprintf("%-12s\t", "-")
		} else {
			names := []string{}
			for _, a := range as {
				names = append(names, a.GetName())
			}
			str += fmt.Sprintf("%-12s\t", strings.Join(names, ","))
		}
		i++
		if i == len(policy) {
			str += "]"
		} else if i%width == 0 {
			str += "]\n [\t"
		}
	}
	str += "]"
	return policy, str
}
