package main

import (
	"log"

	"github.com/InstaySystem/is_v2-be/internal/container"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/api"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/config"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/worker"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	ctn, err := container.NewContainer(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	sv := api.NewServer(cfg, ctn)

	mqWorker := worker.NewMessageQueueWorker(ctn.MQPro, ctn.SMTPPro, ctn.Log)
	mqWorker.Start()

	ch := make(chan error, 1)
	go func() {
		if err := sv.Start(); err != nil {
			ch <- err
		}
	}()

	log.Printf("Server is running at: http://localhost:%d", cfg.Server.Port)

	sv.GracefulShutdown(ch)

	ctn.Cleanup()
}
