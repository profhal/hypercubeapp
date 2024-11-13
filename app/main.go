package main

import (
	"fmt"
	"hypercubeapp/hypercube"
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

	fmt.Println("Kicking off hypercube...")

	hCube.Run()

	fmt.Println()

	fmt.Print("::Press enter to shut down the hypercube::")
	fmt.Scanln()

}
