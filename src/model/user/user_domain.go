package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
)

type userDomanain struct {
	id       string  
	email    string 
	password string 
	name     string 
	age      int8   
}

func (ud *userDomanain) GetJSONValue() (string, error) {
	json, err := json.Marshal(ud)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(json), nil
}
/*
Domain não deve ser exportavel, para isso temos os objetos request e response
Domain é para se seguir a regra de negócio
O controller não chama diretamente o domain contudo chama uma interface dele
Representação no banco de dados
*/

func (ud *userDomanain) GetEmail() string {
	return ud.email
}

func (ud *userDomanain) GetPassword() string {
	return ud.password
}

func (ud *userDomanain) GetName() string {
	return ud.name
}

func (ud *userDomanain) GetAge() int8 {
	return int8(ud.age)
}

func (ud *userDomanain) GetId() string {
	return ud.id
}

func (ud *userDomanain) SetID(id string) {
	ud.id = id
}

func (ud *userDomanain) SetName(name string) {
	ud.name = name
}

func (ud *userDomanain) SetAge(age int8) {
	ud.age = age
}

func NewUserDomain(email string, password string, name string, age int8) UserDomanainInterface {
	return &userDomanain{
		email: email,
		password: password,
		name: name,
		age: age,
	}
}

func NewUserDomainUpdate( name string, age int8) UserDomanainInterface {
	return &userDomanain{
		name: name,
		age: age,
	}
}

func NewUserDomainLogin(email string, password string) UserDomanainInterface {
	return &userDomanain{
		email: email,
		password: password,
	}
}

func (ud *userDomanain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomanainInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetJSONValue() (string, error)
	EncryptPassword()
	SetID(string)
	SetName(string)
	SetAge(int8)
	GenerreteTokeJwt() (string, *rest_err.RestErr)
}