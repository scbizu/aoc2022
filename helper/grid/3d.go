package grid

type XYZVec struct {
	X int
	Y int
	Z int
}

func (v XYZVec) Distance(v2 XYZVec) int {
	return Abs(v.X, v2.X) + Abs(v.Y, v2.Y) + Abs(v.Z, v2.Z)
}

func (v XYZVec) Add(v2 XYZVec) XYZVec {
	return XYZVec{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func (v XYZVec) Neighbors6() []XYZVec {
	return []XYZVec{
		{X: v.X + 1, Y: v.Y, Z: v.Z},
		{X: v.X - 1, Y: v.Y, Z: v.Z},
		{X: v.X, Y: v.Y + 1, Z: v.Z},
		{X: v.X, Y: v.Y - 1, Z: v.Z},
		{X: v.X, Y: v.Y, Z: v.Z + 1},
		{X: v.X, Y: v.Y, Z: v.Z - 1},
	}
}
