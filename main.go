package main

import (
	"app/account"
	"app/encrypter"
	"app/files"
	"fmt"

	"github.com/joho/godotenv"
)



func main()  {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось найти env файл")
	}
	accounts := account.CreateVault(files.NewJsonDb("data.key"), *encrypter.NewEncrypter())
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

