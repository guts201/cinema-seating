package server

import (
	"cinema/pkg/ent"
	"cinema/pkg/ent/migrate"
	mykit "cinema/pkg/genkit/pkg/api"
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/reflection"

	pb0 "cinema/api"
	redis2 "cinema/internal/redis"
	cinema "cinema/internal/server/clientcinema"
	bkcinema "cinema/internal/server/backofficecinema"
	config "cinema/pkg/config"

	dbe "cinema/pkg/database"
)

// Serve ...
func Serve(cfg *config.Config) {
	service := newService(cfg, []mykit.Option{}...)

	logger := service.Logger()
	server := service.Server()

	driver, err := dbe.Open("mysql_cinema", cfg.GetDatabase())
	if err != nil {
		logger.Fatal("can not open database", zap.Error(err))
	}
	entClient := ent.NewClient(ent.Driver(driver))
	defer func() {
		if err := entClient.Close(); err != nil {
			logger.Fatal("can not close ent client", zap.Error(err))
		}
	}()

	if err = entClient.Schema.Create(context.Background(), migrate.WithDropIndex(true)); err != nil {
		logger.Fatal("can not init my database", zap.Error(err))
	}

	redisClient := redis2.New(cfg, logger, cfg.GetRedis().GetEnabled())

	pb0.RegisterClientCinemaServer(server, cinema.NewServer(entClient, redisClient))
	pb0.RegisterBackofficeCinemaServer(server, bkcinema.NewServer(entClient))

	reflection.Register(server)

	service.Serve()
}
