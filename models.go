package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type app struct {
	Router *mux.Router
	DB     *sql.DB
}

type user struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type certificate struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt int    `json:"createdAt"`
	OwnerID   string `json:"ownerID"`
	Year      int    `json:"year"`
	Note      string `json:"note"`
}

type transfer struct {
	To     string `json:"id"`
	Status string `json:"status"`
}
