package pente

import (
	"bytes"
	"fmt"
)

const (
	boardSize = 19
)

type P struct {
	hasWon        bool
	board         [boardSize][boardSize]int
	currentPlayer int
	captured      [3]int
}

type pair struct {
	p0, p1 point
}

type point struct {
	row, col int
}

func (p *P) HasWon() bool {
	return p.hasWon
}

func (p *P) Winner() (int, error) {
	if p.hasWon {
		return p.currentPlayer, nil
	}
	return 0, ErrNoWinner
}

func (p *P) CurrentPlayer() int {
	return p.currentPlayer
}

func Init() *P {
	return &P{
		hasWon:        false,
		currentPlayer: 1,
	}
}

func (p *P) Place(row, col int) error {
	if p.hasWon {
		return ErrGameOver
	}

	if row >= boardSize || col >= boardSize {
		return ErrInvalidPlacement
	}

	if p.board[row][col] > 0 {
		return ErrInvalidPlacement
	}
	defer p.incrementPlayer()

	p.board[row][col] = p.currentPlayer

	pts := p.placeIncur5seq(row, col)
	if len(pts) >= 5 {
		p.hasWon = true
	}

	pairs := p.placeIncurCapture(row, col)
	p.removePairs(pairs)
	p.captured[p.currentPlayer] += len(pairs)
	if p.captured[p.currentPlayer] >= 5 {
		p.hasWon = true
	}

	return nil
}

func (p *P) placeIncur5seq(row, col int) []point {
	var subseq_0, subseq_1 []point

	seq := make([]point, 0)

	origin := point{row, col}

	subseq_0 = p.subSeqHelper(1, 0, origin)
	subseq_1 = p.subSeqHelper(-1, 0, origin)

	if (len(subseq_0) + len(subseq_1)) > 5 {
		return append(subseq_0, subseq_1[1:]...)
	}

	subseq_0 = p.subSeqHelper(0, 1, origin)
	subseq_1 = p.subSeqHelper(0, -1, origin)
	if (len(subseq_0) + len(subseq_1)) > 5 {
		return append(subseq_0, subseq_1[1:]...)
	}

	subseq_0 = p.subSeqHelper(1, 1, origin)
	subseq_1 = p.subSeqHelper(-1, -1, origin)
	if (len(subseq_0) + len(subseq_1)) > 5 {
		return append(subseq_0, subseq_1[1:]...)
	}

	subseq_0 = p.subSeqHelper(1, -1, origin)
	subseq_1 = p.subSeqHelper(-1, 1, origin)
	if (len(subseq_0) + len(subseq_1)) > 5 {
		return append(subseq_0, subseq_1[1:]...)
	}

	return seq
}

func (p *P) placeIncurCapture(row, col int) []pair {

	capturedPairs := make([]pair, 0)

	origin := point{row, col}

	var _pair *pair

	_pair = p.capturedHelper(1, 0, origin)
	capturedPairs = addIfNotNil(_pair, capturedPairs)
	_pair = p.capturedHelper(0, 1, origin)
	capturedPairs = addIfNotNil(_pair, capturedPairs)
	_pair = p.capturedHelper(-1, 0, origin)
	capturedPairs = addIfNotNil(_pair, capturedPairs)
	_pair = p.capturedHelper(0, -1, origin)
	capturedPairs = addIfNotNil(_pair, capturedPairs)
	_pair = p.capturedHelper(1, 1, origin)
	capturedPairs = addIfNotNil(_pair, capturedPairs)
	_pair = p.capturedHelper(-1, 1, origin)
	capturedPairs = addIfNotNil(_pair, capturedPairs)
	_pair = p.capturedHelper(1, -1, origin)
	capturedPairs = addIfNotNil(_pair, capturedPairs)
	_pair = p.capturedHelper(-1, -1, origin)
	capturedPairs = addIfNotNil(_pair, capturedPairs)

	return capturedPairs
}

func (p *P) removePairs(pairs []pair) {
	for _, pair := range pairs {
		p.board[pair.p0.row][pair.p0.col] = 0
		p.board[pair.p1.row][pair.p1.col] = 0
	}
}

func addIfNotNil(p *pair, lst []pair) []pair {
	if p != nil {
		lst = append(lst, *p)
	}
	return lst
}

func (p *P) isSurrounded(p0, p1, p2, p3 point) bool {
	return p.val(p0) == p.val(p3) && p.val(p1) == p.val(p2) && p.val(p0) != p.val(p1)
}

func (p *P) val(p0 point) int {
	return p.board[p0.row][p0.col]
}

func (p *P) String() string {
	var buf bytes.Buffer

	X := []byte("  X ")
	O := []byte("  O ")
	dash := []byte("  - ")
	nl := []byte("\n")

	buf.Write([]byte("    "))
	for i := 0; i < boardSize; i++ {
		buf.Write([]byte(fmt.Sprintf(" %2d ", i)))
	}
	buf.Write(nl)

	for i, r := range p.board {
		buf.Write([]byte(fmt.Sprintf(" %2d ", i)))
		for _, c := range r {
			switch c {
			case 1:
				buf.Write(X)
			case 2:
				buf.Write(O)
			default:
				buf.Write(dash)
			}
		}
		buf.Write(nl)
	}
	return buf.String()
}

func (p *P) incrementPlayer() {
	p.currentPlayer += 1
	if p.currentPlayer > 2 {
		p.currentPlayer = 1
	}
}

func (p *P) capturedHelper(rowinc, colinc int, origin point) *pair {
	rowEnd := origin.row + (rowinc * 3)
	colEnd := origin.col + (colinc * 3)

	if isOutOfBounds(rowEnd, colEnd) {
		return nil
	}

	var pts = new([3]point)
	for i := 1; i < 4; i++ {
		pts[i-1] = point{row: origin.row + (i * rowinc), col: origin.col + (i * colinc)}
	}
	if p.isSurrounded(origin, pts[0], pts[1], pts[2]) {
		return &pair{pts[0], pts[1]}
	}
	return nil
}

func (p *P) subSeqHelper(rowinc, colinc int, origin point) []point {
	var p_next point
	var rowval, colval int
	originval := p.val(origin)

	seq := make([]point, 1)
	seq[0] = origin
	for i := 1; ; i++ {
		rowval = origin.row + (i * rowinc)
		colval = origin.col + (i * colinc)

		if isOutOfBounds(rowval, colval) {
			break
		}

		p_next = point{row: rowval, col: colval}

		if p.val(p_next) == originval {
			seq = append(seq, p_next)
		} else {
			break
		}
	}
	return seq
}

func isOutOfBounds(row, col int) bool {
	if row < 0 || row >= boardSize {
		return true
	}

	if col < 0 || row >= boardSize {
		return true
	}

	return false
}
