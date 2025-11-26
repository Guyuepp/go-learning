package bank

import (
	"errors"
	"sync"
)

type BankAccount struct {
	id       int
	password string
	money    float64
	mu       sync.Mutex
}

var m sync.Map

var nextId = struct {
	sync.Mutex
	val int
}{
	val: 1,
}

func AdjuestMoney(id int, password string, money float64) (rest float64, err error) {
	v, ok := m.Load(id)
	if !ok {
		return 0, errors.New("account not exists")
	}
	acc := v.(*BankAccount)
	if acc.password != password {
		return 0, errors.New("wrong password")
	}
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.money += money
	return acc.money, nil
}

func Signup(password string, money float64) (id int, err error) {
	nextId.Lock()
	id = nextId.val
	nextId.val++
	nextId.Unlock()
	acc := &BankAccount{
		id:       id,
		password: password,
		money:    money,
		mu:       sync.Mutex{},
	}
	m.Store(id, acc)
	return id, nil
}
