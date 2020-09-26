package utils

import (
	"fmt"
	"log"
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
