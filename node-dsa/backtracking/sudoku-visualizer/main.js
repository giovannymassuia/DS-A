import {
  createBoard,
  fillCell,
  highlightCell,
  unhighlightCell
} from "./board.js";
import {Sudoku} from "./sudoku.js";

const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));
let time = 100;
const speedSlider = document.getElementById("speed-slider");
const speedValue = document.getElementById("speed-value");
speedSlider.addEventListener("input", (e) => {
  const speed = e.target.value;
  speedValue.innerText = speed;
  time = parseInt(speed, 10);
});
speedValue.innerText = time;
speedSlider.value = time;

const sudoku = new Sudoku();

async function fillCellCb ({row, col, value, fixed = false}) {
  await sleep(time);
  fillCell(row, col, value);
  highlightCell(row, col, 'fill');
  await sleep(time);
  if(!fixed) unhighlightCell(row, col, 'fill');
}
sudoku.register("fillCell", fillCellCb);

sudoku.register("invalidCell", async ({row, col, value, fixed = false}) => {
  await sleep(time * 2);
  highlightCell(row, col, 'error');
  await sleep(time * 2);
  if(!fixed) unhighlightCell(row, col, 'error');
});

sudoku.register("checkCell", async ({row, col, value}) => {
  await sleep(time/8);
  highlightCell(row, col, 'check');
  await sleep(time/8);
  unhighlightCell(row, col, 'check');
});

let initialBoard = [];
let solvedBoard = [];
async function init() {
  initialBoard = [];
  solvedBoard = [];

  createBoard();

  const game = await sudoku.generateSudokuBoard();
  initialBoard = game.initialBoard;
  solvedBoard = game.solvedBoard;

  console.log("Initial Board:", initialBoard);
  console.log("Solved Board:", solvedBoard);
}

async function solve() {
  console.log("Solving Sudoku...");
  await sudoku.solve(initialBoard);

  // mess up the board to test validation
  // initialBoard[1][4] = 3; // Invalid move
  // await fillCellCb({ row: 1, col: 4, value: 3, fixed: true });

  console.log("Valid Solution:", await sudoku.validateBoard(initialBoard));
}

createBoard();
initialBoard = [
  [0, 0, 0, 1, 7, 0, 8, 0, 0],
  [8, 0, 0, 0, 0, 2, 7, 1, 9],
  [0, 4, 7, 0, 0, 8, 6, 2, 5],
  [3, 1, 4, 9, 0, 7, 0, 0, 0],
  [0, 8, 6, 0, 3, 0, 1, 9, 2],
  [2, 9, 5, 6, 0, 1, 4, 3, 0],
  [5, 3, 0, 7, 0, 0, 0, 8, 0],
  [6, 7, 0, 2, 4, 3, 9, 5, 1],
  [4, 2, 1, 8, 5, 0, 0, 7, 6]
]
for(let r = 0; r < 9; r++) {
  for(let c = 0; c < 9; c++) {
    fillCell(r, c, initialBoard[r][c]);
  }
}

document.getElementById("solve-btn").addEventListener("click", solve);
document.getElementById("generate-btn").addEventListener("click", init);


