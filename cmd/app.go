package main

import (
	context "context"
	flag "flag"
	fmt "fmt"
	http "net/http"
	os "os"
	signal "os/signal"
	atomic "sync/atomic"
	time "time"

	mariadb "go-rest-api-template/pkg/infra/database/mariadb"
	mongodb "go-rest-api-template/pkg/infra/database/mongodb"
	sys "go-rest-api-template/pkg/infra/system"
)

var healthy int32
var env string
var serverPort int

func main() {
	sys.Info("[Server initializing]")

	flag.StringVar(&env, "env", "prod", "Environment")
	flag.Parse()

	sys.LoadProperties(env)
	serverPort = sys.Properties.ServerPort

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	server := buildServer()

	mariadb.Connect()
	mongodb.Connect()

	startServer(server)
	<-quit
	stopServer(server)
}

func buildServer() *http.Server {
	mux := http.NewServeMux()

	handleRouters(mux)

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", serverPort),
		Handler:      tracing()(mux),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
}

func startServer(server *http.Server) {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sys.Fatal("[Could not listen on port %d] err:%v", serverPort, err)
		}
	}()

	atomic.StoreInt32(&healthy, 1)
	sys.Info("[Server is ready to handle requests at port %d]", serverPort)
}

func stopServer(server *http.Server) {
	sys.Info("[Server shutting down]")

	atomic.StoreInt32(&healthy, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		mariadb.Disconnect()
		mongodb.Disconnect()
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		sys.Fatal("[Could not gracefully shutdown the server] err:%v", err)
	}

	sys.Info("[Server shutdown complete]")
}
