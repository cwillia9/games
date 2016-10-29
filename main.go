package main

import (
	"fmt"
	"github.com/cwillia9/games/pente"
)

func main() {
	p := pente.Init()

	fmt.Println(p)
	fmt.Println("Player 1's turn")

	var row, col int
	for !p.HasWon() {
		_, err := fmt.Scanf("%d %d\n", &row, &col)
		fmt.Println("row:", row, "col:", col)
		if err != nil {
			continue
		}

		err = p.Place(row, col)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(p)
		fmt.Println("Player", p.CurrentPlayer(), "'s turn")
	}

	fmt.Println("Game over!")
	fmt.Println("Player", p.CurrentPlayer(), "wins")
}
