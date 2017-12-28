package repositories

import (
	"github.com/jmoiron/sqlx"
	"go-clean-arch/domain/models"
)
import "go-clean-arch/domain/repositories"

type userRepository struct {
	Db *sqlx.DB
}

func NewUserRepository(Db *sqlx.DB) repositories.UserRepository {
	return &userRepository{Db}
}

func (ur *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	query := "select user_nme from users"
	err := ur.Db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
