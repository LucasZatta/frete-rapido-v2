package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucaszatta/frete-rapido-v2/internal/database"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/handlers"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/repository"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/service"
)

func main() {
	// cfg, err := config.Load(os.Args)
	// if err != nil {
	// 	panic(err)
	// }

	db := database.New() //put it inside env maybe

	quoteRepository := repository.New(db)
	quoteService := service.New(quoteRepository)
	quoteHttp := handlers.New(quoteService)

	// r := chi.NewRouter()

	// r.Get("/products/{id}", quoteHttp.GetQuote)

	// http.ListenAndServe(":8081", r)

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	r := gin.Default()

	r.GET("/", quoteHttp.GetQuote)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
