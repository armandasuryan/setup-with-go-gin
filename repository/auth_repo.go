package repository

import (
	"errors"
	"gin-starter/model"
	"gin-starter/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepoMethod struct {
	db  *gorm.DB
	log *logrus.Logger
}

func AuthRepo(db *gorm.DB, log *logrus.Logger) *AuthRepoMethod {
	return &AuthRepoMethod{db, log}
}

func (r *AuthRepoMethod) GetDataUserRepo(username string) (model.ResponseLogin, error) {
	r.log.Println("Execute function GetDataUserRepo")

	var user model.ResponseLogin
	if err := r.db.Table(utils.TABEL_USER).
		Select(`users.id, users.username, emp.name, emp.email, role.role_name`).
		Joins(`LEFT JOIN employee emp ON users.id_employee = emp.id`).
		Joins(`LEFT JOIN role ON role.id = emp.id_role`).
		Where("users.deleted_at IS NULL and users.username = ?", username).
		Find(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Error("record not found in get data user profile")
			return model.ResponseLogin{}, errors.New("record not found")
		}

		r.log.Error("Failed to get detail data user :", err)
		return model.ResponseLogin{}, err
	}

	return user, nil
}

func (r *AuthRepoMethod) VerifyDataUserRepo(username, passwd string) string {

	var login model.Login

	// check if username exist
	if err := r.db.Table(utils.TABEL_USER).
		Where("username = ?", username).
		First(&login).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Error("User not found")
			return "User not found"
		}
		r.log.Error("Failed to get username in database")
		return "Failed to get username in database"
	}

	// verify password
	if !utils.VerifyPassword(passwd, login.Password) {
		r.log.Error("Password dosn't match")
		return "Password dosn't match"
	}

	return ""
}

func (r *AuthRepoMethod) GetUserProfile(username string) (model.UserProfile, error) {
	r.log.Println("Execute function GetUserProfile")

	var user model.UserProfile
	if err := r.db.Table(utils.TABEL_USER).
		Select(`users.id, emp.name, emp.email`).
		Joins(`LEFT JOIN employee emp ON users.id_employee = emp.id`).
		Where(`users.deleted_at IS NULL and users.username = ?`, username).
		Find(&user).Error; err != nil {
		r.log.Error("Failed to get user profile : ", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Error("record not found in get user profile")
			return model.UserProfile{}, errors.New("record not found")
		}

		r.log.Error("Failed get detail user profile : ", err)
		return model.UserProfile{}, nil
	}

	return user, nil
}
