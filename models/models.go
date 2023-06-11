package models

import "time"

type User struct {
	UserId    string  `gorm:"column:user_id; primary_key" json:"user_id"`
	Password  *string `gorm:"column:password; type: varchar(255)" json:"password" validate:"required"`
	FirstName *string `gorm:"column:first_name; type: varchar(50)" json:"first_name" validate:"required"`
	LastName  *string `gorm:"column:last_name; type: varchar(50)" json:"last_name" validate:"required"`
	Phone     *string `gorm:"column:phone; type: varchar(20)" json:"phone" validate:"required"`
	Email     string  `gorm:"column:email; type: varchar(255)" json:"email" validate:"email,required"`
	// Addresses         []Address `json:"addresses" validate:"required"`
	// Cart              []Product `json:"cart"`
	// Orders            []Order   `json:"orders"`
}

type UserPublic struct {
	UserId    string  `json:"user_id"`
	FirstName *string `json:"first_name" validate:"required"`
	LastName  *string `json:"last_name" validate:"required"`
	Email     string  `json:"email" validate:"email,required"`
}

type CreateUserReq struct {
	User User `json:"user" validate:"required"`
}

type CreateUserResp struct {
	User UserPublic `json:"user"`
}

type LoginReq struct {
	Email    string  `json:"email" validate:"email,required"`
	Password *string `gorm:"column:password; type: varchar(255)" json:"password" validate:"required"`
}

type LoginResp struct {
	Token string `json:"token" validate:"required"`
}

type GetUserByIdReq struct {
	UserId string `json:"user_id"`
}

type GetUserByIdResp struct {
	UserId    string  `json:"user_id"`
	FirstName *string `json:"first_name" validate:"required"`
	LastName  *string `json:"last_name" validate:"required"`
	Email     string  `json:"email" validate:"email,required"`
}

type Address struct {
	Address    *string `json:"address"`
	Country    *string `json:"country"`
	PostalCode *string `json:"portal_code"`
}

type Product struct {
	ProductId   string  `json:"product_id"`
	ProductName *string `json:"product_name"`
	Price       *int    `json:"price"`
	Rating      *int    `json:"rating"`
	Image       *string `json:"image"`
}

type Order struct {
	OrderId        string    `json:"order_id"`
	OrderCart      []Product `json:"order_cart"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Price          int       `json:"price"`
	Payment_Method int       `json:"payment_method"`
	Status         string    `json:"status"`
}
