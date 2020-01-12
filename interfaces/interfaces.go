package interfaces

import "github.com/picbook1/models"

type UserInterface interface {
	Users() ([]models.User, []error)
	User(id uint) (*models.User, []error)
	UpdateUser(comment *models.User) (*models.User, []error)
	DeleteUser(id uint) (*models.User, []error)
	Close()
}



