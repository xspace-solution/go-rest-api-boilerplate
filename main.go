package main

import (
	"net/http"

	"github.com/katochojiro/go-api-boilerplate/config"
	"github.com/katochojiro/go-api-boilerplate/controllers"
	"github.com/katochojiro/go-api-boilerplate/middleware"
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
			Next: r,
		},
	)
}
