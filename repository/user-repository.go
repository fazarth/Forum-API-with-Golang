package repository

import (
	"log"

	"buddyku/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(user model.User) model.User
	UpdateUser(user model.User) model.User
	VerifyCredential(username string, password string) interface{}
	IsDuplicateUserName(username string) (tx *gorm.DB)
	FindByUserName(username string) model.User
	ProfileUser(userID string) model.User
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user model.User) model.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user model.User) model.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser model.User
		db.connection.Find(&tempUser, user.UserID)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(username string, password string) interface{} {
	var user model.User
	// res := db.connection.Where("user_name = ?", user_name).Take(&user)
	res := db.connection.Raw("SELECT users.user_id, users.user_name, users.password, users.comment, users.is_usable, users.create_user, users.create_date, users.update_user, users.update_date FROM users WHERE users.user_name = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateUserName(username string) (tx *gorm.DB) {
	var user model.User
	return db.connection.Where("user_name = ?", username).Take(&user)
}

func (db *userConnection) FindByUserName(username string) model.User {
	var user model.User
	db.connection.Where("user_name = ?", username).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) model.User {
	var user model.User
	db.connection.Preload("Employees").Preload("Employees.User").Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
