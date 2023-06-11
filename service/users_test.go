package service

import (
	"testing"

	"github.com/cyee02/ecommerce-service/db"
	"github.com/cyee02/ecommerce-service/helper"
	"github.com/cyee02/ecommerce-service/models"
	"github.com/cyee02/ecommerce-service/test/mock"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"gotest.tools/assert"
)

func Test_CreateUser(t *testing.T) {
	InitUserService()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mock.NewMockDbInf(ctrl)
	db.SetDbService(mockDb)

	// Set User variables
	email := "abcd1234"
	password := helper.String_ptr("Newpassword")
	firstName := helper.String_ptr("John")
	lastName := helper.String_ptr("Smith")
	phone := helper.String_ptr("12345678")

	tests := []struct {
		name  string
		args  *models.User
		setup func(t *testing.T)
		want  *models.User
	}{
		{
			name: "Successfully Creates User",
			args: &models.User{
				Email:     email,
				Password:  password,
				FirstName: firstName,
				LastName:  lastName,
				Phone:     phone,
			},
			setup: func(t *testing.T) {
				mockDb.EXPECT().GetUserByEmail(gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
				mockDb.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
			want: &models.User{
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(t)
			user, _ := UserService.CreateUser(test.args)
			assert.Equal(t, test.want.FirstName, user.FirstName)
			assert.Equal(t, test.want.LastName, user.LastName)
			assert.Equal(t, test.want.Email, user.Email)
		})
	}
}

func Test_LoginUser(t *testing.T) {
	InitUserService()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mock.NewMockDbInf(ctrl)
	db.SetDbService(mockDb)

	// Set User variables
	user := &models.User{
		Email:    "abcd1234",
		Password: helper.String_ptr("Newpassword"),
	}
	encryptedPassword, _ := EncryptPassword(*user.Password)

	tests := []struct {
		name  string
		args  *models.User
		setup func(t *testing.T)
		want  interface{}
	}{
		{
			name: "Successfully Login",
			args: user,
			setup: func(t *testing.T) {
				mockDb.EXPECT().GetUserByEmail(gomock.Any()).Return(&models.User{Email: user.Email, Password: encryptedPassword}, nil)
			},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(t)
			_, err := UserService.Login(&user.Email, user.Password)
			assert.Equal(t, test.want, err)
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	InitUserService()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mock.NewMockDbInf(ctrl)
	db.SetDbService(mockDb)

	// Set User variables
	user := &models.User{
		Email:    "abcd1234",
		Password: helper.String_ptr("Newpassword"),
	}
	encryptedPassword, _ := EncryptPassword(*user.Password)

	tests := []struct {
		name  string
		args  *models.User
		setup func(t *testing.T)
		want  interface{}
	}{
		{
			name: "Successfully Login",
			args: user,
			setup: func(t *testing.T) {
				mockDb.EXPECT().GetUserByEmail(gomock.Any()).Return(&models.User{Email: user.Email, Password: encryptedPassword}, nil)
			},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(t)
			_, err := UserService.Login(&user.Email, user.Password)
			assert.Equal(t, test.want, err)
		})
	}
}

func Test_DeleteUser(t *testing.T) {
	InitUserService()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mock.NewMockDbInf(ctrl)
	db.SetDbService(mockDb)

	// Set User variables
	user := &models.User{
		UserId:    "1234",
		Email:     "abcd1234",
		FirstName: helper.String_ptr("John"),
		LastName:  helper.String_ptr("Smith"),
	}

	tests := []struct {
		name  string
		args  *models.User
		setup func(t *testing.T)
		want  *models.User
	}{
		{
			name: "Successfully delete user",
			args: user,
			setup: func(t *testing.T) {
				mockDb.EXPECT().GetUserByEmail(gomock.Any()).Return(user, nil)
				mockDb.EXPECT().DeleteUser(gomock.Any()).Return(nil)
			},
			want: user,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(t)
			user, _ := UserService.DeleteUser(user)
			assert.Equal(t, test.want.UserId, user.UserId)
			assert.Equal(t, test.want.Email, user.Email)
			assert.Equal(t, test.want.FirstName, user.FirstName)
			assert.Equal(t, test.want.LastName, user.LastName)
		})
	}
}
