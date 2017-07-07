package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)


func main()  {
	db, _ := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "",
		Database: "postgres",
	})

defer db.Close()
	pg.Scan()

}