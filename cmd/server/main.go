package main

import (
	"github.com/NikitaBarysh/metrics_and_alertinc/internal/config"
	"github.com/NikitaBarysh/metrics_and_alertinc/internal/handlers"
	"github.com/NikitaBarysh/metrics_and_alertinc/internal/router"
	"github.com/NikitaBarysh/metrics_and_alertinc/internal/storage/repositories"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	flags := config.NewFlagNames()
	flags.ParseFlags()

	memStorage := repositories.NewMemStorage()
	handler := handlers.NewHandler(memStorage)
	router := router.NewRouter(handler)
	chiRouter := chi.NewRouter()
	chiRouter.Mount("/", router.Register())
	err := http.ListenAndServe(flags.FlagRunAddr, chiRouter)
	if err != nil {
		panic(err)
	}
}
