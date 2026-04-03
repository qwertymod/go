package main

import (
	"app/account"
	"app/files"
	"fmt"
)
func getMenu() (userChoise int){
	fmt.Println(`	1 - создать аккаунт
	2 - найти аккаунт
	3 - вывести все аккаунты
	4 - выход
Выберите вариант: `)
	fmt.Scanln(&userChoise)
	return 
}


func menu(userChoise int, accounts *account.Vault) error {
	switch userChoise {
	case 1:
		acc, err := account.CreateAcc()
		if err != nil {
			return err
		}
		accounts.Accounts = append(accounts.Accounts, *acc)
		data, err := accounts.ToBytes()
		files.WriteFile(data, "data.json")
	case 2:
		var login string
		fmt.Print("Введите логин: ")
		fmt.Scanln(&login)
		for _, value := range accounts.Accounts {
			if value.Login == login {
				value.PrintAcc()
			}
		}
	case 3:
		for _, value := range accounts.Accounts {
			value.PrintAcc()
		}
	default:
		fmt.Println("Введите подходящее число")
	}


	return nil
}