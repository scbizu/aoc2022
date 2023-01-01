package seq

import "testing"

var cseq = NewCircular([]int{1, 2, -3, 3, -2, 0, 4})

func TestMove(t *testing.T) {
	cseq.Move(0, 1).Move(0, 2).Move(1, -3).Move(2, 3).Move(2, -2).Move(3, 0).Move(5, 4)
	t.Log(cseq.List())
}
