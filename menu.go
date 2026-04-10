package main

import (
	"app/account"
	"app/files"
	"fmt"
	"strings"

	"github.com/fatih/color"
)
func getMenu() (userChoise int){
	fmt.Println(`	1 - создать аккаунт
	2 - найти аккаунт
	3 - удалить аккаунт
	4 - выход
Выберите вариант: `)
	fmt.Scanln(&userChoise)
	return 
}


func findAccount(accounts *account.Vault) {
	var url string
	fmt.Print("Введите URL для поиска: ")
	fmt.Scanln(&url)
	finds := accounts.FindAccount(func(acc account.Account, str string) bool {
		return strings.Contains(acc.Link, str)
	}, url)

	if len(*finds) == 0 {
		color.Red("Не удалось найти акканут")
	}
	
	for _, acc := range *finds {
		acc.PrintAcc()
	}
}

func deleteAcc(vault *account.Vault)  {
	var url string 
	fmt.Print("Введите url для удаления: ")
	fmt.Scanln(&url)

	isDel := vault.DeleteAcc(url)
	if isDel {
		color.Green("Акканут удален")
	} else {
		color.Red("Аккаунт не найден")
	}

	data, _ := vault.ToBytes()
	db := files.NewJsonDb("data.json")
	db.Write(data)
}


func menu(userChoise int, accounts *account.VaultDb) error {

	switch userChoise {
	case 1:
		accounts.AddAccount()
	case 2:
		findAccount(&accounts.Vault)
	case 3:
		deleteAcc(&accounts.Vault)
	}
	
	return nil
}

