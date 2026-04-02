package main

import (
	"app/account"
	"app/files"
	"fmt"
)



func main()  {
	createAcc()
}


func createAcc() {
	acc, err := account.CreateAcc()
	if err != nil {
		fmt.Println(err)
		return 
	}

	file, err := acc.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
	}

	files.WriteFile(file, "data.json")

}