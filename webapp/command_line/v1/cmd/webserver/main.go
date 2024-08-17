package main

import (
	"context"
	"fmt"
	game "gowithtests/webapp/command_line/v1"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/rs/zerolog"
)

const (
	dbFileName  = "game.db.json"
	servicePort = 5000
)

func main() {

	errs := make(chan error)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store, err := game.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system store, %v ", err)
	}
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
