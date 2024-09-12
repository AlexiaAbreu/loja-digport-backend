package main

import (
	"github.com/AlexiaAbreu/loja-digport-backend/db"
	_ "github.com/lib/pq"
)

func main() {
	db.InitDB()
	StartServer()

}
