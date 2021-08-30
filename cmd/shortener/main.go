package main

import (
	"log"
	"net/http"

	"github.com/spinel/go-musthave-shortener-tpl/internal/app/repository"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/router"
)

func main() {
	repo, err := repository.New()
	if err != nil {
		panic(err)
	}

	go func() {
		http.Handle("/", router.Router(repo))
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
