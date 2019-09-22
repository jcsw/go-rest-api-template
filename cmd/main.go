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

	uuid "github.com/google/uuid"

	mariadb "github.com/jcsw/go-rest-api-template/pkg/database/mariadb"
	mongodb "github.com/jcsw/go-rest-api-template/pkg/database/mongodb"
	"github.com/jcsw/go-rest-api-template/pkg/hello"
	monitor "github.com/jcsw/go-rest-api-template/pkg/monitor"
	sys "github.com/jcsw/go-rest-api-template/pkg/system"
)

var healthy int32
var env string

func main() {
	sys.Info("Server initializing...")

	flag.StringVar(&env, "env", "prod", "Environment")
	flag.Parse()

	sys.LoadProperties(env)

	mariadb.Connect()
	mongodb.Connect()

	port := sys.Properties.ServerPort

	router := http.NewServeMux()
	router.HandleFunc("/health", health)
	router.HandleFunc("/monitor", monitor.Handler)

	router.HandleFunc("/", hello.Handler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      tracing()(router),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		sys.Info("Server shutting down...")

		atomic.StoreInt32(&healthy, 0)

		mariadb.Disconnect()
		mongodb.Disconnect()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			sys.Fatal("Could not gracefully shutdown the server, err:%v", err)
		}

		sys.Info("Shutdown complete, bye bye.")
	}()

	atomic.StoreInt32(&healthy, 1)
	sys.Info("Server is ready to handle requests at port [%d].", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		sys.Fatal("Could not listen on port [%d], err:%v", port, err)
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func wrapResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func tracing() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = uuid.New().String()
			}

			sys.Info("--> [%s] request [%s] on [%s] by %s", requestID, r.Method, r.URL.Path, r.RemoteAddr)

			ctx := context.WithValue(r.Context(), sys.ContextKeyRequestID, requestID)
			w.Header().Set("X-Request-Id", requestID)

			wrw := wrapResponseWriter(w)
			next.ServeHTTP(wrw, r.WithContext(ctx))

			sys.Info("<-- [%s] response [%d] in [%v]", requestID, wrw.statusCode, time.Since(start))
		})
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}
