package main

import (
	"fmt"
	"food_ordering/domain"
)

func main() {
	client := domain.CreateUser("John", "Doe", 30, 1000)
	fmt.Println(client.Balance())
	client.AddMoney(100)
	fmt.Println(client.Balance())
}
