package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/i-dontneedaname/study-todoapp/internal/core/logger"
	core_postgres_pool "github.com/i-dontneedaname/study-todoapp/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/i-dontneedaname/study-todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/i-dontneedaname/study-todoapp/internal/core/transport/http/server"
	users_postgres_repository "github.com/i-dontneedaname/study-todoapp/internal/features/users/repository/postgres"
	users_service "github.com/i-dontneedaname/study-todoapp/internal/features/users/service"
	users_transport_http "github.com/i-dontneedaname/study-todoapp/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Initializing postgres connection pool...")
	connectionPool, err := core_postgres_pool.NewConnectionPool(ctx, core_postgres_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer connectionPool.Close()

	logger.Debug("Initializing feature...", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(connectionPool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("Initializing HTTP server...")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	httpServer.RegisterApiRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}
