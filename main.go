package main

import (
	"./sqlDriver"
	//"fmt"
)

func main() {
	var (
		host     string
		user     string
		port     string
		_type    string
		database string
		pass     string
	)
	host = "127.0.0.1"
	user = "root"
	port = "3306"
	_type = "tcp"
	database = "test"
	pass = "root"
	d := sqlDrvier.Driver{}.SetDriver(port, host, user, pass, database, _type)
	d.Select()
}
