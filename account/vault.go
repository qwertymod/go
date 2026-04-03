package account

import (
	"encoding/json"
	"time"
)



type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateTime time.Time `json:"updateTime"`
}



func CreateVault() (Vault) {
	val := Vault {
		UpdateTime : time.Now(),
	}

	return val
}


func (acc *Vault) ToBytes() ([]byte, error){
	data, err := json.Marshal(*acc)

	if err != nil {
		return nil, err
	}

	return data , nil
}

