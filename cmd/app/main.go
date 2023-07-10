package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alexander272/mattermost_bot/internal/config"
	"github.com/Alexander272/mattermost_bot/internal/server"
	"github.com/Alexander272/mattermost_bot/internal/services"
	transport "github.com/Alexander272/mattermost_bot/internal/transport/http"
	"github.com/Alexander272/mattermost_bot/pkg/logger"
	"github.com/Alexander272/mattermost_bot/pkg/mattermost"
)

func main() {
	// if err := gotenv.Load(".env"); err != nil {
	// 	logger.Fatalf("error loading env variables: %s", err.Error())
	// }

	conf, err := config.Init("configs/config.yaml")
	if err != nil {
		logger.Fatalf("error initializing configs: %s", err.Error())
	}
	logger.Init(os.Stdout, conf.Environment)

	//* Dependencies

	mostClient := mattermost.NewMattermostClient(mattermost.Config{Server: conf.Most.Server, Token: conf.Most.Token})

	//* Services, Repos & API Handlers
	services := services.NewServices(services.Deps{
		MostClient: mostClient,
		ChannelId:  conf.Most.ChannelId,
	})
	handlers := transport.NewHandler(services)

	//* HTTP Server
	srv := server.NewServer(conf, handlers.Init(conf))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	logger.Infof("Application started on port: %s", conf.Http.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
