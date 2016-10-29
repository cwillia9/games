package pente

import (
	"reflect"
	"testing"
)

func TestPlace(t *testing.T) {
	p := Init()

	var err error
	err = p.Place(10, 10)
	if err != nil {
		t.Error("Got unexpected error placing")
	}
	err = p.Place(10, 10)
	if err == nil {
		t.Error("Expected error, didn't get one")
	}

	p1 := Init()

	p1.board[10][10] = 1

	if p.String() != p1.String() {
		t.Error("Boards did not match")
	}

}

func TestCaptureLogic(t *testing.T) {
	p := Init()

	p.board[10][10] = 1
	p.board[10][11] = 2
	p.board[10][12] = 2
	p.board[10][13] = 1

	res := p.placeIncurCapture(10, 10)
	expect := []pair{
		pair{
			point{10, 11},
			point{10, 12},
		},
	}
	if !reflect.DeepEqual(res, expect) {
		t.Error("Expected:", expect, "got:", res)
	}

	expect = append([]pair{
		pair{
			point{11, 10},
			point{12, 10},
		},
	}, expect...)

	p.board[11][10] = 2
	p.board[12][10] = 2
	p.board[13][10] = 1

	res = p.placeIncurCapture(10, 10)
	if !reflect.DeepEqual(res, expect) {
		t.Error("Expected:", expect, "got:", res)
	}
}

func Test5InARowLogic(t *testing.T) {
	p := Init()

	p.board[10][10] = 1
	p.board[10][11] = 1
	p.board[10][12] = 1
	p.board[10][13] = 1
	p.board[10][14] = 1

	seq := p.placeIncur5seq(10, 11)
	if len(seq) != 5 {
		t.Error("Expected a sequence of 5")
	}

}
