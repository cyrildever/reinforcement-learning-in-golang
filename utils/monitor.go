package utils

import (
	"fmt"
	"log"
	"rl-algo/model"
)

// LiveMonitor ...
type LiveMonitor struct {
	Output  string // The name of the followed value, eg. "reward"
	steps   float64
	average float64
}

// ComputeAndLog takes a value at the end of a step and log the new average
func (m *LiveMonitor) ComputeAndLog(v float64) {
	m.steps += 1
	m.average = ((m.steps-1)*m.average + v) / m.steps
	log.Println(fmt.Sprintf("average_%s=%.6f;steps=%d", m.Output, m.average, int(m.steps)))
}

// LogStateValues displays the passed states and their respective values only at the passed steps.
func (m *LiveMonitor) LogStateValues(sv model.StateValue, forStates []model.State, width int, displaySteps ...int) {
	if width == 0 {
		width = 8
	}
	m.steps += 1
	if contains(displaySteps, int(m.steps)) {
		log.Println("====================", m.Output)
		log.Println(fmt.Sprintf("step=%d", int(m.steps)))
		fmt.Println(sv.Print(forStates, width))
	}
}

func contains(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
