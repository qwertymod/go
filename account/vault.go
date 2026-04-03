package account

import (
	"app/files"
	"encoding/json"
	"time"

	"github.com/fatih/color"
)



type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateTime time.Time `json:"updateTime"`
}


type slot struct {
	Login string
	password string
	Massage []string
}

func CreateSlot(a, b, c string) (d slot) {
	d.Login = a
	d.password = b
	d.Massage = make([]string, 0, 1)
	d.Massage = append(d.Massage, c)
	return 
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

