package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct{
	Login 		string `json:"login" xml:"test" ` // Тег — это буквально мост между именем в JSON и именем поля в Go-структуре.
	Password 	string `json:"password"`
	Link 		string `json:"link"`
	TimeCreate time.Time `json:"timeCreate"`
	TimeUpdate time.Time `json:"timeUpdate"`
}


func (acc *Account) ToBytes() ([]byte, error){
	file, err := json.Marshal(acc)

	if err != nil {
		return nil, err
	}

	return file , nil
}


func (acc *Account) PrintAcc() {
	color.Cyan("Ваш логин: %s \nВаш пароль: %s \nВаша ссылка: %s", acc.Login, acc.Password, acc.Link)
} // сделал из функции метод для аккаунта

func (acc *Account) createPassword(n int) {
	symbols := []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-+*!#@")
	password := make([]rune,  n)
	for index := range password{
		password[index] = symbols[rand.IntN(len(symbols))]
	}

	acc.Password = string(password)
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
		Login : login,
		Password : password,
		Link : link,
		TimeCreate : time.Now(),
		TimeUpdate : time.Now(),
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

