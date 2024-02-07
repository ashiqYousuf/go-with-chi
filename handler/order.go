package handler

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create method")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list method")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetByID method")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateByID method")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteByID method")
}
