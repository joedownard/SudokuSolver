import numpy as np
import copy


class Solver:
    def __deepcopy__(self, memodict):
        cls = self.__class__
        copied = cls.__new__(cls)
        copied.X = {key: [value for value in self.X[key]] for key in self.X}
        copied.Y = {key: [value for value in self.Y[key]] for key in self.Y}
        return copied

    def __init__(self, sudoku):
        self.X = []
        self.Y = {}
        self.build_dictionaries()
        self.convert_to_dictionary_and_fill()
        self.constrain_for_input(sudoku)

    def build_dictionaries(self):
        col_counter = 0
        for i in range(729):
            cell_val = (i % 9) + 1
            row_num = i // 81
            col_num = col_counter
            box_num = (i // 81 // 3) * 3 + (col_counter // 3)

            cell_constraint = ("cell_constraint", (row_num, col_num))
            row_constraint = ("row_constraint", (row_num, cell_val))
            col_constraint = ("col_constraint", (col_num, cell_val))
            box_constraint = ("box_constraint", (box_num, cell_val))

            self.Y[(i // 81, col_counter, cell_val)] = (cell_constraint, row_constraint, col_constraint, box_constraint)

            if col_counter <= 7:
                if i % 9 == 8:
                    col_counter += 1
            elif col_counter == 8:
                if i % 9 == 8:
                    col_counter = 0

        for i in range(9):
            for j in range(9):
                self.X.append(("cell_constraint", (i, j)))
        for i in range(9):
            for j in range(9):
                self.X.append(("row_constraint", (i, j + 1)))
        for i in range(9):
            for j in range(9):
                self.X.append(("col_constraint", (i, j + 1)))
        for i in range(9):
            for j in range(9):
                self.X.append(("box_constraint", (i, j + 1)))

    def convert_to_dictionary_and_fill(self):
        self.X = {constraint: set() for constraint in self.X}
        for cell_val, constraints in self.Y.items():
            for constraint in constraints:
                self.X[constraint].add(cell_val)

    def constrain_for_input(self, sudoku):
        for (row_idx, col_idx), value in np.ndenumerate(sudoku):
            if value != 0:
                self.fill_cell((row_idx, col_idx, value))

    def fill_cell(self, col):
        for constraint1 in self.Y[col]:
            for cell_that_satisfies_constraint in self.X[constraint1]:
                for constraint2 in self.Y[cell_that_satisfies_constraint]:
                    if constraint1 != constraint2:
                        self.X[constraint2].remove(cell_that_satisfies_constraint)
                self.Y.pop(cell_that_satisfies_constraint)
            self.X.pop(constraint1)


def is_board_valid(sudoku):
    rows = [[] for _ in range(9)]
    cols = [[] for _ in range(9)]
    boxes = [[] for _ in range(9)]
    for (row_idx, col_idx), value in np.ndenumerate(sudoku):
        if value != 0:
            box_index = (row_idx // 3) * 3 + (col_idx // 3)
            if value in rows[row_idx]:
                return False
            if value in cols[col_idx]:
                return False
            if value in boxes[box_index]:
                return False

            rows[row_idx].append(value)
            cols[col_idx].append(value)
            boxes[box_index].append(value)
    return True


def solve(state, partial_solution):
    least_options = min(state.X, key=lambda constraint: len(state.X[constraint]))

    while len(state.X[least_options]) == 1:
        partial_solution.append(next(iter(state.X[least_options])))
        state.fill_cell(next(iter(state.X[least_options])))
        if len(state.X) == 0:
            return partial_solution
        least_options = min(state.X, key=lambda constraint: len(state.X[constraint]))

    for option in state.X[least_options]:
        new_state = copy.deepcopy(state)
        new_state.fill_cell(option)
        partial_solution.append(option)

        deep_state = solve(new_state, partial_solution)
        if deep_state is not None:
            return deep_state
    return None


def sudoku_solver(sudoku):
    if not is_board_valid(sudoku):
        return np.full_like(sudoku, -1)

    solver = Solver(sudoku)
    solution = solve(solver, [])
    if solution:
        for (row, col, value) in solution:
            sudoku[row][col] = value
    else:
        return np.full_like(sudoku, -1)

    return sudoku
