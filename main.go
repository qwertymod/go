package main

import (
	"fmt"
	"app/account"
)



func main() {
	acc, err := account.CreateAcc()
	if err != nil {
		fmt.Println(err)
		return 
	}
	account.PrintAcc(acc)
}