package main

import (
	"net/http"

	"github.com/jsweet314/go-api-boilerplate/config"
	"github.com/jsweet314/go-api-boilerplate/controllers"
	"github.com/jsweet314/go-api-boilerplate/middleware"
)

func init() {
	config.LoadEnvVars()
}

func main() {
	r := controllers.Startup()
	db := config.ConnectToDatabase()
	defer db.Close()
	http.ListenAndServe(
		":8080",
		&middleware.CORSMiddleware{
			r,
		},
	)
}
