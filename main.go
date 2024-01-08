package main

import (
	db "MyMoneyAPI/database"
	rt "MyMoneyAPI/routes"
)

func main() {

	db.DbConnect()
	rt.HandleRoutes()

}
