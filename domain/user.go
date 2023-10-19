package domain

type user struct {
	name     string
	lastName string
	age      int64
	balance  int64
}

func CreateUser(name, lastName string, age, balance int64) *user {
	return &user{
		name:     name,
		lastName: lastName,
		age:      age,
		balance:  balance,
	}
}

func (u *user) Name() string {
	return u.name
}

func (u *user) AddMoney(amount int64) {
	u.balance += amount
}

func (u *user) Balance() int64 {
	return u.balance
}
