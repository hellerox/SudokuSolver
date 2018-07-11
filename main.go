package main

import (
	"fmt"
	"os"
)

type cell struct {
	id     int
	value  int
	solved bool
	pVal   []int
}

type boardInt [9][9]int
type board [9][9]cell

func main() {

	bi := boardInt{
		{0, 0, 0, 2, 6, 0, 7, 0, 1},
		{6, 8, 0, 0, 7, 0, 0, 9, 0},
		{1, 9, 0, 0, 0, 4, 5, 0, 0},
		{8, 2, 0, 1, 0, 0, 0, 4, 0},
		{0, 0, 4, 6, 0, 2, 9, 0, 0},
		{0, 5, 0, 0, 0, 3, 0, 2, 8},
		{0, 0, 9, 3, 0, 0, 0, 7, 4},
		{0, 4, 0, 0, 5, 0, 0, 3, 6},
		{7, 0, 3, 0, 1, 8, 0, 0, 0},
	}

	board := bi.convertCellBoard()
	board.solve()
}

func (bi boardInt) convertCellBoard() board {
	var b board
	var cuad int
	var s bool = false

	for c := 0; c <= 8; c++ {
		for r := 0; r <= 8; r++ {
			if c >= 0 && c <= 2 && r <= 2 && r >= 0 {
				cuad = 1
			}
			if c >= 0 && c <= 2 && r <= 5 && r >= 3 {
				cuad = 2
			}
			if c >= 0 && c <= 2 && r <= 8 && r >= 6 {
				cuad = 3
			}
			if c >= 3 && c <= 5 && r <= 2 && r >= 0 {
				cuad = 4
			}
			if c >= 3 && c <= 5 && r <= 5 && r >= 3 {
				cuad = 5
			}
			if c >= 3 && c <= 5 && r <= 8 && r >= 6 {
				cuad = 6
			}
			if c >= 6 && c <= 8 && r <= 2 && r >= 0 {
				cuad = 7
			}
			if c >= 6 && c <= 8 && r <= 5 && r >= 3 {
				cuad = 8
			}
			if c >= 6 && c <= 8 && r <= 8 && r >= 6 {
				cuad = 9
			}

			if bi[c][r] > 9 {
				fmt.Println("Número mayor a 9, inválido.")
				os.Exit(3)
			}

			if bi[c][r] != 0 {
				s = true
			} else {
				s = false
			}

			b[c][r] = cell{id: cuad, value: bi[c][r], solved: s, pVal: []int{}}
		}
	}

	return b
}

func (b board) print(pval bool) {
	for c := 0; c <= 8; c++ {
		for r := 0; r <= 8; r++ {
			fmt.Print(b[c][r].value)
			if pval == true {
				fmt.Print(b[c][r].pVal)
			}
			fmt.Print(" ")
			if r == 8 {
				fmt.Println("\n")
			}
		}
	}
}

func (b board) findInRow(r int, val int) int {

	for i := 0; i < 8; i++ {
		if b[r][i].value == val {
			return i

		}
	}
	return -1
}

func (b board) findInColumn(c int, val int) int {

	for i := 0; i < 8; i++ {
		if b[i][c].value == val {
			return i
		}
	}
	return -1
}

func (b board) possibleValues() (board,int) {

	var changes int
	for c := 0; c < 9; c++ {
		for r := 0; r < 9; r++ {
			mc := make(map[int]int)
			mr := make(map[int]int)

			if b[c][r].solved == false {
				b[c][r].pVal = []int{}
				for i := 0; i < 9; i++ {
					if b[c][i].value != 0 {
						mr[b[c][i].value] = 1
					}
				}
				for i := 0; i < 9; i++ {
					if b[i][r].value != 0 {
						mc[b[i][r].value] = 1
					}
				}

				for i := 1; i <= 9; i++ {
					if _, ok := mc[i]; !(ok) {
						if _, ok := mr[i]; !(ok) {
							b[c][r].pVal = append(b[c][r].pVal, i)
						}
					}
				}

				if len(b[c][r].pVal) == 1 {
					b[c][r].solved = true
					b[c][r].value = b[c][r].pVal[0]
					b[c][r].pVal = []int{}
					changes++
				}
			}
		}
	}

	b.print(true)

	return b,changes
}


func (b board)solve() board {
	var b2 board
	var i1 int

	b2,i1=b.possibleValues()
	for{
		b2,i1=b2.possibleValues()
		fmt.Println(i1)
		if i1==0{break}
		}

	return b2
}