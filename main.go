package main

import (
	"todo_list/conf"
	"todo_list/routers"
)

func main() {
	conf.Init()
	r := routers.NewRouter()
	r.Run(":8888")
}
