package repository

import "app/source/domain/entities"

func InsertUser(user *entities.User) error {
	return db.Create(user).Error
}
