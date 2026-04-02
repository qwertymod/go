package main

import (
	"fmt"
)



func main() {
	acc, err := createAcc()
	if err != nil {
		fmt.Println(err)
		return 
	}
	printAcc(acc)
}