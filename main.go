package main

import (
	"app/account"
	"app/files"
	"fmt"
)



func main()  {
	//var l account.Account// - могу создать структуру из другого пакета благодаря заглавной букве 
	acc, err := account.CreateAcc()
	if err != nil {
		fmt.Println(err)
		return 
	}
	files.WriteFile()

	acc.PrintAcc()
}