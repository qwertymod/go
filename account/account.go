package  account


import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type Account struct{
	login 		string
	password 	string
	link 		string
}

type accountBonus struct {
	timeCreate time.Time
	timeUpdate time.Time
	Account 
}

func (acc *Account) PrintAcc() {
	fmt.Printf("Ваш логин: %s \nВаш пароль: %s \nВаша ссылка: %s", acc.login, acc.password, acc.link)
} // сделал из функции метод для аккаунта

func (acc *Account) createPassword(n int) {
	symbols := []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-+*!#@")
	password := make([]rune,  n)
	for index := range password{
		password[index] = symbols[rand.IntN(len(symbols))]
	}

	acc.password = string(password)
}

func CreateAcc() (*Account, error) {
	var login, password, link string

	login, password, link = getUserAcc()

	if login == "" {
		return nil, errors.New("Invalid login")
	}

	_, err := url.ParseRequestURI(link)
	if err != nil {
		return nil, err
	}

	acc := &Account{
		login : login,
		password : password,
		link : link,
	}

	if password == ""{
		acc.createPassword(12)
	}

	return acc, nil
}


func getUserAcc() (login, password, link string) {
	fmt.Print("Input login: ")
	fmt.Scanln(&login)
	fmt.Print("Input password: ")
	fmt.Scanln(&password)
	fmt.Print("Input link: ")
	fmt.Scanln(&link)

	return login, password, link
}

