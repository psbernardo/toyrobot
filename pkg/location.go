package pkg

type TableTop struct {
	X int
	Y int
}

func NewTableTop(X, Y int) *TableTop {
	return &TableTop{
		X: X,
		Y: Y,
	}
}

func (t TableTop) IsInsideTheTablePot(X, Y int) bool {
	if X < 0 || Y < 0 {
		return false
	}
	if X > t.X {
		return false
	}

	if Y > t.Y {
		return false
	}

	return true
}

type Location struct {
	X int
	Y int
}

// it defends on direction we if we will add or deduct
// NORTH MOVE(Y)++
// SOUTH MOVE(Y)--

// EAST MOVE  (X)++)
// WEST MOVE  (X)--
func NewMoveDirections(tabletop *TableTop) map[string]MoveLocationFunc {
	return map[string]MoveLocationFunc{
		North: MoveNorth(tabletop),
		East:  MoveEast(tabletop),
		West:  MoveWest(),
		South: MoveSouth(),
	}

}

func MoveNorth(tabletop *TableTop) func(l *Location) {

	return func(l *Location) {
		if l.Y < tabletop.Y {
			l.Y++
		}

	}
}

func MoveSouth() func(l *Location) {
	return func(l *Location) {
		if l.Y > 0 {
			l.Y--
		}

	}
}

func MoveEast(tabletop *TableTop) func(l *Location) {
	return func(l *Location) {
		if l.X < tabletop.X {
			l.X++
		}

	}
}

func MoveWest() func(l *Location) {
	return func(l *Location) {
		if l.X > 0 {
			l.X--
		}

	}
}
