package main

import (
	"gin_test/db"
	"gin_test/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
