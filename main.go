package main

import (
	"app/account"
	"app/files"
	"fmt"
)



func main()  {
	accounts := account.CreateVault(files.NewJsonDb("data.json"))
	for {

		userChoise := getMenu()

		if userChoise == 4 {
			break
		}

		err := menu(userChoise, accounts)

		if err != nil {
			fmt.Println(err)
		}
	}


}

