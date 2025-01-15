package main

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	config2 "github.com/sefikcan/news/internal/config"
	"github.com/sefikcan/news/internal/couchbase"
	"github.com/sefikcan/news/internal/elasticsearch"
	"github.com/sefikcan/news/internal/handlers"
	repository2 "github.com/sefikcan/news/internal/repository"
	"github.com/sefikcan/news/internal/routes"
	"github.com/sefikcan/news/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config := config2.LoadConfig("dev")

	couchbaseClient := couchbase.NewCouchbaseClient(config.Couchbase.Host, config.Couchbase.UserName, config.Couchbase.Password, config.Couchbase.Bucket)

	elasticClient := elasticsearch.NewElasticsearchClient(config.Elasticsearch.Url)

	repository := repository2.NewNewsRepository(couchbaseClient, elasticClient)

	newsService := service.NewNewsService(repository)

	newsHandler := handlers.NewNewsHandler(newsService)

	router := mux.NewRouter()
	routes.RegisterRoutes(router, newsHandler)

	server := &http.Server{
		Addr:         config.Server.Host + ":" + config.Server.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Server is starting at %s ...", config.Server.Host+":"+config.Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server couldn't be shut down properly: %v", err)
	}
	log.Println("Server closed.")
}
