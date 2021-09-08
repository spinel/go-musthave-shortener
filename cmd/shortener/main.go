package main

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spinel/go-musthave-shortener/internal/app/repository/web"
	"github.com/spinel/go-musthave-shortener/internal/app/router"
	"github.com/spinel/go-musthave-shortener/internal/app/store"
)

func main() {
	// gob db file
	gobFileName := "urls.gob"

	// gob writer
	var gobWrite *bytes.Buffer

	//goob reader
	gobRead, _ := os.Open(gobFileName)

	// gob storage
	s := store.NewStore(gobFileName, gobWrite, gobRead)

	// memory storage
	memory, _ := s.GetData()

	// Entity interface
	entityRepo := web.NewEntityRepo(memory)
	defer gobRead.Close()

	server := &http.Server{Addr: ":8080", Handler: router.NewRouter(entityRepo)}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				memory := entityRepo.GetMemory()
				err := s.SaveData(memory)
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)

	// Set os signals to the chan
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

	ticker.Stop()
	done <- true
}
