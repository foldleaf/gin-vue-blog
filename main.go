package main

import (
	"gin-vue-blog/model"
	"gin-vue-blog/routes"
)

func main() {

	model.InitDb()
	routes.InitRouter()
}
