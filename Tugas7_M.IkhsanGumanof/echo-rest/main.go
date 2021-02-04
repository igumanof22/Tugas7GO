package main

import (
	db "Tugas7_M.IkhsanGumanof/echo-rest/db"
	routes "Tugas7_M.IkhsanGumanof/echo-rest/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))

}
