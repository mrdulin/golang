package repositories

import models "github.com/mrdulin/golang/go-clean-arch/domain/models/user"

type UserRepository interface {
	FindAll() ([]*models.User, error)
}
