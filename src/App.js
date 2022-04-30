import './App.css';
import { useEffect, useState } from 'react'
import Amplify, { API } from 'aws-amplify'
import awsExports from "./aws-exports";
Amplify.configure(awsExports);

let classNames = require('classnames');

function App() {
  const [board, setBoard] = useState([[]])

  useEffect(() => {
    let defaultBoard = [[], [], [], [], [], [], [], [], []];
    for (let i = 0; i < 9; i++) {
      for (let j = 0; j < 9; j++) {
        defaultBoard[i][j] = "";
      }
    }
    setBoard(defaultBoard);
  }, [])

  const onCellChange = (i, j, val) => {
    if (val.length <= 1) {
      const newBoard = [...board];
      newBoard[i][j] = val;
      setBoard(newBoard);
    }
  }

  const onSolve = async () => {
    const request = {
      body: {
        "board": board
      }
    };

    const result = await API.post("sudokusolverapi", "/solveSudoku", request);
  }

  return (
    <div className="central">
      <div className="board">
        {board.map((row, i) => {
          return row.map((cell, j) => {
            var cellClass = classNames({
              'cell': true,
              'cell-bottom-border': i === 2 || i === 5,
              'cell-right-border': j === 2 || j === 5
            });

            return <div key={(i, j)} className={cellClass} >
              <input className='number' value={board[i][j]} onChange={(evt) => onCellChange(i, j, evt.target.value)}>
              </input>
            </div>
          })
        })}
      </div>

      <button className="solveButton" onClick={onSolve}>
        Solve
      </button>
    </div>
  );
}

export default App;
