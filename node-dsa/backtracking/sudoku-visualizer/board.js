export function createBoard() {
  const container = document.getElementById("board");
  container.innerHTML = "";

  for (let i = 0; i < 3; i++) {
    const gridRow = document.createElement("div");
    gridRow.className = "row";

    for (let j = 0; j < 3; j++) {
      const grid = document.createElement("div");
      grid.className = "grid";

      for (let row = 0; row < 3; row++) {
        const rowDiv = document.createElement("div");
        rowDiv.className = "row";

        for (let col = 0; col < 3; col++) {
          const cell = document.createElement("div");
          cell.className = "cell";
          cell.id = `cell-${i * 3 + row}-${j * 3 + col}`;
          rowDiv.appendChild(cell);
        }

        grid.appendChild(rowDiv);
      }

      gridRow.appendChild(grid);
    }

    container.appendChild(gridRow);
  }
}

export function highlightCell(row, col, type) {
  const cell = document.getElementById(`cell-${row}-${col}`);
  cell.classList.add("highlight-" + (type || "default"));
}

export function unhighlightCell(row, col, type) {
  const cell = document.getElementById(`cell-${row}-${col}`);
  cell.classList.remove("highlight-" + (type || "default"));
}

export function fillCell(row, col, value) {
  const cell = document.getElementById(`cell-${row}-${col}`);
  if(!cell) return;
  if (value === 0) {
    cell.innerText = "";
  } else {
    cell.innerText = value;
  }
}
