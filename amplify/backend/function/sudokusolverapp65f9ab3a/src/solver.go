package main

import "fmt"

type cell struct {
	cellValue int
	rowNum    int
	colNum    int
}

type constraint struct {
	cellValue int
	rowNum    int
	colNum    int
	boxNum    int
}

func makeRowConstraint(i int, j int) constraint {
	return constraint{
		cellValue: j,
		rowNum:    i,
	}
}
func makeColConstraint(i int, j int) constraint {
	return constraint{
		cellValue: j,
		colNum:    i,
	}
}
func makeBoxConstraint(i int, j int) constraint {
	return constraint{
		cellValue: j,
		boxNum:    i,
	}
}
func makeCellConstraint(i int, j int) constraint {
	return constraint{
		rowNum: i,
		colNum: j,
	}
}

type solveState struct {
	cellToConstraints map[cell]map[constraint]bool
	constraintToCells map[constraint]map[cell]bool
	board             [9][9]int
}

func initialiseState(board [9][9]int) solveState {
	state := solveState{
		board:             board,
		cellToConstraints: make(map[cell]map[constraint]bool),
		constraintToCells: make(map[constraint]map[cell]bool),
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			for cellValue := 1; cellValue < 10; cellValue++ {
				box := (((i-1)/3)*3 + ((j - 1) / 3)) + 1

				rowConstraint := makeRowConstraint(i, cellValue)
				colConstraint := makeColConstraint(j, cellValue)
				boxConstraint := makeBoxConstraint(box, cellValue)
				cellConstraint := makeCellConstraint(i, j)

				c := cell{
					cellValue: cellValue,
					rowNum:    i,
					colNum:    j,
				}

				state.cellToConstraints[c] = map[constraint]bool{rowConstraint: true, colConstraint: true, boxConstraint: true, cellConstraint: true}
			}
		}
	}

	for c, constraints := range state.cellToConstraints {
		for constraint := range constraints {
			_, exists := state.constraintToCells[constraint]
			if !exists {
				state.constraintToCells[constraint] = make(map[cell]bool)
			}

			state.constraintToCells[constraint][c] = true
		}
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != 0 {
				state.fillCell(cell{
					cellValue: board[row][col],
					rowNum:    row + 1,
					colNum:    col + 1,
				})
			}
		}
	}

	return state
}

func (s solveState) fillCell(c cell) solveState {
	s.board[c.rowNum-1][c.colNum-1] = c.cellValue

	for filled_constraint := range s.cellToConstraints[c] {
		for cell_which_fills_constraint := range s.constraintToCells[filled_constraint] {
			for other_constraint_filled_by_cell := range s.cellToConstraints[cell_which_fills_constraint] {
				if filled_constraint != other_constraint_filled_by_cell {
					delete(s.constraintToCells[other_constraint_filled_by_cell], cell_which_fills_constraint)
				}
			}
			delete(s.cellToConstraints, cell_which_fills_constraint)
		}
		delete(s.constraintToCells, filled_constraint)
	}

	return s
}

func (s solveState) getFewestOptions() constraint {
	count := 10
	var current constraint
	for cons, cells := range s.constraintToCells {
		if len(cells) < count {
			current = cons
			count = len(cells)
		}
	}
	return current
}

func solve(state solveState) (solveState, bool) {
	if len(state.constraintToCells) == 0 {
		return state, true
	}

	constraintWithFewestCells := state.getFewestOptions()

	for cellOption := range state.constraintToCells[constraintWithFewestCells] {
		newState := state.deepcopy()
		newState = newState.fillCell(cellOption)
		newState, solved := solve(newState)
		if solved {
			return newState, true
		}

	}

	return state, false
}

func isBoardValid(board [9][9]int) bool {
	var rows [9][9]bool
	var cols [9][9]bool
	var boxes [9][9]bool

	for row_idx := 0; row_idx < 9; row_idx++ {
		for col_idx := 0; col_idx < 9; col_idx++ {
			box_idx := (row_idx/3)*3 + (col_idx / 3)

			value := board[row_idx][col_idx] - 1
			if value == -1 {
				continue
			}

			if rows[row_idx][value] || cols[col_idx][value] || boxes[box_idx][value] {
				return false
			}

			rows[row_idx][value], cols[col_idx][value], boxes[box_idx][value] = true, true, true
		}
	}

	return true
}

func SolveBoard(board [9][9]int) [9][9]int {
	if !isBoardValid(board) {
		return board
	}

	state := initialiseState(board)
	state, solved := solve(state)
	fmt.Println(solved)

	if solved {
		return state.board
	} else {
		return board
	}
}
