package user

import (
	// roles "schollinggo/modules/Master/role"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByUserName(username string) (User, error)
	FindByID(ID uint64) (User, error)
	Update(user User) (User, error)
	FindAll() ([]User, error)
	DeleteByUuid(Uuid string) ([]User, error)
	FindByUuid(Uuid string) (User, error)
	DeleteByid(id uint64) ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	// return user, nil
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUserName(user_name string) (User, error) {
	var user User

	err := r.db.Debug().Where("user_name = ?", user_name).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID uint64) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUuid(Uuid string) (User, error) {
	var user User

	err := r.db.Where("uuid = ?", Uuid).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Debug().Updates(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) DeleteByUuid(Uuid string) ([]User, error) {
	var users []User

	err := r.db.Where("uuid = ?", Uuid).Delete(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) DeleteByid(id uint64) ([]User, error) {
	var users []User

	err := r.db.Debug().Where("id = ?", id).Delete(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}
