package sequence

// Direction is an alias for order pattern.
type Direction string

const (
	// Increasing order
	Increasing = Direction("increasing")
	// Decreasing order
	Decreasing = Direction("decreasing")
	// Stable order
	Stable = Direction("stable")
)

// Analyze detects the direction of sequence.
func Analyze(input []int) Direction {
	if len(input) <= 1 {
		return Stable
	}

	seen := make(map[Direction]struct{})
	last := input[0]

	for i := 1; i < len(input); i++ {
		if input[i] == last {
			return Stable
		}

		if input[i] > last {
			seen[Increasing] = struct{}{}
		} else {
			seen[Decreasing] = struct{}{}
		}

		last = input[i]
	}

	_, iok := seen[Increasing]
	_, dok := seen[Decreasing]

	if iok && dok {
		return Stable
	}

	if iok {
		return Increasing
	}

	return Decreasing
}
