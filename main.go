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

type board [9][9]cell


func main() {
	/*
	//Simpler board, solved without Solve()
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
	}*/

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

	/*

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
		}


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
		}*/

	defer elapsed("Sudoku Solver")()

	board := bi.convertCellBoard()
	fmt.Println("Board Inicial")
	board.print(false)
	bs := board.simpleSolve()
	fmt.Println("\n\n ---------- Resultado final ----------")
	bs.solve().print(false)

}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
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

			b[c][r] = cell{cuad: cuad, value: bi[c][r], solved: s, pVal: []int{}}
		}
	}

	return b
}

func (b board) print(pval bool) {
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

func (b board) possibleValues() (board, int) {
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

	return b, changes
}

func (b board) simpleSolve() board {
	var b2 board
	var i1 int

	b2, i1 = b.possibleValues()
	
	for {
		b2, i1 = b2.possibleValues()
		if i1 == 0 {
			break
		}
	}

	return b2
}

func (pb board) solve() (wb board) {
	col := []board{}
	col = append(col, pb)
	var b3 board
	i := 0

	for {
		cx, rx := unsolved(col[i])
		if cx ==-1 && rx==-1{
			fmt.Println("Posible final")
			col[i].print(false)
			col[i].validate()
			fmt.Println("Iteracion", i)
			fmt.Println("Pasos encontrados",len(col))
			wb=col[i]
			return wb
		}

		for _, val := range (col[i])[cx][rx].pVal {
			if val != 0 {
				b3 = col[i]
				b3[cx][rx].solved = true
				b3[cx][rx].value = val
				b3, _ = b3.possibleValues()
				col = append(col, b3)
			}
		}
		i++
	}
}

func (b board) validate() int {
	var fs int

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

func unsolved(b board) (int, int) {

	for c := 0; c < 9; c++ {
		for r := 0; r < 9; r++ {
			if b[c][r].solved == false {
				return c, r
			}
		}
	}

	return -1, -1
}
