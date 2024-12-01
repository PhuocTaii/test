package server

import (
	"context"
	"errors"
	"fmt"
	_ "ia-03-backend/docs"
	"ia-03-backend/pkg/http/middleware"
	"ia-03-backend/pkg/http/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option func(*HTTPServer)

type HTTPServer struct {
	Name                    string
	Port                    int64
	StrictSlash             bool
	Routes                  []route.Route
	GroupRoutes             []route.GroupRoute
	Middlewares             []func(c *gin.Context)
	GinOptions              []route.GinOption
	GracefulShutdownTimeout time.Duration
	OnCloseFunc             func()
}

func NewHTTPServer(options ...Option) *HTTPServer {
	s := &HTTPServer{}
	for _, option := range options {
		option(s)
	}
	return s
}

func (server *HTTPServer) Run() {
	// Request ID middleware
	server.Middlewares = append(server.Middlewares, middleware.RequestId())
	server.Middlewares = append(server.Middlewares, middleware.Recover())
	server.Middlewares = append(server.Middlewares, middleware.CORSMiddleware())

	// Setup route
	r := route.NewGin(
		route.AddMiddlewares(server.Middlewares...),
		route.AddHealthCheckRoute(),
		route.AddNoRouteHandler(),
		route.SetStrictSlash(server.StrictSlash),
		route.SetMaximumMultipartSize(10000000),
		route.AddGroupRoutes(server.GroupRoutes),
		route.AddRoutes(server.Routes),
		route.AddGinOptions(server.GinOptions...),
	)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Create an HTTP server instance with the Gin engine
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", server.Port),
		Handler: r,
	}

	// Channel to listen for interrupt signals (e.g., CTRL+C, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		fmt.Printf("http server started on port %v", server.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("could not start server: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	<-quit
	fmt.Printf("shutting down server...")

	// Graceful shutdown with a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	fmt.Printf("server exited gracefully")
}

func AddName(n string) Option {
	return func(server *HTTPServer) {
		server.Name = n
	}
}

func AddPort(port int64) Option {
	return func(server *HTTPServer) {
		server.Port = port
	}
}

func AddMiddlewares(m []func(c *gin.Context)) Option {
	return func(server *HTTPServer) {
		server.Middlewares = append(server.Middlewares, m...)
	}
}

func (server *HTTPServer) AddRoutes(r []route.Route) {
	server.Routes = r
}

func (server *HTTPServer) AddGroupRoutes(gr []route.GroupRoute) {
	server.GroupRoutes = gr
}

func SetStrictSlash(strict bool) Option {
	return func(server *HTTPServer) {
		server.StrictSlash = strict
	}
}

func SetGracefulShutdownTimeout(t time.Duration) Option {
	return func(server *HTTPServer) {
		server.GracefulShutdownTimeout = t
	}
}

func Start(server *HTTPServer, err error) {
	if err != nil {
		panic(err)
	}
	server.Run()
}
