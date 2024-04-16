package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/andreluizmicro/go-mocker-api/configs"
	"github.com/andreluizmicro/go-mocker-api/internal/application"
	"github.com/andreluizmicro/go-mocker-api/internal/infrastructure/repository"
	"github.com/andreluizmicro/go-mocker-api/internal/infrastructure/web/handlers"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

var sessionId = uuid.New().String()

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nEncerrando o programa...")
		removeFiles(sessionId)
		defer os.Exit(0)
	}()

	cfg, err := configs.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	webServerPort := cfg.WebServerPort
	if len(os.Args) >= 2 {
		webServerPort = os.Args[1]
	}

	aplicacao := Generate(webServerPort)
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Generate(port string) *cli.App {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "port",
				Value: "9000",
				Usage: "Porta da aplicação",
			},
			&cli.StringFlag{
				Name:  "uri",
				Value: "mock",
				Usage: "URI para realizar as requisições",
			},
		},
		Action: func(cCtx *cli.Context) error {
			port := cCtx.Args().Get(0)
			uri := cCtx.Args().Get(1)

			StartWebServer(port, uri)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	return app
}

func StartWebServer(port, uri string) {
	mux := http.NewServeMux()

	mockRepository := repository.NewMockRepository(sessionId)
	mockService := application.NewMockService(mockRepository)
	mockHandler := handlers.NewMockHandler(mockService)

	fmt.Println(port)

	mux.HandleFunc("POST /mock", mockHandler.Create)
	mux.HandleFunc(fmt.Sprintf("GET /%s/{session_id}", uri), mockHandler.Find)

	http.ListenAndServe(":"+port, mux)
}

func removeFiles(sessionId string) {
	os.RemoveAll(
		fmt.Sprintf("./tmp/mock/%s", sessionId),
	)
}
