package grid

type Line struct {
	from Vec
	to   Vec
}

func NewLine(v1, v2 Vec) *Line {
	return &Line{from: v1, to: v2}
}

func (l *Line) OnDraw(fn func(v Vec) error) error {
	if l.from.X == l.to.X {
		if l.from.Y > l.to.Y {
			l.from, l.to = l.to, l.from
		}
		for i := l.from.Y; i <= l.to.Y; i++ {
			if err := fn(Vec{l.from.X, i}); err != nil {
				return err
			}
		}
	} else {
		if l.from.X > l.to.X {
			l.from, l.to = l.to, l.from
		}
		for i := l.from.X; i <= l.to.X; i++ {
			if err := fn(Vec{i, l.from.Y}); err != nil {
				return err
			}
		}
	}
	return nil
}
