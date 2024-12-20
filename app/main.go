package main

import (
	"fmt"
	"hypercubeapp/network"
	"math"
	"os"
	"strconv"
	"time"
)

func doHypercube() {

	dimension := -1

	fmt.Print("Enter hypercube dimension: ")
	fmt.Scan(&dimension)

	for dimension < 0 {

		fmt.Print("Dimension must be non-negative. Try again. Enter hypercube dimension: ")
		fmt.Scan(&dimension)

	}

	fmt.Println()

	fmt.Print("Building " + strconv.Itoa(dimension) + "-D hypercube ... ")

	start := time.Now()

	hCube := network.CreateHypercube(dimension)

	elapsed := time.Since(start)

	fmt.Println("It took", elapsed, "to build.")

	fmt.Println()

	nodeToTouch := 0

	for nodeToTouch > -1 {

		nodeToTouch = -2

		fmt.Println("What node would you like to touch?")

		for nodeToTouch < -1 || nodeToTouch > (int)(math.Pow(2, float64(dimension))-1) {

			fmt.Print("Enter a node number (0, ", (int)(math.Pow(2, float64(dimension))-1), ") or -1 to quit: ")
			fmt.Scan(&nodeToTouch)

		}

		if nodeToTouch > -1 {

			fmt.Println("Touching node", nodeToTouch, "...")

			hCube.Touch(nodeToTouch)

		}
	}

	fmt.Println()

}

func doGrid() {

	rows := -1
	cols := -1

	fmt.Print("Enter the number of rows and cols (e.g. 3 2): ")
	fmt.Scan(&rows)
	fmt.Scan(&cols)

	for rows < 1 || cols < 1 {

		fmt.Println("Dimensions must be positive.")
		fmt.Print("Try again. Enter the number of rows and cols (e.g. 3 2): ")
		fmt.Scan(&rows)
		fmt.Scan(&cols)

	}

	fmt.Println()

	fmt.Print("Building " + strconv.Itoa(rows) + "x" + strconv.Itoa(cols) + " grid ... ")

	start := time.Now()

	grid := network.CreateGrid(rows, cols)

	elapsed := time.Since(start)

	fmt.Println("It took", elapsed, "to build.")

	fmt.Println()

	rowToTouch := 1

	for rowToTouch > -1 {

		rowToTouch = -2
		colToTouch := -2

		fmt.Println("What node would you like to touch? Rows 0 -", rows-1, ". Cols 0 -", cols-1)

		for (rowToTouch < -1 || rowToTouch > rows-1) || (colToTouch < -1 || colToTouch > cols-1) {

			fmt.Print("Enter a node number x y (enter -1 -1 to quit): ")
			fmt.Scan(&rowToTouch)
			fmt.Scan(&colToTouch)

			if (rowToTouch < -1 || rowToTouch > rows-1) || (colToTouch < -1 || colToTouch > cols-1) {

				fmt.Println("Rows 0 -", rows-1, ":: Cols 0 -", cols-1)

			}

		}

		if rowToTouch == -1 {

			break

		} else if rowToTouch >= 0 {

			fmt.Println("Touching node (" + strconv.Itoa(rowToTouch) + ", " + strconv.Itoa(colToTouch) + ") ...")

			grid.Touch(rowToTouch, colToTouch)

		}
	}

	fmt.Println()

}

func doRing() {

	nodeCount := 0

	fmt.Print("Enter the number nodes in the ring (> 0): ")
	fmt.Scan(&nodeCount)

	for nodeCount < 1 {

		fmt.Println("Number of nodes must be positive.")
		fmt.Print("Try again. Enter the number nodes in the ring (> 0): ")
		fmt.Scan(&nodeCount)

	}

	fmt.Println()

	fmt.Print("Building ring of length ", nodeCount, " ...")

	start := time.Now()

	ring := network.CreateRing(nodeCount)

	elapsed := time.Since(start)

	fmt.Println("It took", elapsed, "to build.")

	fmt.Println()

	keepGoing := true

	for keepGoing {

		option := -1

		fmt.Println()

		for option < 0 || option > 3 {

			fmt.Println("What would you like to do with the ring?")
			fmt.Println("0. Quit (return to main menu)")
			fmt.Println("1. Execute Loop")
			fmt.Println("2. Execute Election")
			fmt.Println("----------------------------------------")
			fmt.Print("Enter option: ")
			fmt.Scan(&option)

			if option < 0 || option > 3 {

				fmt.Println()
				fmt.Println("Invalid option:", strconv.Itoa(option)+". Try again.")
				fmt.Println()

			}

		}

		switch option {
		case 0:

			keepGoing = false

		case 1:

			nodeToTouch := -1

			fmt.Print("Enter would you like to touch (0 - " + strconv.Itoa(nodeCount-1) + ") or -1 to exit: ")
			fmt.Scan(&nodeToTouch)

			for nodeToTouch < -1 {

				fmt.Print("Try again. Enter would you like to touch (0 - " + strconv.Itoa(nodeCount-1) + ") or -1 to exit: ")
				fmt.Scan(&nodeToTouch)

			}

			if nodeToTouch == -1 {

				break

			} else {

				direction := "undefined"

				fmt.Print("Enter the direction to loop (\"left\" or \"right\"): ")
				fmt.Scan(&direction)

				for direction != "left" && direction != "right" {

					fmt.Print("Try again. Enter the direction to loop (\"left\" or \"right\"): ")
					fmt.Scan(&direction)

				}

				fmt.Println("Touching node " + strconv.Itoa(nodeToTouch) + " ...")

				ring.Loop(nodeToTouch, direction)

			}

		case 2:

			ring.RunElection()

		}

	}

	fmt.Println()

}

func doQuit() {

	fmt.Println()
	fmt.Println("Goodbye")
	fmt.Println()

	os.Exit(0)

}

func main() {

	var actionMap map[int]func() = map[int]func(){
		0: doQuit,
		1: doHypercube,
		2: doGrid,
		3: doRing,
	}

	var option int

	for {

		option = -1

		fmt.Println("Choose a configuration to build.")
		fmt.Println("0. Quit")
		fmt.Println("1. Hypercube")
		fmt.Println("2. Grid")
		fmt.Println("3. Ring")
		fmt.Println("--------------------------------")
		fmt.Print("Enter option: ")
		fmt.Scan(&option)

		action, ok := actionMap[option]

		if ok {
			action()
		} else {
			fmt.Println()
			fmt.Println("Invalid option:", strconv.Itoa(option)+". Try again.")
			fmt.Println()

		}

	}

}
