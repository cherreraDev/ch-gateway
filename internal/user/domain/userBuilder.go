package domain

import "github.com/google/uuid"

type UserBuilder struct {
	id       uuid.UUID
	userName string
	password string
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) WithId(id uuid.UUID) *UserBuilder {
	b.id = id
	return b
}

func (b *UserBuilder) WithUserName(userName string) *UserBuilder {
	b.userName = userName
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.password = password
	return b
}

func (b *UserBuilder) Build() User {
	return User{
		id:       b.id,
		userName: b.userName,
		password: b.password,
	}
}
