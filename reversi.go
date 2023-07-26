package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	next      string
	stone     string
	rev_stone string
	// スキップした数
	skipNum int
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

	b.skipNum = 0
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
	existEmpty := true
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

	if existEmpty {
		b.skipNext()
	} else {
		b.game = false
	}
}

func (b *Board) skip() bool {

	// 全ての座標で配置できるか確認
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.isEmpty(i, j) {
				if b.allocable(b.next, i, j) {
					b.skipNum = 0
					return false
				}
			}
		}
	}
	b.skipNum += 1
	return true
}

func (b *Board) setStone(x int, y int) bool {
	// 置けるかの確認
	if !b.allocable(b.next, x, y) {
		return false
	}
	// おける場合：ひっくり返す
	b.board[x][y] = b.next
	b.turnStone(b.next, x, y)
	return true
}

type Way struct {
	i, j int
}

func (b *Board) allocableList(stone string, x int, y int) []Way {
	var flippableWays []Way
	// // 空でなければ置けない
	// if !b.isEmpty(x, y) {
	// 	return flippableWays
	// }

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
		// fmt.Println("way: ", way)
		b.turnStoneOfWay(stone, x, y, way)
	}
}

func (b *Board) turnStoneOfWay(stone string, x int, y int, way Way) {
	i := way.i
	j := way.j
	for s := 1; s < 8; s++ {
		xis := x + i*s
		yjs := y + j*s
		// fmt.Println("ひっくり返す座標", xis, yjs)
		// 同じ種類の石の場合、配置の終了
		if b.board[xis][yjs] == stone {
			break
		}
		b.board[xis][yjs] = stone
	}
}

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

// 勝敗を判定する関数
func (b *Board) determineWinner() string {
	cntBlack := 0
	cntWhite := 0

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			cntBlack = b.CountBlack(cntBlack, b.board[i][j])
			cntWhite = b.CountWhite(cntWhite, b.board[i][j])
		}
	}

	if cntBlack > cntWhite {
		return fmt.Sprintf("%s が %d 対 %d で勝利しました！", b.BLACK, cntBlack, cntWhite)
	} else if cntWhite > cntBlack {
		return fmt.Sprintf("%s が %d 対 %d で勝利しました！", b.WHITE, cntWhite, cntBlack)
	} else {
		return "引き分けです！"
	}
}

func (b *Board) skipNext() {
	if b.next == b.BLACK {
		b.next = b.WHITE
	} else if b.next == b.WHITE {
		b.next = b.BLACK
	}
}

func main() {

	var b Board
	// b.game = false
	//b.board =
	b.EMPTY = " "
	b.BLACK = "●"
	b.WHITE = "○"
	b.next = b.WHITE
	//b.rev_stone = ""
	b.initialize()
	b.ShowBoard()

	// コンソールからの入力を受け付ける
	scanner := bufio.NewScanner(os.Stdin)

	// ゲーム実行中フラグがtrueのあいだループする
	for b.game {
		if b.skip() {
			// コマの交代
			fmt.Println(b.next + "はどこにも配置できません。スキップします。")
			// 二人ともスキップした場合、終了
			if b.skipNum == 2 {
				break
			}
			b.skipNext()
			continue
		}
		fmt.Println(b.next + "のターンです")
		fmt.Print("駒をおくx座標を入力してください:")
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())

		fmt.Print("駒をおくy座標を入力してください:")
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())

		fmt.Println("x: ", x)
		fmt.Println("y: ", y)
		if !b.setStone(x, y) {
			continue
		}

		b.ShowBoard()

	}
	// 勝敗を判定する
	result := b.determineWinner()
	fmt.Println(result)
	fmt.Println("ゲーム終了！")

	// scanner.Close()
}
