package main

func deepcopyCellToConstraints(cellToConstraints map[cell]map[constraint]bool) map[cell]map[constraint]bool {
	newCellToConstraints := make(map[cell]map[constraint]bool)
	for cell, constraintsMap := range cellToConstraints {
		newCellToConstraints[cell] = make(map[constraint]bool)
		for constraint := range constraintsMap {
			newCellToConstraints[cell][constraint] = true
		}
	}

	return newCellToConstraints
}

func deepcopyConstraintsToCells(constraintsToCells map[constraint]map[cell]bool) map[constraint]map[cell]bool {
	newConstraintsToCell := make(map[constraint]map[cell]bool)
	for constraint, cellsMap := range constraintsToCells {
		newConstraintsToCell[constraint] = make(map[cell]bool)
		for cell := range cellsMap {
			newConstraintsToCell[constraint][cell] = true
		}
	}
	return newConstraintsToCell
}

func (s solveState) deepcopy() solveState {
	return solveState{
		cellToConstraints: deepcopyCellToConstraints(s.cellToConstraints),
		constraintToCells: deepcopyConstraintsToCells(s.constraintToCells),
		board:             s.board,
	}
}
