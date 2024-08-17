package main

import (
	"context"
	"fmt"
	game "gowithtests/webapp/command_line/v2"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/rs/zerolog"
)

const (
	dbFileName  = "/Users/aks/mine/go/gowithtests/webapp/command_line/v2/game.db.json"
	servicePort = 8888
)

func main() {

	errs := make(chan error)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	store, close, err := game.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		logger.Fatal().Err(err)
	}
	defer close()

	server := game.NewPlayerServer(store)
	httpServer := http.Server{
		Addr:    ":" + strconv.Itoa(servicePort),
		Handler: server,
	}

	// start the http server in another go routine
	go func() {
		logger.Info().Msg(fmt.Sprintf("Application HTTP server is now starting on port %d", servicePort))
		errs <- httpServer.ListenAndServe()
	}()

	// wait for os signal in another go routine to stop it gracefully
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		logger.Info().Msg(fmt.Sprintf("caught %s application will attempt to gracefully shutdown", sig.String()))
		errs <- fmt.Errorf("%s", sig)
	}()

	// wait on error channel for signal or server startup failed
	err = <-errs
	logger.Error().AnErr("Recieved error on error chan", err)
	if errHttpServer := httpServer.Shutdown(context.Background()); errHttpServer != nil {
		logger.Error().AnErr("failed to gracefully shutdown http server", errHttpServer)
	}
	logger.Info().Msg("server shutdown complete, application will exit now")
}
