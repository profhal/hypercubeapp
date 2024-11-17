package network

import (
	"strconv"
)

type Grid struct {
	Master
	rowCount  int
	colCount  int
	nodeCount int
	nodes     [][]*gridNode
	inputQ    chan string
}

func CreateGrid(rowCount int, colCount int) *Grid {

	grid := new(Grid)

	// Prep the nodes
	//
	grid.rowCount = rowCount
	grid.colCount = colCount
	grid.nodeCount = grid.rowCount * grid.colCount
	grid.inputQ = make(chan string)

	grid.nodes = make([][]*gridNode, 0, grid.rowCount)

	for r := 0; r < grid.rowCount; r++ {

		grid.nodes = append(grid.nodes, make([]*gridNode, 0, grid.colCount))

		for c := 0; c < grid.colCount; c++ {

			grid.nodes[r] = append(grid.nodes[r], new(gridNode))

			// node[r][c] is in position (c, r) in the grid.
			grid.nodes[r][c].id = "(" + strconv.Itoa(c) + ", " + strconv.Itoa(r) + ")"
			grid.nodes[r][c].neighborCount = 0
			grid.nodes[r][c].inputQ = make(chan string, 4)

			grid.nodes[r][c].Start(grid, "0")

		}
	}

	// Wire the grid
	//
	for r := range grid.nodes {

		for c := range grid.nodes[r] {

			// Remember! node[r][c] is in grid position (c, r))
			//
			// Given the grid,
			//
			//                   ^
			//                   |
			//   rowCount - 1 >  o---o---o---o---o---o
			//                   |   |   |   |   |   |
			//                   o---o---o---o---o---o
			//                   |   |   |   |   |   |
			//                   o---o---o---o---o---o->
			//                                       ^ colCount - 1
			// Observe that
			//     - the origin (0, 0) has no down or left
			//     - the left side (0, x), x < rowCount-1, there is no left
			//     - the upper left corner (0, rowCount - 1) has no left or up
			//     - the lower right corner (colCount-1, 0) has no right or down
			//     - the right side (colCount-1, x), x < rowCount-1, has no right
			//     - the upper right corner (colCount - 1, rowCount - 1) has no up or right
			//     - the bottom wall (x, 0), 0 < x < colCount - 1, has no down
			//     - the top wall (x, rowCount-1), 0 < x < colCount - 1, has now up
			//     - all interior points have all four neighbors
			//
			if c == 0 {

				if r == 0 {

					// (0, 0) has no down or left
					//
					grid.nodes[r][c].up = grid.nodes[r+1][c] // up

					grid.nodes[r][c].right = grid.nodes[r][c+1] // right

				} else if r < grid.rowCount-1 {

					// 0 < x < rowCount-1 : (0, x) has no left
					//
					grid.nodes[r][c].up = grid.nodes[r+1][c] // up

					grid.nodes[r][c].right = grid.nodes[r][c+1] // right

					grid.nodes[r][c].down = grid.nodes[r-1][c] // down

				} else {

					// (0, rowCount-1) has no up or left
					//
					grid.nodes[r][c].right = grid.nodes[r][c+1] // right

					grid.nodes[r][c].down = grid.nodes[r-1][c] // down

				}

			} else if c == grid.colCount-1 {

				if r == 0 {

					// (colCount-1, 0) has no right or down
					//
					grid.nodes[r][c].left = grid.nodes[r][c-1] // left

					grid.nodes[r][c].up = grid.nodes[r+1][c] // up

				} else if r < grid.rowCount-1 {

					// x < rowCount : (colCount, x) has no right
					//
					grid.nodes[r][c].left = grid.nodes[r][c-1] // left

					grid.nodes[r][c].up = grid.nodes[r+1][c] // up

					grid.nodes[r][c].down = grid.nodes[r-1][c] // down

				} else {

					// (colCount-1, rowCount-1) has no up or right
					//
					grid.nodes[r][c].left = grid.nodes[r][c-1] // left

					grid.nodes[r][c].down = grid.nodes[r-1][c] // down

				}

			} else if r == 0 {

				// x < colCount : (0, x) has no down
				//
				// Note: c = 0 was handled above.
				//
				if c < colCount-1 {

					grid.nodes[r][c].left = grid.nodes[r][c-1] // left

					grid.nodes[r][c].up = grid.nodes[r+1][c] // up

					grid.nodes[r][c].right = grid.nodes[r][c+1] // right

				}

			} else if r == rowCount-1 {

				// already handled (0, rowCount) in the c=0 case. We avoid handling (colCount-1, rowCount-1)
				//
				if c < colCount-1 {

					// 0 < x < colCount : (x, rowCount-1) has no up
					//
					grid.nodes[r][c].left = grid.nodes[r][c-1] // left

					grid.nodes[r][c].right = grid.nodes[r][c+1] // right

					grid.nodes[r][c].down = grid.nodes[r-1][c] // down

				}

			} else {

				// We are somewhere in the interior of the grid. Full wiring.
				//
				grid.nodes[r][c].left = grid.nodes[r][c-1] // left

				grid.nodes[r][c].up = grid.nodes[r+1][c] // up

				grid.nodes[r][c].right = grid.nodes[r][c+1] // right

				grid.nodes[r][c].down = grid.nodes[r-1][c] // down

			}

		}

	}

	return grid

}

// Runs the grid task
func (g *Grid) Touch(row int, col int) {

	g.nodes[row][col].inputQ <- "start"

	<-g.inputQ

}

func (g *Grid) NodeFinished() {
	g.inputQ <- "done"
}
