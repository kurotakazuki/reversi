package main

import (
	"fmt"
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

func TestBoard_allocableList(t *testing.T) {
	b := Board{}
	b.EMPTY = " "
	b.BLACK = "●"
	b.WHITE = "○"
	b.initialize()

	flippableWays := b.allocableList(b.WHITE, 3, 2)
	fmt.Println(flippableWays)
	expected := []Way{
		{0, 1}, // Up
	}

	if len(flippableWays) != len(expected) {
		t.Errorf("Expected %d flippable ways, but got %d", len(expected), len(flippableWays))
	}

	for i, way := range flippableWays {
		if way != expected[i] {
			t.Errorf("Expected way %d to be %v, but got %v", i, expected[i], way)
		}
	}
}

func TestBoard_turnStone(t *testing.T) {
	b := Board{}
	b.EMPTY = " "
	b.BLACK = "●"
	b.WHITE = "○"
	b.next = b.BLACK
	b.initialize()

	b.setStone(3, 5) // Put BLACK stone at (2,3)

	// Check if the stone at (3,3) turned from WHITE to BLACK
	if b.board[3][4] != b.BLACK {
		t.Errorf("Expected board[3][4] to be BLACK, but got %s", b.board[3][4])
	}
}

func TestBoardCount(t *testing.T) {
	var b Board
	b.EMPTY = " "
	b.BLACK = "●"
	b.WHITE = "○"

	blackCounter := 0
	whiteCounter := 0

	blackCounter = b.CountBlack(blackCounter, b.BLACK)
	if blackCounter != 1 {
		t.Errorf("CountBlack function is not working correctly. Expected 1 but got %d", blackCounter)
	}

	whiteCounter = b.CountWhite(whiteCounter, b.WHITE)
	if whiteCounter != 1 {
		t.Errorf("CountWhite function is not working correctly. Expected 1 but got %d", whiteCounter)
	}
}
