package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
)

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
	// existEmpty := false
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

	// if existEmpty {
	// 	fmt.Println(b.stone + "のターンです")
	// } else {
	// 	fmt.Println(b.stone + "ゲーム終了！")
	// 	b.Game = false
	// }
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
