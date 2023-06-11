package service

import (
	"errors"
	"log"

	"github.com/cyee02/ecommerce-service/db"
	"github.com/cyee02/ecommerce-service/helper"
	"github.com/cyee02/ecommerce-service/helper/middleware"
	"github.com/cyee02/ecommerce-service/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var UserService UserInf

// mockgen -source=service/users.go -package=mock -destination=test/mock/users.go
type UserInf interface {
	CreateUser(user *models.User) (*models.UserPublic, error)
	Login(email *string, password *string) (*string, error)
	UpdateUser(req *models.User) (*models.UserPublic, error)
	DeleteUser(user *models.User) (*models.UserPublic, error)
	GetUserById(userId string) (*models.User, error)
}

type userImp struct{}

func SetUserService(service UserInf) {
	UserService = service
}

func InitUserService() {
	var service userImp
	SetUserService(&service)
}

func (*userImp) GetUserById(userId string) (*models.User, error) {
	user, err := db.UserDB.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func EncryptPassword(password string) (*string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return helper.String_ptr(string(encryptedPassword)), nil
}

func (*userImp) CreateUser(user *models.User) (*models.UserPublic, error) {
	// Check if email is in use
	foundUser, err := db.UserDB.GetUserByEmail(user.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("No duplicate email found, proceeding to create user")
	} else if foundUser != nil {
		return nil, errors.New("already have an account with this email")
	} else if err != nil {
		log.Printf("service.CreateUser err=%s", err)
		return nil, err
	}

	// Encrypt password
	encryptedPassword, err := EncryptPassword(*user.Password)
	if err != nil {
		log.Printf("service.CreateUser err=%s", err)
		return nil, err
	}
	user.Password = encryptedPassword

	// Generate user id
	user.UserId = uuid.New().String()

	// Update db
	err = db.UserDB.CreateUser(user)
	if err != nil {
		log.Printf("service.CreateUser err=%s", err)
		return nil, err
	}
	return &models.UserPublic{
		UserId:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func (*userImp) Login(email *string, password *string) (*string, error) {
	// Check if the email address exists
	foundUser, err := db.UserDB.GetUserByEmail(*email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		} else {
			log.Printf("service.Login err=%s", err)
			return nil, err
		}
	}

	// Validate password
	err = bcrypt.CompareHashAndPassword([]byte(*foundUser.Password), []byte(*password))
	if err != nil {
		log.Printf("service.Login err=%s", err)
		return nil, errors.New("invalid username or password")
	}

	// Create token
	token, err := middleware.GenToken(foundUser)
	if err != nil {
		log.Printf("service.Login err=%s", err)
		return nil, errors.New("failed to generate client token")
	}
	return token, nil
}

func (*userImp) UpdateUser(user *models.User) (*models.UserPublic, error) {
	// Check that userId and email matches
	foundUser, err := db.UserDB.GetUserByEmail(user.Email)
	if err != nil {
		log.Printf("service.UpdateUser err=%s", err)
		return nil, err
	}
	if foundUser.UserId != user.UserId {
		return nil, errors.New("Email does not match with UserId")
	}

	// Encrypt password
	encryptedPassword, err := EncryptPassword(*user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = encryptedPassword

	// Update db
	err = db.UserDB.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return &models.UserPublic{
		UserId:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func (*userImp) DeleteUser(user *models.User) (*models.UserPublic, error) {
	// Check that userId and email matches
	foundUser, err := db.UserDB.GetUserByEmail(user.Email)
	if err != nil {
		log.Printf("service.DeleteUser err=%s", err)
		return nil, err
	}
	if foundUser.UserId != user.UserId {
		return nil, errors.New("Email does not match with UserId")
	}

	// Delete data from db
	err = db.UserDB.DeleteUser(user)
	if err != nil {
		log.Printf("service.DeleteUser err=%s", err)
		return nil, err
	}
	return &models.UserPublic{
		UserId:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}
