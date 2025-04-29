package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100) / 100
	fmt.Println("new sum:", sum)

	fmt.Println("pi:", math.Pi)

	circleRad :=15.5
	circleCircum := circleRad* 2 * math.Pi  
	fmt.Printf("circumference: %.2f \n", circleCircum)

}
