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

func (b *Board) ShowBoard() {
	cntBlack := 0
	cntWhite := 0

	fmt.Println(" |0|1|2|3|4|5|6|7|")
	fmt.Println("――――――――――――――")
	for i, row := range b.board {
		fmt.Printf("%d|", i)
		for _, s := range row {
			fmt.Print(s + "|")

			if s == b.EMPTY {
				//existEmpty = true
			} else if s == b.BLACK {
				cntBlack++
			} else if s == b.WHITE {
				cntWhite++
			}
		}
		fmt.Println()
		fmt.Println("――――――――――――――")
	}

	fmt.Println(b.BLACK+":", cntBlack)
	fmt.Println(b.WHITE+":", cntWhite)
	fmt.Println("――――――――――――――")

}

func (b *Board) allocable(rock string, x int, y int) bool {
	// 空でなければ置けない
	if !b.isEmpty {
		return false
	}
	//

	// 全方位の計算
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {

			// 真ん中方向除外
			if i == 0 && j == 0 {
				continue
			}
			// 盤面外
			if x+i < 0 || x+i >= 8 || y+j < 0 || y+j >= 8 {
				continue
			}

			if board[y+j][x+i] != color[*other] {
				continue
			}

			for s := 2; s < 8; s++ {
				if x+i*s >= 0 && x+i*s < 8 && y+j*s >= 0 && y+j*s < 8 {

					if board[y+j*s][x+i*s] == nil {
						break
					}

					if board[y+j*s][x+i*s] == color[player] {
						return true
					}
				}
			}
		}
	}
	return false
}

func (b *Board) isEmpty(x int, y int) bool {
	if b.board[x][y] == b.EMPTY {
		return true
	}
	return false
}
func (b *Board) isOut(x int, y int) bool {
	if x < 0 || 7 < x || y < 0 || 7 < y {
		return true
	}
	return false
}

func main() {

	var b Board

	b.EMPTY = " "
	b.BLACK = "●"
	b.WHITE = "○"
	b.stone = ""
	b.initialize()
	b.ShowBoard()

}
