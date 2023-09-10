package main
import (
	"bufio"
	"fmt"
	// "math"
	"os"
	// "sort"
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
		} else if currentY > 600 {
			currentY = 200
		}
		array = append(array, currentY)
		firstY := array[0]
		y := (((currentY - firstY) / x) * x) + firstY
		Vari := Variance(array,x)
		Yrange := (int(Vari) / 10)-30
		fmt.Printf("%d %d\n", (y -3 - (Yrange/50)), (y +3 + (Yrange/50)))
	}
}
func ava(i []int) float64 {
	z := 0.0
	for x := 0; x < len(i); x++ {
		z += float64(i[x])
	}
	return (z / float64(len(i)))
}
func Variance(i []int,n int) float64 {
	DeltaX := ava(i)
	TheX := 0.0
	Variance := 0.0
	if n < 10 {
	for x := 0; x < len(i); x++ {
		z := (float64(i[x]) - float64(DeltaX))
		TheX += (z * z)
	}
	Variance = TheX / (float64(len(i)))
	} else {
		for x := 0 ; x < 10 ; x++ {
			z := (float64(i[x]) - float64(DeltaX))
			TheX += (z * z)
		}
		Variance = TheX / (float64(10))

	}
	return Variance
}



func printError(err error) {
	fmt.Println("ERROR: " + err.Error())
	os.Exit(1)
}
