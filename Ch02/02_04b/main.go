package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Sum: ", intSum)


	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Sum: ", floatSum)

	total := float64(i1) + f2
	fmt.Println("total: ", total)

	product := float64(i1) * f2
	fmt.Println("product: ", product)

}
