package main

import "testing"

func TestFizzBuzz01(t *testing.T) {
    cntBlack, cntWhite := ShowBoard()
    fmt.Print(str)
	if cntBlack != "2" {
		t.Error("Test01 is failed")
	}
}

// func TestFizzBuzz01(t *testing.T) {
// 	str := fizzbuzz(1)
// 	if str != "1" {
// 		t.Error("Test01 is failed")
// 	}
// }

// func TestFizzBuzz01(t *testing.T) {
// 	str := fizzbuzz(1)
// 	if str != "1" {
// 		t.Error("Test01 is failed")
// 	}
// }