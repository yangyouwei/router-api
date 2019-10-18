package main

import (
	db "github.com/yangyouwei/router-api/databases"
	. "github.com/yangyouwei/router-api/router"
)
func main()  {
	defer db.SqlDB.Close()

	router := InitRouter()
	router.Run(":8080")
}
