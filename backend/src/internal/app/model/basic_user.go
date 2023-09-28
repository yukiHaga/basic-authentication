package model

type BasicUserId int64

// PasswordとConfirmationはRailsの家蔵属性として入っていたから、一応入れておいた
type BasicUser struct {
	Name     string
	Password string
}

func NewBasicUser(name, password string) *BasicUser {
	return &BasicUser{
		Name:     name,
		Password: password,
	}
}
