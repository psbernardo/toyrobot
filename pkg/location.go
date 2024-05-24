package pkg

type TableTop struct {
	X int
	Y int
}

func NewTableTop(X, Y int) TableTop {
	return TableTop{
		X: X,
		Y: Y,
	}
}

func (t TableTop) IsInsideTheTablePot(X, Y int) bool {
	if X < 0 || Y < 0 {
		return false
	}
	if X > (t.X - 1) {
		return false
	}

	if Y > (t.Y - 1) {
		return false
	}

	return true
}

type Location struct {
	X int
	Y int
}
