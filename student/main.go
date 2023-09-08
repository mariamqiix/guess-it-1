package main

import (
	"bufio"
	"fmt"
	"math"
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
		array = append(array, currentY)
		firstY := array[0]
		y := (((currentY - firstY) / x) * x) + firstY
		Vari := Variance(array)
		Yrange := (int(Vari) / 10)
		fmt.Print((y - (Yrange / 5)))
		fmt.Print(" - ")
		fmt.Println(y + (Yrange / 5))
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

func Variance(i []int) float64 {
	DeltaX := ava(i)
	TheX := 0.0
	for x := 0; x < len(i); x++ {
		z := (float64(i[x]) - float64(DeltaX))
		TheX += (z * z)
	}
	Variance := TheX / (float64(len(i)))
	return Variance
}

func StandardDeviation(i []int) float64 {
	x := Variance(i)
	return (math.Sqrt(x))
}

func printError(err error) {
	fmt.Println("ERROR: " + err.Error())
	os.Exit(1)
}
