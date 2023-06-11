package helper

import (
	"encoding/json"
	"log"

	"github.com/cyee02/ecommerce-service/models"
)

func ConvertUserArrayToPublic(users []*models.User) []*models.UserPublic {
	usersPublic := make([]*models.UserPublic, len(users))
	for i, user := range users {
		userPublic := &models.UserPublic{
			UserId:    user.UserId,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}
		usersPublic[i] = userPublic
	}
	return usersPublic
}

func CastAnyToUser(anyType any) (*models.User, error) {
	var user models.User
	data, err := json.Marshal(anyType)
	if err != nil {
		log.Panicf("Marhshalling err: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Fatalf("Unmarshalling err : %+v", err)
		return nil, err
	}
	return &user, nil
}
