package main

import (
	"net/http"

	"sysdevops.biz/admin/models"
	"sysdevops.biz/config"
)

func main() {
	models.Post{}.Migrate()
	models.User{}.Migrate()
	models.Category{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}
