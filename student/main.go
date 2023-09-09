package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	x := 0
	var array []int
	for scanner.Scan() {
		currentY, _ := strconv.Atoi(scanner.Text())
		x++
		if currentY < 20 {
			currentY = 100
		} else if currentY < 600 {
			currentY = 200
		}
		array = append(array, currentY)
		firstY := array[0]
		y := (((currentY - firstY) / x) * x) + firstY
		Vari := Variance(array,x)
		Yrange := (int(Vari) / 10)
		fmt.Printf("%d %d\n", (y - (Yrange )), (y + (Yrange )))

	}
}

func ava(i []int) float64 {
	z := 0.0
	for x := 0; x < len(i); x++ {
		z += float64(i[x])
	}
	return (z / float64(len(i)))
}

func Median(i []int) float64 {
	sort.Ints(i)
	Median := 0.0
	if len(i)%2 == 1 {
		Median = float64(i[(len(i) / 2)])
	} else {
		Median = (float64(i[(len(i)/2)]) + float64(i[(len(i)/2)-1])) / 2.0
	}
	return Median
}

func Variance(i []int,n int) float64 {
	DeltaX := ava(i)
	TheX := 0.0
	if n <= 10 {
	for x := 0; x < len(i); x++ {
		z := (float64(i[x]) - float64(DeltaX))
		TheX += (z * z)
	}} else {
		for x := n-11 ; x < n ; x++ {
			z := (float64(i[x]) - float64(DeltaX))
			TheX += (z * z)
		}
	}
	Variance := TheX / (float64(10))
	return Variance
}


func printError(err error) {
	fmt.Println("ERROR: " + err.Error())
	os.Exit(1)
}
