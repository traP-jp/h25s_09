package utils

import "math/rand/v2"

type Bug struct {
	Name        string
	Probability float64
}

var (
	BackendBugs = map[int]Bug{}
)

func DetermineDispatchBug(id int) (Bug, bool) {
	if bug, exists := BackendBugs[id]; exists {
		if bug.Probability >= 1.0 || (bug.Probability > 0 && rand.Float64() < bug.Probability) {
			return bug, true
		}
	}
	return Bug{}, false
}
