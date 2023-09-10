package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	x := 0
	var array []int
	for scanner.Scan() {
		currentY, _ := strconv.Atoi(scanner.Text())
		x++
		if currentY < 50 {
			currentY = 100
		} else if currentY > 350 {
			currentY = 200
		}
		array = append(array, currentY)
		y := math.Round(ava(array))
		standard := math.Round(StandardDeviation(array))
		fmt.Printf("%d %d\n", int((y - standard)), int((y + standard)))
	}
}

func ava(i []int) float64 {
	z := 0.0
	for x := 0; x < len(i); x++ {
		z += float64(i[x])
	}
	return (z / float64(len(i)))
}

func Variance(i []int) float64 {
	DeltaX := ava(i)
	TheX := 0.0
	Variance := 0.0
	for x := 0; x < len(i); x++ {
		z := (float64(i[x]) - float64(DeltaX))
		TheX += (z * z)
	}
	Variance = TheX / (float64(len(i)))
	return Variance
}

func StandardDeviation(i []int) float64 {
	x := Variance(i)
	return (math.Sqrt(x))
}
