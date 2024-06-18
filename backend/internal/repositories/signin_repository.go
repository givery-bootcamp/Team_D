package repositories

import (
	"myapp/internal/entities"
	"myapp/internal/external"

	"gorm.io/gorm"
)

type SignInRepository struct {
	Conn *gorm.DB
}

func NewSignInRepository() *SignInRepository {
	return &SignInRepository{
		Conn: external.GetDB(),
	}
}

func (r *SignInRepository) SignIn(username, password string) (*entities.User, error) {
	var user entities.User
	result := r.Conn.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
