package logicop

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

// Break is implicit so no need to type it down
func EnumEvalImplicit(direction Direction) Direction {
	switch direction {
	case West:
		return direction
	case South:
		return direction
	case East:
		return direction
	case North:
		return direction
	}
	return 0
}

// Switch can also be explicit
func EnumEvalExplicit(direction Direction) Direction {
	switch {
	case direction == West:
		return direction
	case direction == South:
		return direction
	case direction == East:
		return direction
	case direction == North:
		return direction
	}
	return 0
}
