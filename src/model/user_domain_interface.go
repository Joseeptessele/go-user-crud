package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() string
	SetID(string)
	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		name:     name,
		password: password,
		age:      age,
		id:       "",
	}
}
