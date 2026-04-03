package main

import (
	"app/account"
	"fmt"
)



func main()  {
	accounts := account.CreateVault()
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

