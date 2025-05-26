/**
 * Sudoku Generator
 * This function generates a random Sudoku puzzle by filling the empty cells
 * it uses backtracking to fill the cells with numbers from 1 to 9
 * difficulty level is determined by the number of empty cells
 * 30 empty cells = easy
 * 40 empty cells = medium
 * 50 empty cells = hard
 * @return { solvedBoard: string[][], initialBoard: string[][] }
 */
function generateSudoku(difficulty: number): { solvedBoard: string[][]; initialBoard: string[][] } {
  const board: string[][] = Array.from({ length: 9 }, () => Array(9).fill("."));
  const solvedBoard: string[][] = Array.from({ length: 9 }, () => Array(9).fill("."));

  function isValid(board: string[][], row: number, col: number, num: string): boolean {
    for (let i = 0; i < board.length; i++) {
      if (board[row][i] === num || board[i][col] === num) {
        return false;
      }
    }

    const gridSize = Math.sqrt(board.length);
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

  function solveSudoku(board: string[][]): boolean {
    for (let row = 0; row < board.length; row++) {
      for (let col = 0; col < board.length; col++) {
        if (board[row][col] === ".") {
          for (let num = 1; num <= 9; num++) {
            const strNum = String(num);
            if (isValid(board, row, col, strNum)) {
              board[row][col] = strNum;
              if (solveSudoku(board)) {
                return true;
              }
              board[row][col] = "."; // Backtrack
            }
          }
          return false; // No valid number found
        }
      }
    }
    return true; // Solved
  }

  function fillBoard(board: string[][]): void {
    for (let row = 0; row < board.length; row++) {
      for (let col = 0; col < board.length; col++) {
        if (board[row][col] === ".") {
          const num = Math.floor(Math.random() * 9) + 1;
          const strNum = String(num);
          if (isValid(board, row, col, strNum)) {
            board[row][col] = strNum;
            fillBoard(board);
            if (!solveSudoku(board)) {
              board[row][col] = "."; // Backtrack
            }
          }
        }
      }
    }
  }

  function removeNumbers(board: string[][], count: number): void {
    let removed = 0;
    while (removed < count) {
      const row = Math.floor(Math.random() * 9);
      const col = Math.floor(Math.random() * 9);
      if (board[row][col] !== ".") {
        board[row][col] = ".";
        removed++;
      }
    }
  }

  fillBoard(solvedBoard);
  removeNumbers(solvedBoard, difficulty);
  for (let row = 0; row < board.length; row++) {
    for (let col = 0; col < board.length; col++) {
      board[row][col] = solvedBoard[row][col];
    }
  }
  removeNumbers(board, difficulty);
  return { solvedBoard, initialBoard: board };
}
