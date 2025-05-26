import {deepEqual} from "node:assert";
import test from "node:test";

/*
  Problem: Write a program to solve a Sudoku puzzle by filling the empty cells.
  Example:
   - Input: board = [
      ["5","3",".",".","7",".",".",".","."],
      ["6",".",".","1","9","5",".",".","."],
      [".","9","8",".",".",".",".","6","."],
      ["8",".",".",".","6",".",".",".","3"],
      ["4",".",".","8",".","3",".",".","1"],
      ["7",".",".",".","2",".",".",".","6"],
      [".","6",".",".",".",".","2","8","."],
      [".",".",".","4","1","9",".",".","5"],
      [".",".",".",".","8",".",".","7","9"]
    ]
*  Output: [
      ["5","3","4","6","7","8","9","1","2"],
      ["6","7","2","1","9","5","3","4","8"],
      ["1","9","8","3","4","2","5","6","7"],
      ["8","5","9","7","6","1","4","2","3"],
      ["4","2","6","8","5","3","7","9","1"],
      ["7","1","3","9","2","4","8","5","6"],
      ["9","6","1","5","3","7","2","8","4"],
      ["2","8","7","4","1","9","6","3","5"],
      ["3","4","5","2","8","6","1","7","9"]
    ]
 */


function isValid(board: string[][], row: number, col: number, num: string): boolean {
  // Check if the number is already in the row or column
  for (let i = 0; i < board.length; i++) {
    if (board[row][i] === num || board[i][col] === num) {
      return false; // Check row and column
    }
  }

  // Check if the board has grids and determine the grid size
  // 3x3 board = 1 grid
  // 4x4 board = 2 grids
  // 9x9 board = 3 grids
  const gridSize = board.length === 4 ? 2 : board.length === 9 ? 3 : 1;

  const startRow = Math.floor(row / gridSize) * gridSize;
  const startCol = Math.floor(col / gridSize) * gridSize;
  for (let i = startRow; i < startRow + gridSize; i++) {
    for (let j = startCol; j < startCol + gridSize; j++) {
      if (board[i][j] === num) {
        return false;
      }
    }
  }
  return true;
}

function solveSudoku(board: string[][], numbers = 9): boolean {
  const boardSize = board.length;
  for (let row = 0; row < boardSize; row++) {
    for (let col = 0; col < boardSize; col++) {
      if (board[row][col] === ".") {
        for (let num = 1; num <= numbers; num++) {
          const strNum = num.toString();
          if (isValid(board, row, col, strNum)) {
            board[row][col] = strNum; // Place the number
            if (solveSudoku(board, numbers)) {
              return true; // If solved, return true
            }
            board[row][col] = "."; // Backtrack
          }
        }
        return false; // No valid number found, backtrack
      }
    }
  }
  return true; // Solved
}

test("sudoku-solver", () => {
  const board = [
    ["5", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
  ];

  solveSudoku(board);

  // pretty print the board, grid by grid with spaces around them
  for (let i = 0; i < board.length; i++) {
    for (let j = 0; j < board[i].length; j++) {
      process.stdout.write(board[i][j] + " ");
      if ((j + 1) % 3 === 0 && j < board[i].length - 1) {
        process.stdout.write("  "); // Add space between grids
      }
    }
    console.log(); // New line after each row
    if ((i + 1) % 3 === 0 && i < board.length - 1) {
      console.log(); // Add a blank line between grids
    }
  }

  deepEqual(board, [
      ["5", "3", "4",   "6", "7", "8",   "9", "1", "2"],
      ["6", "7", "2",   "1", "9", "5",   "3", "4", "8"],
      ["1", "9", "8",   "3", "4", "2",   "5", "6", "7"],

      ["8", "5", "9",   "7", "6", "1",   "4", "2", "3"],
      ["4", "2", "6",   "8", "5", "3",   "7", "9", "1"],
      ["7", "1", "3",   "9", "2", "4",   "8", "5", "6"],

      ["9", "6", "1",   "5", "3", "7",   "2", "8", "4"],
      ["2", "8", "7",   "4", "1", "9",   "6", "3", "5"],
      ["3", "4", "5",   "2", "8", "6",   "1", "7", "9"],
    ]);
});

test("sudoku-solver 98% solved", () => {
  const board = [
    ["5", "3", ".", "6", "7", "8", ".", "1", "2"],
    ["6", "7", "2", "1", "9", "5", "3", "4", "8"],
    ["1", "9", "8", "3", "4", "2", "5", "6", "7"],
    [".", "5", "9", "7", "6", "1", "4", "2", "3"],
    ["4", "2", "6", "8", "5", "3", "7", "9", "1"],
    ["7", "1", "3", "9", "2", "4", "8", "5", "6"],
    ["9", "6", "1", "5", "3", "7", "2", "8", "4"],
    ["2", "8", "7", "4", "1", "9", "6", "3", "5"],
    ["3", "4", "5", "2", "8", "6", "1", "7", "9"],
  ];

  solveSudoku(board);

  console.log(board);

  deepEqual(board, [
    ["5", "3", "4", "6", "7", "8", "9", "1", "2"],
    ["6", "7", "2", "1", "9", "5", "3", "4", "8"],
    ["1", "9", "8", "3", "4", "2", "5", "6", "7"],
    ["8", "5", "9", "7", "6", "1", "4", "2", "3"],
    ["4", "2", "6", "8", "5", "3", "7", "9", "1"],
    ["7", "1", "3", "9", "2", "4", "8", "5", "6"],
    ["9", "6", "1", "5", "3", "7", "2", "8", "4"],
    ["2", "8", "7", "4", "1", "9", "6", "3", "5"],
    ["3", "4", "5", "2", "8", "6", "1", "7", "9"],
  ]);
});

// test with a board 4x4 and only 4 numbers
test("sudoku-solver 4x4 with 4 numbers", () => {
  const board = [
    ["1", ".", ".", "."],
    [".", "2", ".", "."],
    [".", ".", "3", "."],
    [".", ".", ".", "4"],
  ];

  solveSudoku(board, 4);

  console.log(board);

  deepEqual(board, [
      ["1", "3", "4", "2"],
      ["4", "2", "1", "3"],
      ["2", "4", "3", "1"],
      ["3", "1", "2", "4"],
    ]);
});

// test with a board 3x3 and only 3 numbers
test("sudoku-solver 3x3 with 3 numbers", () => {
  const board = [
    ["1", ".", "."],
    [".", "2", "."],
    [".", ".", "3"],
  ];


  solveSudoku(board, 3);

  console.log(board);

  deepEqual(board, [
      ["1", "3", "2"],
      ["3", "2", "1"],
      ["2", "1", "3"],
    ]);
});
