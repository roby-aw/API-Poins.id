package store_test

import (
	"api-redeem-point/business/store"
	"os"
	"testing"
)

var service store.Service
var store1, store2, store3 store.Store
var AuthStore store.AuthStore
var input, inputnull store.InputPoin

var token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MywiRW1haWwiOiJ0ZXN0MUBnbWFpbC5jb20iLCJDdXN0b21lciI6dHJ1ZSwiZXhwIjoxNjU2ODcxMzEzfQ.a9O_RMxG7iJR4tMdBZmL6JY2lZKsiQv3rSegnhv1C00"

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestLoginStore(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.LoginStore(&AuthStore)
		if result == nil {
			t.Error("Expect found the result")
		}
	})
}
func TestInputPoin(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.InputPoin(&input)
		if result == nil {
			t.Error("Expect found the result")
		}
	})
	t.Run("Expect found the error", func(t *testing.T) {
		_, err := service.InputPoin(&inputnull)
		if err == nil {
			t.Error("Expect error not nill")
		}
	})
}

func setup() {
	store1.ID = 1
	store1.Email = "store1@gmail.com"
	store1.Alamat = "jl. store 1"
	store1.Store = "store1"
	store1.Password = "passwordstore1"

	store2.ID = 2
	store2.Email = "store2@gmail.com"
	store2.Alamat = "jl. store 2"
	store2.Store = "store2"
	store2.Password = "passwordstore2"

	store3.ID = 3
	store3.Email = "store3@gmail.com"
	store3.Alamat = "jl. store 3"
	store3.Store = "store3"
	store3.Password = "passwordstore3"

	repo := newInMemoryRepository()
	service = store.NewService(&repo)

	AuthStore.Email = store1.Email
	AuthStore.Password = store1.Password

	input.Store_id = 1
	input.Customer_id = 1
	input.Amount = 10000
}

type InMemoryRepository struct {
	Store    map[int]store.Store
	AllStore []store.Store
}

func newInMemoryRepository() InMemoryRepository {
	var repo InMemoryRepository

	repo.Store = make(map[int]store.Store)
	repo.Store[int(store1.ID)] = store1
	repo.Store[int(store2.ID)] = store2
	repo.Store[int(store3.ID)] = store3

	repo.AllStore = []store.Store{}
	repo.AllStore = append(repo.AllStore, store1)
	repo.AllStore = append(repo.AllStore, store2)
	repo.AllStore = append(repo.AllStore, store3)

	return repo
}

func (repo *InMemoryRepository) SignStore(Auth *store.AuthStore) (*store.ResponseLoginStore, error) {
	var Res store.ResponseLoginStore
	for _, v := range repo.AllStore {
		if v.Email == Auth.Email {
			if v.Password == Auth.Password {
				Res.Store.ID = v.ID
				Res.Store.Email = v.Email
				Res.Store.Store = v.Store
				Res.Store.Alamat = v.Alamat
				Res.Store.Password = v.Password
				Res.Token = token
			}
		}
	}
	return &Res, nil
}

func (repo *InMemoryRepository) InputPoin(input *store.InputPoin) (*int, error) {
	price := input.Amount
	var i int
	for i = 0; price >= 100; i = i + 1 {
		price = price - 100
	}
	return &i, nil
}
