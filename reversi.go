package main

import "fmt"

type Board struct {
	// ゲーム実行中フラグ
	game bool
	// オセロ版に対応した多次元配列
	board [8][8]string
	// 定数
	EMPTY string
	BLACK string
	WHITE string
	// 変数
	stone     string
	rev_stone string
}

func (b *Board) initialize() {
	// オセロ版の要素を全てクリアする
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.board[i][j] = b.EMPTY
		}
	}

	// 初期状態の配置
	b.board[3][3] = b.BLACK
	b.board[3][4] = b.WHITE
	b.board[4][3] = b.WHITE
	b.board[4][4] = b.BLACK

	// 次うつ駒の色を指定
	b.stone = b.BLACK
	b.rev_stone = b.WHITE

	// ゲーム実行中フラグ
	b.game = true
}

func (b *Board) CountBlack(cntBlack int, s string) int {
	if s == b.EMPTY {
		//existEmpty = true
	} else if s == b.BLACK {
		cntBlack++
	}

	return cntBlack
}

func (b *Board) CountWhite(cntWhite int, s string) int {
	if s == b.EMPTY {
		//existEmpty = true
	} else if s == b.WHITE {
		cntWhite++
	}

	return cntWhite
}

func (b *Board) ShowBoard() {
	// existEmpty := false
	cntBlack := 0
	cntWhite := 0

	fmt.Println(" |0|1|2|3|4|5|6|7|")
	fmt.Println("――――――――――――――")
	for i, row := range b.board {
		fmt.Printf("%d|", i)
		for _, s := range row {
			fmt.Print(s + "|")
			cntBlack = b.CountBlack(cntBlack, s)
			cntWhite = b.CountWhite(cntWhite, s)

		}
		fmt.Println()
		fmt.Println("――――――――――――――")
	}

	fmt.Println(b.BLACK+":", cntBlack)
	fmt.Println(b.WHITE+":", cntWhite)
	fmt.Println("――――――――――――――")

	// if existEmpty {
	// 	fmt.Println(b.stone + "のターンです")
	// } else {
	// 	fmt.Println(b.stone + "ゲーム終了！")
	// 	b.Game = false
	// }

}

type Way struct {
	i, j int
}

func (b *Board) allocableList(stone string, x int, y int) []Way {
	var flippableWays []Way
	// 空でなければ置けない
	if !b.isEmpty(x, y) {
		return flippableWays
	}

	// 全方位の計算
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {

			// 真ん中方向は除外
			if i == 0 && j == 0 {
				continue
			}
			xi := x + i
			yj := y + j
			// 盤面外
			if isOut(xi, yj) {
				continue
			}
			// 隣の石が同じ種類の石の場合は、配置できない
			if b.board[xi][yj] == stone {
				continue
			}

			for s := 2; s < 8; s++ {
				// 盤面外のときは、配置できない
				if isOut(x+i*s, y+j*s) {
					break
				}
				// 空のときは、配置できない
				if b.isEmpty(x+i*s, y+j*s) {
					break
				}
				// 同じ種類の石の場合、配置できる
				if b.board[x+i*s][y+j*s] == stone {
					flippableWays = append(flippableWays, Way{i, j})
				}
			}
		}
	}
	return flippableWays
}

func (b *Board) turnStone(stone string, x int, y int) {
	for _, way := range b.allocableList(stone, x, y) {
		b.turnStoneOfWay(stone, x, y, way)
	}
}

func (b *Board) turnStoneOfWay(stone string, x int, y int, way Way) {
	i := way.i
	j := way.j
	for s := 1; s < 8; s++ {
		xis := x + i*s
		yjs := y + j*s
		b.board[xis][yjs] = stone
		// 同じ種類の石の場合、配置できる
		if b.board[xis][yjs] == stone {
			break
		}
	}
}

// func (b *Board) turnLeftUp(x int, y int) {
// 	if y > 1 && x > 1 {
// 		// となりの駒
// 		next := b.board[y-1][x-1]

// 		// となりの駒が裏駒の場合
// 		if next == b.rev_stone {
// 			// さらにその一つとなりから順に確認
// 			for i := 2; true; i++ {
// 				if x-i < 0 || y-i < 0 || b.board[y-i][x-i] == b.EMPTY {
// 					// 駒がない場合終了
// 					break
// 				} else if b.board[y-i][x-i] == b.stone {
// 					// 自駒の場合
// 					// あいだの駒をすべて自駒にひっくりかえす
// 					for t := 1; t < i; t++ {
// 						// 配列の要素を上書き
// 						b.board[y-t][x-t] = b.stone
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
// }

// func (b *Board) TurnUp(x, y int) {
// 	if y > 1 {
// 		next := b.Stones[y-1][x]

// 		if next == b.RevStone {
// 			for i := 2; ; i++ {
// 				if y-i < 0 || b.Stones[y-i][x] == Empty {
// 					break
// 				} else if b.Stones[y-i][x] == b.Stone {
// 					for t := 1; t < i; t++ {
// 						b.Stones[y-t][x] = b.Stone
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
// }

// func (b *Board) TurnRightUp(x, y int) {
// 	if y > 1 && x < 6 {
// 		next := b.Stones[y-1][x+1]

// 		if next == b.RevStone {
// 			for i := 2; ; i++ {
// 				if x+i > 7 || y-i < 0 || b.Stones[y-i][x+i] == Empty {
// 					break
// 				} else if b.Stones[y-i][x+i] == b.Stone {
// 					for t := 1; t < i; t++ {
// 						b.Stones[y-t][x+t] = b.Stone
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
// }

// func (b *Board) TurnDown(x, y int) {
// 	if y < 6 {
// 		next := b.Stones[y+1][x]

// 		if next == b.RevStone {
// 			for i := 2; ; i++ {
// 				if y+i > 7 || b.Stones[y+i][x] == Empty {
// 					break
// 				} else if b.Stones[y+i][x] == b.Stone {
// 					for t := 1; t < i; t++ {
// 						b.Stones[y+t][x] = b.Stone
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
// }

// func (b *Board) TurnRight(x, y int) {
// 	if x < 6 {
// 		next := b.Stones[y][x+1]

// 		if next == b.RevStone {
// 			for i := 2; ; i++ {
// 				if x+i >
// // 同様に他の方向のひっくり返す処理を実装する
// // turnUp, turnRightUp, turnLeft, turnRight, turnLeftDown, turnDown, turnRightDown

func (b *Board) isEmpty(x int, y int) bool {
	if b.board[x][y] == b.EMPTY {
		return true
	}
	return false
}
func isOut(x int, y int) bool {
	if x < 0 || 7 < x || y < 0 || 7 < y {
		return true
	}
	return false
}
func (b *Board) allocable(stone string, x int, y int) bool {
	// 空でなければ置けない
	if !b.isEmpty(x, y) {
		return false
	}

	// 全方位の計算
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {

			// 真ん中方向は除外
			if i == 0 && j == 0 {
				continue
			}
			xi := x + i
			yj := y + j
			// 盤面外
			if isOut(xi, yj) {
				continue
			}
			// 隣の石が同じ種類の石の場合は、配置できない
			if b.board[xi][yj] == stone {
				continue
			}

			for s := 2; s < 8; s++ {
				// 盤面外のときは、配置できない
				if isOut(x+i*s, y+j*s) {
					break
				}
				// 空のときは、配置できない
				if b.isEmpty(x+i*s, y+j*s) {
					break
				}
				// 同じ種類の石の場合、配置できる
				if b.board[x+i*s][y+j*s] == stone {
					return true
				}
			}
		}
	}
	return false
}

func main() {

	var b Board
	// b.game = false
	//b.board =
	b.EMPTY = " "
	b.BLACK = "●"
	b.WHITE = "○"
	b.stone = ""
	//b.rev_stone = ""
	b.initialize()
	b.ShowBoard()

	// // コンソールからの入力を受け付ける
	// scanner := bufio.NewScanner(os.Stdin)

	// // ゲーム実行中フラグがtrueのあいだループする
	// for game {
	// 	fmt.Print("駒をおくx座標を入力してください:")
	// 	scanner.Scan()
	// 	x, _ := strconv.Atoi(scanner.Text())

	// 	fmt.Print("駒をおくy座標を入力してください:")
	// 	scanner.Scan()
	// 	y, _ := strconv.Atoi(scanner.Text())

	// 	setStone(x, y)
	// }

	// scanner.Close()
}
