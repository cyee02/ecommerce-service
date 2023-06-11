package db

import (
	"log"

	"github.com/cyee02/ecommerce-service/helper/config"
	"github.com/cyee02/ecommerce-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UserDB DbInf

// mockgen -source=db/db.go -package=mock -destination=test/mock/userDb.go
type DbInf interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(userId string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type dbImp struct {
	client *gorm.DB
}

func SetDbService(service DbInf) {
	UserDB = service
}

func InitMySQL() {
	dsn := config.Config.DbUsername + ":" + config.Config.DbPassword + "@tcp(" + config.Config.DbHost + ":" + config.Config.Port + ")/" + config.Config.DbName
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	log.Printf("Connected to host [%s] at port [%s] to db name [%s]", config.Config.DbHost, config.Config.Port, config.Config.DbName)
	SetDbService(dbImp{client: database})
}

func (db dbImp) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.client.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("db.GetUserByEmail err=%s", err)
		return nil, err
	}
	return &user, nil
}

func (db dbImp) GetUserById(userId string) (*models.User, error) {
	var user models.User
	err := db.client.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		log.Printf("db.GetUserById err=%s", err)
		return nil, err
	}
	return &user, nil
}

func (db dbImp) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := db.client.Find(&users).Error
	if err != nil {
		log.Printf("db.GetAllUsers err=%s", err)
		return nil, err
	}
	return users, nil
}

func (db dbImp) CreateUser(user *models.User) error {
	err := db.client.Create(&user).Error
	if err != nil {
		log.Printf("db.CreateUser err=%s", err)
		return err
	}
	return nil
}

func (db dbImp) UpdateUser(user *models.User) error {
	err := db.client.Save(&user).Error
	if err != nil {
		log.Printf("db.UpdateUser err=%s", err)
		return err
	}
	return nil
}

func (db dbImp) DeleteUser(user *models.User) error {
	err := db.client.Delete(&user).Error
	if err != nil {
		log.Printf("db.DeleteUser err=%s", err)
		return err
	}
	return nil
}
