package demo

import (
	"fmt"
	"time"
)


type Account struct {
	ID          string `bson:"_id"`
	Phone       string  `bson:"phone"`
	CountryCode string            `bson:"countryCode"`
	CreateAt    time.Time         `bson:"createAt"`
}


func Add(msg string){
	acc := Account{
		CreateAt:time.Now(),
	}
	fmt.Println(acc, acc.ID, acc.CreateAt)
	fmt.Println(msg)
}