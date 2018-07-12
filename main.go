package main

import (
	"fmt"
	"os"
	"time"
)

type cell struct {
	cuad   int
	value  int
	solved bool
	pVal   []int
}

type boardInt [9][9]int

type board struct {
	boardd [9][9]cell
	padre  int
}

func main() {

	/*bi := boardInt{
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
	bi := boardInt{
		{0, 0, 0, 6, 0, 0, 4, 0, 0},
		{7, 0, 0, 0, 0, 3, 6, 0, 0},
		{0, 0, 0, 0, 9, 1, 0, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 5, 0, 1, 8, 0, 0, 0, 3},
		{0, 0, 0, 3, 0, 6, 0, 4, 5},
		{0, 4, 0, 2, 0, 0, 0, 6, 0},
		{9, 0, 3, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 1, 0, 0},
	}

	bi := boardInt{
		{0, 2, 0, 6, 0, 8, 0, 0, 0},
		{5, 8, 0, 0, 0, 9, 7, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{3, 7, 0, 0, 0, 0, 5, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 8, 0, 0, 0, 0, 1, 3},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 9, 8, 0, 0, 0, 3, 6},
		{0, 0, 0, 3, 0, 6, 0, 9, 0},
	}

	bi := boardInt{
		{0, 0, 0, 4, 8, 0, 2, 0, 0},
		{9, 0, 0, 7, 2, 0, 0, 1, 8},
		{8, 0, 0, 0, 0, 6, 0, 0, 0},
		{6, 0, 0, 0, 0, 2, 0, 8, 0},
		{4, 0, 0, 0, 0, 0, 0, 0, 7},
		{0, 9, 0, 1, 0, 0, 0, 0, 2},
		{0, 0, 0, 8, 0, 0, 0, 0, 5},
		{7, 8, 0, 0, 4, 3, 0, 0, 6},
		{0, 0, 6, 0, 5, 7, 0, 0, 0},
	}*/

	bi := boardInt{
		{0, 0, 0, 8, 1, 9, 5, 0, 3},
		{0, 0, 6, 0, 0, 0, 0, 0, 8},
		{0, 0, 0, 3, 0, 6, 2, 0, 0},
		{0, 0, 3, 0, 0, 0, 0, 2, 6},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 4, 0, 0, 0, 0, 8, 0, 0},
		{0, 0, 2, 4, 0, 8, 0, 0, 0},
		{9, 0, 0, 0, 0, 0, 4, 0, 0},
		{3, 0, 4, 7, 2, 1, 0, 0, 0},
	}

	defer elapsed("SudokuSolver")()

	board := bi.convertCellBoard()
	fmt.Println("Board Inicial")
	board.print(true)
	bs := board.simpleSolve()
	fmt.Println("---------- Resultado final ----------")
	bs.solve().print(false)

}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func (bi boardInt) convertCellBoard() board {
	var bb board
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

			bb.boardd[c][r] = cell{cuad: cuad, value: bi[c][r], solved: s, pVal: []int{}}
		}
	}

	return bb
}

func (bb board) print(pval bool) {
	b := bb.boardd
	fmt.Println("------------------------------------------------")
	for c := 0; c <= 8; c++ {
		for r := 0; r <= 8; r++ {
			fmt.Print(b[c][r].value)
			if pval == true {
				fmt.Print(b[c][r].pVal)
				fmt.Print(b[c][r].solved)
			}
			fmt.Print("  ")
			if r == 8 {
				fmt.Println("\n")
			}
		}
	}
	fmt.Println("------------------------------------------------ \n")
}

func (bb board) possibleValues() (board, int) {
	b := bb.boardd
	var changes int
	for c := 0; c < 9; c++ {
		for r := 0; r < 9; r++ {
			mc := make(map[int]int)
			mr := make(map[int]int)
			mg := make(map[int]int)

			//Si está resuelto se lo brinca (resuelto as in de entrada)
			if b[c][r].solved == false {

				//Revisa números existentes en columnas
				for i := 0; i < 9; i++ {
					if b[c][i].value != 0 {
						mr[b[c][i].value] = 1
					}
				}

				//Revisa números existentes el rows
				for i := 0; i < 9; i++ {
					if b[i][r].value != 0 {
						mc[b[i][r].value] = 1
					}
				}

				//Revisa dentro de su cubo
				for cx := 0; cx < 9; cx++ {
					for rx := 0; rx < 9; rx++ {
						if b[c][r].cuad == b[cx][rx].cuad && c != cx && c != rx {
							mg[b[cx][rx].value] = 1
						}
					}
				}

				b[c][r].pVal = []int{}
				for i := 1; i <= 9; i++ {
					if _, ok := mc[i]; !(ok) {
						if _, ok := mr[i]; !(ok) {
							if _, ok := mg[i]; !(ok) {
								if len(b[c][r].pVal) == 0 || b[c][r].pVal[0] < i {
									//fmt.Println("Agrego: ", c,r,i, b[c][r].pVal)
									b[c][r].pVal = append(b[c][r].pVal, i)
								}
							}
						}
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

	//b.print(true)
	bb.boardd = b
	return bb, changes
}

func (b board) simpleSolve() board {
	var b2 board
	var i1 int

	b2, i1 = b.possibleValues()
	for {
		b2, i1 = b2.possibleValues()
		fmt.Println("Realice cambios en esta: ", i1)
		if i1 == 0 {
			break
		}
	}
	fmt.Println("Simple Solve----")
	b2.print(true)
	return b2
}

func (pbb board) solve() (wb board) {
	colb := []board{}
	colb = append(colb, pbb)
	var b3 board
	i := 0

	for {
		cx, rx := unsolved(colb[i])
		//fmt.Println("cxrx", cx, rx)

		if cx ==-1 && rx==-1{
			fmt.Println("Posible final")
			colb[i].print(false)
			colb[i].validate()
			fmt.Println("Iteracion", i)
			fmt.Println("Pasos encontrados",len(colb))
			wb.boardd=colb[i].boardd
			return wb
		}

		for _, val := range (colb[i]).boardd[cx][rx].pVal {
			if val != 0 {
				b3 = colb[i]
				b3.padre = i
				b3.boardd[cx][rx].solved = true
				b3.boardd[cx][rx].value = val
				b3, _ = b3.possibleValues()
				colb = append(colb, b3)
			}
		}
		//fmt.Println("--")
		//fmt.Println("Me faltan por resolver:", colb[i].validate())
		//fmt.Println("Iteracion", i)
		i++
	}
}

func (bb board) validate() int {
	var fs int
	b := bb.boardd

	for c := 0; c < 9; c++ {
		sum := 0
		for r := 0; r < 9; r++ {
			sum += b[c][r].value

			if b[c][r].solved == false {
				fs++

			}
		}

		if sum!=45 {
			fmt.Println("\n \n Resultado o Sudoku inválido")
			os.Exit(3)
		}
	}
	return fs
}

func unsolved(bb board) (int, int) {
	b := bb.boardd
	for c := 0; c < 9; c++ {
		for r := 0; r < 9; r++ {
			if b[c][r].solved == false {
				return c, r
			}
		}
	}
	return -1, -1
}
