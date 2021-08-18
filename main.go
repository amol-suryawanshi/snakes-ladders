package main

import (
	"fmt"
	"os"
	"snakes-ladders/entity"
	"snakes-ladders/usecases"
)

func main() {
	fmt.Println("The Game has started")
	var char string
	var choice string
	var snakeCount int
	var ladderCount int
	//Get user input to select dice
	fmt.Printf("\n\t1. Normal\n\t2. Crooked\n\tSelect type of dice? : ")
	fmt.Scanln(&choice)
	var dice entity.Dice
	switch choice {
	case "1":
		dice = usecases.NewDice(entity.NormalDice)
	case "2":
		dice = usecases.NewDice(entity.CrookedDice)
	default:
		fmt.Println("Invalid choice")
		os.Exit(1)
	}

	fmt.Print("\n\tEnter Number of Snakes on board : ")
	fmt.Scanln(&snakeCount)
	snakes := make(map[int]int)
	for i := 0; i < snakeCount; i++ {
		var start, end int
		fmt.Printf("Enter Start and End Position of snake %d : ", i+1)
		fmt.Scanln(&start, &end)
		if start <= end {
			fmt.Println("Enter valid positions")
			i--
			continue
		}
		snakes[start] = end
	}

	fmt.Printf("\n\tEnter Number of Ladders on board : ")
	fmt.Scanln(&ladderCount)
	ladders := make(map[int]int)
	for i := 0; i < snakeCount; i++ {
		var start, end int
		fmt.Printf("Enter Start and End Position of Ladder %d : ", i+1)
		fmt.Scanln(&start, &end)
		if start >= end {
			fmt.Println("Enter valid positions")
			i--
			continue
		}
		ladders[start] = end
	}

	board := usecases.NewBoard(dice, usecases.NewPawn(), usecases.NewSnakes(snakes), usecases.NewLadders(ladders))

	fmt.Println("Board is Ready to play\n\n\tLets start!!!")
	for usecases.GetPos(&board.Pawn) >= 0 && usecases.GetPos(&board.Pawn) < 100 {
		var cRollNum int

		//Get user input to proceed for rolling dice
		fmt.Printf("Roll dice? (y/n): ")
		fmt.Scan(&char)
		switch char {
		case "y", "Y", "yes", "Yes", "YES":
			cRollNum = usecases.RollDice(board.Dice)
		case "n", "N", "no", "No", "NO":
			fmt.Printf("Quiting the game :-(")
			os.Exit(1)
		default:
			continue
		}

		//Move the pawn according to point got from dice rolling
		fmt.Printf("Current roll point: %d\n", cRollNum)
		if (usecases.GetPos(&board.Pawn) + cRollNum) <= 100 {
			usecases.SetPos(&board.Pawn, usecases.GetPos(&board.Pawn)+cRollNum)
		}
		fmt.Printf("Pawn position: %d\n", usecases.GetPos(&board.Pawn))

		//Verify that the current position of pawn is having Snake head or ladder foot
		//if current pawn pos is Snake's head pawn will be move to snake's tail
		//if current pawn pos is Ladder's foot pawn will be move to Ladder's top
		if v, ok := usecases.GetSnakesPos(board.Snakes, usecases.GetPos(&board.Pawn)); ok {
			fmt.Printf("Oops, bitten by snake back to: %d\n", v)
			usecases.SetPos(&board.Pawn, v)
		} else if v, ok := usecases.GetLaddersPos(board.Ladders, usecases.GetPos(&board.Pawn)); ok {
			fmt.Printf("Hurray, got the ladder climb to: %d\n", v)
			usecases.SetPos(&board.Pawn, v)
		}
	}
	if usecases.GetPos(&board.Pawn) == 100 {
		fmt.Printf("Game Completed Successfully :-)")
	}
}
