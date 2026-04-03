package main

import (
	"app/account"
	"fmt"

	"github.com/fatih/color"
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


func findAccount(accounts *account.Vault) {
	var url string
	fmt.Print("Введите URL для поиска: ")
	fmt.Scanln(&url)
	finds := accounts.FindAccountByUrl(url)

	if len(*finds) == 0 {
		color.Red("Не удалось найти акканут")
	}
	
	for _, acc := range *finds {
		acc.PrintAcc()
	}
}


func menu(userChoise int, accounts *account.Vault) error {
	switch userChoise {
	case 1:
		accounts.AddAccount()
	case 2:
		findAccount(accounts)
	case 3:
		for _, value := range accounts.Accounts {
			value.PrintAcc()
		}
	default:
		fmt.Println("Введите подходящее число")
	}


	return nil
}