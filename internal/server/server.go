package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func NewServer(r *gin.Engine) *http.Server {
	// gin.SetMode(gin.ReleaseMode)  //set mode using env
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
