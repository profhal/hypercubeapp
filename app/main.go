package main

import (
	"fmt"
	"hypercubeapp/hypercube"
	"math"
	"strconv"
	"time"
)

func main() {

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

	hCube := hypercube.CreateHypercube(dimension)

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
