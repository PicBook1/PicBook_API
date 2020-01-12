package interfaces

import "github.com/picbook1/models"

type UserInterface interface {
	Users() ([]models.User, []error)
	User(id uint) (*models.User, []error)
	UpdateUser(comment *models.User) (*models.User, []error)
	DeleteUser(id uint) (*models.User, []error)
	Close()
}
type GalleryInterface interface {
	Categories() ([]models.Gallery, []error)
	Gallery(id uint) (*models.Gallery, []error)
	UpdateGallery(gallery *models.Gallery) (*models.Gallery, []error)
	DeleteGallery(id uint) (*models.Gallery, []error)
	StoreGallery(gallery *models.Gallery) (*models.Gallery, []error)
	SearchGallery(title string) (models.Gallery, []error)
	
}


