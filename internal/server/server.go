package server

import (
	"net/http"
)

type Server struct {
	port    int
	idleTO  int
	readTO  int
	writeTO int
}

func NewServer() *http.Server {
	// r := chi.NewRouter()
	// port, _ := strconv.Atoi(os.Getenv("PORT"))
	// idleTO, _ := strconv.Atoi(os.Getenv("IDLETO"))
	// readTO, _ := strconv.Atoi(os.Getenv("READTO"))
	// writeTO, _ := strconv.Atoi(os.Getenv("WRITTO"))

	// NewServer := &Server{
	// 	port:    port,
	// 	idleTO:  idleTO,
	// 	readTO:  readTO,
	// 	writeTO: writeTO,
	// }

	// Declare Server config
	server := &http.Server{
		// Addr: fmt.Sprintf(":%d", NewServer.port),
		// Handler:      , //maybe declare routes here and use raw server
		// IdleTimeout:  time.Minute,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 30 * time.Second,
	}

	return server
}

func RegisterRoutes()
