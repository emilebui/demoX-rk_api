package cmd

import (
	"context"
	"fmt"
	"github.com/emilebui/demoX-rk_api/internal/handlers"
	"github.com/emilebui/demoX-rk_api/internal/repositories"
	"github.com/emilebui/demoX-rk_api/internal/services"
	"github.com/emilebui/demoX-rk_api/pkg/conn"
	"github.com/google/uuid"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkecho "github.com/rookie-ninja/rk-echo/boot"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
)

func Run() {
	boot := rkboot.NewBoot()

	appConf := rkentry.GlobalAppCtx.GetConfigEntry("my-config")
	logger := rkentry.GlobalAppCtx.GetLoggerEntry("app-logger")
	echoEntry := rkecho.GetEchoEntry("demoX-rk_api")

	// Init api instance
	instanceID := uuid.New()

	// Add shutdown hook function
	boot.AddShutdownHookFunc("shutdown-hook", func() {
		fmt.Println("Shutting down ...")
	})
	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Register handlers, repo, services, etc, ...

	// Register Repo
	repo := registerRepo(appConf, logger)

	// Register services
	service := services.NewService(repo, logger)

	// Register Handlers
	demoHandler := handlers.NewDemoHandler()
	msgHandler := handlers.NewMessageHandler(service, instanceID.String())

	// Register API
	demoHandler.RegisterAPI(echoEntry)
	msgHandler.RegisterAPI(echoEntry)

	boot.WaitForShutdownSig(context.TODO())
}

func registerRepo(appConf *rkentry.ConfigEntry, logger *rkentry.LoggerEntry) repositories.Repo {

	kafkaMode := appConf.GetBool("kafka_mode")
	var repo repositories.Repo

	if kafkaMode {
		logger.Info("Initializing repository with Kafka ...")
		p := conn.CreateProducer(appConf)
		repo = repositories.NewKafkaRepo(p, appConf.GetString("kafka_topic"))

	} else {
		logger.Info("Initializing repository with Redis ...")
		redisClient := conn.GetRedisConn(appConf)
		repo = repositories.NewRedisRepo(redisClient, appConf.GetString("redis_queue_key"))
	}

	return repo
}
