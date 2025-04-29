package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	anum := 42
	fmt.Println(str1, str2, str3)
	strLen, err := fmt.Println("The val is: ", anum)
	if err == nil {
		// this means no error
		fmt.Println("String len: ", strLen)
	}

	fmt.Printf("Val of num: %v\n", anum)
	fmt.Printf("Datatype: %T \n", anum)
}
