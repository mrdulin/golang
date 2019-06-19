package repositories

import (
	"go-clean-arch/domain/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
}
