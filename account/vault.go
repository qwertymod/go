package account

import (
	"app/encrypter"
	"encoding/json"
	"strings"
	"time"
	"github.com/fatih/color"
)
type Db interface {
	Write([]byte)
	Read()([]byte, error)
}

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateTime time.Time `json:"updateTime"`
}
type VaultDb struct {
	Vault
	DataB Db
	enc encrypter.Encrypter
}


func (vault *Vault) FindAccount(checker func(Account, string)bool, str string) *[]Account {
	var acc []Account
	for _, value := range vault.Accounts {
		isMatched := checker(value, str)
		if isMatched {
			acc = append(acc, value)
		}
	}

	return &acc
}


func CheckerUrl(account Account, url string) bool {
	return strings.Contains(account.Link, url)
}

func CheckeLogin(account Account, login string) bool {
	return strings.Contains(account.Login, login)
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


func CreateVault(dataB Db, enc encrypter.Encrypter) (*VaultDb) {
	data, err := dataB.Read()
	if err != nil {
	return  &VaultDb {
		Vault : Vault {
			Accounts : make([]Account, 0 , 10),
			UpdateTime : time.Now(),
		},
		enc : enc,
		DataB: dataB,
		}
	}
	data = enc.Decrypter(data)
	var val = VaultDb{DataB: dataB, enc : enc}
	err = json.Unmarshal(data, &val.Vault)
	if err != nil {
		color.Black(err.Error())
	}
	return &val
}


func (vault *VaultDb) AddAccount ()  {
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
	data = vault.enc.Encrypt(data)
	vault.DataB.Write(data)

}


func (vault *Vault) ToBytes() ([]byte, error){
	data, err := json.Marshal(*vault)

	if err != nil {
		return nil, err
	}

	return data , nil
}

