// Copyright © 2019 Mephis Pheies <mephistommm@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package main

import "fmt"

// https://leetcode.com/problems/sudoku-solver/
// https://leetcode.com/problems/sudoku-solver/discuss/15796/Singapore-prime-minister-Lee-Hsien-Loong's-Sudoku-Solver-code-runs-in-1ms

const (
	ROW_CELL   = 9
	ROW_SUBBOX = 3
	TOTAL_CELL = ROW_CELL * ROW_CELL
	BLANK      = 0
	ONES       = 0x3fe // Binary 1111111110
)

var (
	inBlock = [TOTAL_CELL]int{}
	inCol   = [TOTAL_CELL]int{}
	inRow   = [TOTAL_CELL]int{}
)

func init() {
	for i := 0; i < ROW_CELL; i++ {
		for j := 0; j < ROW_CELL; j++ {
			square := ROW_CELL*i + j
			inRow[square] = i
			inCol[square] = j
			inBlock[square] = (i/ROW_SUBBOX)*ROW_SUBBOX + (j / ROW_SUBBOX)
		}
	}
}

type solver struct {

	// Each int is a 9-bit array
	block [ROW_CELL]int
	col   [ROW_CELL]int
	row   [ROW_CELL]int

	// Records entries 1-9 in the grid, as the corresponding bit set to 1
	entry    [TOTAL_CELL]int
	sequence [TOTAL_CELL]int

	seqPtr int
}

func (slr *solver) swapSeqEntries(s1, s2 int) {
	temp := slr.sequence[s1]
	slr.sequence[s1] = slr.sequence[s2]
	slr.sequence[s2] = temp
}

func (slr *solver) initEntry(i, j int, v byte) {
	square := ROW_CELL*i + j
	valbit := 1 << uint(v-'0')

	slr.entry[square] = valbit
	slr.block[inBlock[square]] &= ^valbit
	slr.col[inCol[square]] &= ^valbit
	slr.row[inRow[square]] &= ^valbit

	seqPtr2 := slr.seqPtr
	for seqPtr2 < TOTAL_CELL && slr.sequence[seqPtr2] != square {
		seqPtr2++
	}

	slr.swapSeqEntries(slr.seqPtr, seqPtr2)
	slr.seqPtr++
}

func (slr *solver) nextSeq(s int) int {
	s2, minBitCount := TOTAL_CELL, 100
	for t := s; t < TOTAL_CELL; t++ {
		square := slr.sequence[t]
		possibles := slr.block[inBlock[square]] & slr.row[inRow[square]] & slr.col[inCol[square]]
		bitCount := 0
		for possibles > 0 {
			possibles &= ^(possibles & -possibles)
			bitCount++
		}

		if bitCount < minBitCount {
			minBitCount = bitCount
			s2 = t
		}
	}

	return s2
}

func (slr *solver) place(s int) bool {
	if s >= TOTAL_CELL {
		return true
	}

	s2 := slr.nextSeq(s)
	slr.swapSeqEntries(s, s2)

	square := slr.sequence[s]
	ib, ir, ic := inBlock[square], inRow[square], inCol[square]
	possibles := slr.block[ib] & slr.row[ir] & slr.col[ic]
	for possibles > 0 {
		valbit := possibles & (-possibles)
		possibles &= ^valbit
		slr.entry[square] = valbit
		slr.block[ib] &= ^valbit
		slr.row[ir] &= ^valbit
		slr.col[ic] &= ^valbit

		if slr.place(s + 1) {
			return true
		}

		slr.entry[square] = BLANK
		slr.block[ib] |= valbit
		slr.row[ir] |= valbit
		slr.col[ic] |= valbit
	}

	slr.swapSeqEntries(s, s2)
	return false
}

func (slr *solver) solve(board [][]byte) {
	// clear entry
	for square := 0; square < TOTAL_CELL; square++ {
		slr.sequence[square] = square
		slr.entry[square] = BLANK
	}

	for i := 0; i < ROW_CELL; i++ {
		slr.block[i] = ONES
		slr.row[i] = ONES
		slr.col[i] = ONES
	}

	slr.seqPtr = 0

	// init entry
	for i := 0; i < ROW_CELL; i++ {
		for j := 0; j < ROW_CELL; j++ {
			if board[i][j] != '.' {
				slr.initEntry(i, j, board[i][j])
			}
		}
	}

	slr.place(slr.seqPtr)
	slr.fillBoard(board)
}

func (slr *solver) fillBoard(board [][]byte) {
	for i := 0; i < ROW_CELL; i++ {
		for j := 0; j < ROW_CELL; j++ {
			if board[i][j] == '.' {
				valbit := slr.entry[i*ROW_CELL+j]
				for val := uint(1); val < 10; val++ {
					if valbit == 1<<val {
						board[i][j] = '0' + byte(val)
					}
				}
			}
		}
	}
}

var defaultSolver = &solver{}

// 0 ms	2 MB
func solveSudoku(board [][]byte) {
	defaultSolver.solve(board)
}

func showBoard(board [][]byte) {
	for i := 0; i < ROW_CELL; i++ {
		for j := 0; j < ROW_CELL; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println("")
	}
}

func main() {
	b1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	showBoard(b1)
	solveSudoku(b1)
	showBoard(b1)

}
