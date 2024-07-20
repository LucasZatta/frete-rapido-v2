package main

import (
	"github.com/lucaszatta/frete-rapido-v2/internal/database"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/http"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/repository"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/service"
)

type env struct {
	//db connection
	//controller
}

func main() {
	// cfg, err := config.Load(os.Args)
	// if err != nil {
	// 	panic(err)
	// }

	db := database.New() //put it inside env maybe

	quoteRepository := repository.New(db.Db)
	quoteService := service.New(quoteRepository)
	quoteHttp := http.New(quoteService)

}
