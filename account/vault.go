package account

import (
	"app/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)



type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateTime time.Time `json:"updateTime"`
}


func (vault *Vault) FindAccountByUrl(url string) *[]Account {
	var acc []Account
	for _, value := range vault.Accounts {
		isMatched := strings.Contains(value.Link, url)
		if isMatched {
			acc = append(acc, value)
		}
	}

	return &acc
}

func (vault *Vault) DeleteAcc(url string) bool {
	isDel := false
	for i := 0 ; i < len(vault.Accounts); {
		isMatched := strings.Contains(vault.Accounts[i].Link, url)
		if isMatched {
			vault.Accounts = append(vault.Accounts[:i], vault.Accounts[i+1:]...)
			isDel = true
		} else {
			i++
		}
	}

	return isDel
}


func CreateVault() (*Vault) {
	file, err := files.ReadFile("data.json")
	if err != nil {
	return  &Vault {
		Accounts : make([]Account, 0 , 10),
		UpdateTime : time.Now(),
		}
	}
	var val = Vault{}
	err = json.Unmarshal(file, &val)
	if err != nil {
		color.Red(err.Error())
	}
	return &val
}


func (vault *Vault) AddAccount ()  {
	acc, err := CreateAcc()
	if err != nil {
		color.Red(err.Error())
	}
	vault.Accounts = append(vault.Accounts, *acc)
	vault.UpdateTime = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteFile(data, "data.json")
}


func (acc *Vault) ToBytes() ([]byte, error){
	data, err := json.Marshal(*acc)

	if err != nil {
		return nil, err
	}

	return data , nil
}

