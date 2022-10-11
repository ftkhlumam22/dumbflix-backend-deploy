package repositories

import (
	"testdumpflix/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	GetUser(ID int) (models.User, error)
	GetUsers(ID int) (models.UsersProfileResponse, error)
	UpdateUser(user models.UsersProfileResponse) (models.UsersProfileResponse, error)
	DeleteUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	// Using Preload("profile") to find data with relation to profile and Preload("Products") for relation to Products here ...
	err := r.db.Preload("Profile").Find(&users).Error

	return users, err
}

func (r *repository) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	// Using Preload("profile") to find data with relation to profile and Preload("Products") for relation to Products here ...
	err := r.db.Preload("Profile").First(&user, ID).Error

	return user, err
}

func (r *repository) GetUsers(ID int) (models.UsersProfileResponse, error) {
	var user models.UsersProfileResponse
	// Using Preload("profile") to find data with relation to profile and Preload("Products") for relation to Products here ...
	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) UpdateUser(user models.UsersProfileResponse) (models.UsersProfileResponse, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
