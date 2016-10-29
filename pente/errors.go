package pente

import (
	"errors"
)

var (
	ErrInvalidPlacement = errors.New("Invalid Placement")
	ErrGameOver         = errors.New("Game has been won")
	ErrNoWinner         = errors.New("Game has not yet been won")
)
