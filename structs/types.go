package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



type User struct {	
	ID primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	UserName string `bson:"user_name"`
	Email string `bson:"email"`
	FirstName string `bson:"first_name"`
	SurName string `bson:"sur_name"`
}

type UserDTO struct {
	ID string `json:"id"`
	UserName string `json:"userName"`
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	SurName string `json:"surName"`
}

type Book struct {
	ID primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Isbn string `bson:"isbn"`
	Author string `bson:"author"`
	Stock int `bson:"stock"`
	Price float64 `bson:"price"`
}

type BookDTO struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Isbn string `json:"isbn"`
	Author string `json:"author"`
	Stock int `json:"stock"`
	Price float64 `json:"price"`
}


type Order struct {
	ID primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Status string `bson:"status"`
	Address string `bson:"address"`
	PaymentMethod string `bson:"payment_method"`
	TotalPrice float64 `bson:"total_price"`
	Vat float64 `bson:"vat"`
}

type OrderLine struct {
	ID primitive.ObjectID `bson:"_id"`
	OrderID string `bson:"order_id"`
	BookID string `bson:"book_id"`
	Quantity int `bson:"quantity"`
	Price float64 `bson:"price"`
}

type OrderDTO struct {
	ID string `json:"id"`
	Status string `json:"status"`
	Address string `json:"address"`
	PaymentMethod string `json:"paymentMethod"`
	TotalPrice float64 `json:"totalPrice"`
	Vat float64 `json:"vat"`
	Lines []OrderLineDTO `json:"lines"`
}

type OrderLineDTO struct {
	Book []BookDTO `json:"book"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
}