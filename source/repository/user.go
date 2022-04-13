package repository

import (
	"app/source/domain/entities"

	"gorm.io/gorm"
)

func InsertUser(user *entities.User) error {
	return db.Create(user).Error
}

func FindUserById(id int) (*entities.User, error) {
	user := entities.User{}
	return &user, db.Where("id = ?", id).First(&user).Error
}

func UpdateUser(user *entities.User) error {
	return db.Model(&user).
		Select("Login", "Password", "Email").
		Updates(entities.User{Login: user.Login, Password: user.Password, Email: user.Email}).Error
}

func DeleteUser(id int) error {
	dbResult := db.Where("id = ?", id).Delete(entities.User{})

	if dbResult.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return dbResult.Error
}
