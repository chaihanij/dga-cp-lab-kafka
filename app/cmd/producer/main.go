package main

import (
	"context"
	"dga-cp-lab-kafka/app/database"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dga-cp-lab-kafka/app/middlewares"
	"github.com/gin-gonic/gin"

	_notifyHttpHandler "dga-cp-lab-kafka/app/layers/deliveries/http/notify"
	//
	_notifyRepo "dga-cp-lab-kafka/app/layers/repositories/notify"
	_notifyUseCase "dga-cp-lab-kafka/app/layers/usecases/notify"
)

func main() {
	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(gin.Logger())
	ginEngine.Use(middlewares.CORSMiddleware())

	p := database.KafkaInitProducer()
	notifyRepo := _notifyRepo.InitRepo(p)
	notifyUseCase := _notifyUseCase.InitUseCase(notifyRepo)
	_notifyHttpHandler.NewEndpointHTTPHandler(ginEngine, notifyUseCase)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: ginEngine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Flush and close the producer and the events channel
	for p.Flush(10000) > 0 {
		log.Print("Still waiting to flush outstanding messages\n")
	}
	p.Close()

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

}
