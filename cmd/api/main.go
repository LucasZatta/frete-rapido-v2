package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucaszatta/frete-rapido-v2/internal/database"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/handlers"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/repository"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/service"
	"github.com/lucaszatta/frete-rapido-v2/internal/server"
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

	r := gin.Default()

	r.POST("/quote", quoteHttp.SimulateQuote)
	r.GET("/metrics", quoteHttp.GetQuotes)

	server := server.NewServer(r)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
