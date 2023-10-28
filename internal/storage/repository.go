package storage

import "github.com/sugrado/go-rest-api-template/internal/app/users"

type UserRepository interface {
	Save(firstName, lastName, email string) (int, error)
	Find(id int) (*users.User, error)
}
