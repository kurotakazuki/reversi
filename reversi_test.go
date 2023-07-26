package main

import (
	"testing"
)

func TestBoardInitialization(t *testing.T) {
	var b Board
	b.EMPTY = " "
	b.BLACK = "●"
	b.WHITE = "○"
	b.initialize()

	// Check if the game is running
	if b.game != true {
		t.Errorf("Expected game to be running but it was not")
	}

	// Check if the correct initial setup is done
	if b.board[3][3] != b.BLACK || b.board[3][4] != b.WHITE ||
		b.board[4][3] != b.WHITE || b.board[4][4] != b.BLACK {
		t.Errorf("Board not correctly initialized")
	}

	// Check if the stone and rev_stone are correctly set
	if b.stone != b.BLACK || b.rev_stone != b.WHITE {
		t.Errorf("Stones not correctly initialized")
	}
}

// func TestBoardCount(t *testing.T) {
// 	var b Board
// 	b.EMPTY = " "
// 	b.BLACK = "●"
// 	b.WHITE = "○"

// 	blackCounter := 0
// 	whiteCounter := 0

// 	blackCounter = b.CountBlack(blackCounter, b.BLACK)
// 	if blackCounter != 1 {
// 		t.Errorf("CountBlack function is not working correctly. Expected 1 but got %d", blackCounter)
// 	}

// 	whiteCounter = b.CountWhite(whiteCounter, b.WHITE)
// 	if whiteCounter != 1 {
// 		t.Errorf("CountWhite function is not working correctly. Expected 1 but got %d", whiteCounter)
// 	}
// }

// // Due to ShowBoard() doesn't return any values and it's output oriented, it's hard to test it properly.
// // We could, however, check if it runs without crashing
// func TestShowBoard(t *testing.T) {
// 	var b Board
// 	b.EMPTY = " "
// 	b.BLACK = "●"
// 	b.WHITE = "○"
// 	b.initialize()

// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Errorf("The code panicked")
// 		}
// 	}()

// 	b.ShowBoard()
// }
