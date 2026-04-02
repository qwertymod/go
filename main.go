package main

import (
	"app/account"
	"app/files"
	"fmt"
)



func main()  {
	files.ReadFile()
	files.WriteFile("scidish", "file.txt")
	acc, err := account.CreateAcc()
	if err != nil {
		fmt.Println(err)
		return 
	}

	acc.PrintAcc()
}