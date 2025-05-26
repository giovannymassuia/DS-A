
export class Sudoku {
  eventHandlers = {};

  register(event, callback) {
    if (!this.eventHandlers[event]) {
      this.eventHandlers[event] = [];
    }
    this.eventHandlers[event].push(callback);
  }

  async notify(event, data) {
    if (this.eventHandlers[event]) {
      for (const callback of this.eventHandlers[event]) {
        await callback(data);
      }
    }
  }

  async validateBoard(board) {
    for (let row = 0; row < 9; row++) {
      for (let col = 0; col < 9; col++) {
        const num = board[row][col];
        if (num !== 0 && !(await this.isValid(board, row, col, num))) {
          await this.notify("invalidCell", { row, col, value: num, fixed: true });
          return false;
        }
      }
    }
    return true;
  }

  async isValid(board, row, col, num) {
    // check row
    for (let i = 0; i < 9; i++) {
      if(board[row][i] === 0 || i === col) continue;
      await this.notify("checkCell", { row, col: i, value: num });
      if (board[row][i] === num) {
        await this.notify("invalidCell", { row, col: i, value: num });
        return false;
      }
    }

    // check column
    for (let i = 0; i < 9; i++) {
      if(board[i][col] === 0 || i === row) continue;
      await this.notify("checkCell", { row, col: i, value: num });
      if (board[i][col] === num) {
        await this.notify("invalidCell", { row: i, col, value: num });
        return false;
      }
    }

    // check 3x3 grid
    const startRow = Math.floor(row / 3) * 3;
    const startCol = Math.floor(col / 3) * 3;
    for (let i = startRow; i < startRow + 3; i++) {
      for (let j = startCol; j < startCol + 3; j++) {
        if (board[i][j] === 0 || (i === row && j === col)) continue;
        await this.notify("checkCell", { row, col: i, value: num });
        if (board[i][j] === num) {
          await this.notify("invalidCell", { row: i, col: j, value: num });
          return false;
        }
      }
    }

    return true;
  }

  async solve(boardParam, notify = true) {
    // clone the board to avoid mutating the original
    const board = JSON.parse(JSON.stringify(boardParam));

    for (let row = 0; row < 9; row++) {
      for (let col = 0; col < 9; col++) {
        if (board[row][col] === 0) {
          for (let num = 1; num <= 9; num++) {
            if(notify) await this.notify('fillCell', { row, col, value: num });
            if (await this.isValid(board, row, col, num)) {
              board[row][col] = num;
              if(notify) await this.notify('fillCell', { row, col, value: num, fixed: true });
              if (await this.solve(board, notify)) {
                await this.notify('validCell', { row, col, value: num });
                return true;
              }
              board[row][col] = 0;
              if(notify) await this.notify('fillCell', { row, col, value: 0 });
            }
            if(notify) await this.notify('fillCell', { row, col, value: 0 });
          }
          return false;
        }
      }
    }
    return true;
  }

// Helper function to fill the board using backtracking
  async _fillAndSolveBoard(board) {
    for (let r = 0; r < 9; r++) {
      for (let c = 0; c < 9; c++) {
        if (board[r][c] === 0) { // Find empty cell
          const numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9];
          // Shuffle numbers for randomness
          for (let i = numbers.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            [numbers[i], numbers[j]] = [numbers[j], numbers[i]];
          }

          for (let num of numbers) {
            await this.notify("fillCell", { row: r, col: c, value: num });
            if (await this.isValid(board, r, c, num)) {
              board[r][c] = num;
              if (await this._fillAndSolveBoard(board)) {
                return true; // Successfully filled from this point
              }

              // Backtrack
              board[r][c] = 0;
              await this.notify("fillCell", { row: r, col: c+1, value: 0 });
            }
          }
          return false; // No number worked for this cell, trigger backtracking
        }
      }
    }
    return true; // All cells are already filled
  }

  async generateSudokuBoard() {
    const solvedBoard = Array.from({ length: 9 }, () => Array(9).fill(0));
    await this._fillAndSolveBoard(solvedBoard);
    const initialBoard = JSON.parse(JSON.stringify(solvedBoard));

    const removeNumbers = async (count) => {
      let removed = 0;
      while (removed < count) {
        const row = Math.floor(Math.random() * 9);
        const col = Math.floor(Math.random() * 9);
        if (initialBoard[row][col] !== 0) {
          const temp = initialBoard[row][col];
          initialBoard[row][col] = 0; // Remove the number
          await this.notify("fillCell", { row, col, value: 0 });
          // Check if the board is still solvable
          // if (!await this.solve(initialBoard, false)) {
          //   initialBoard[row][col] = temp; // Restore the number if not solvable
          //   await this.notify("fillCell", { row, col, value: temp });
          // } else {
            removed++;
          // }
        }
      }
    }

    await removeNumbers(30);

    return {
      solvedBoard,
      initialBoard,
    }
  }
}
