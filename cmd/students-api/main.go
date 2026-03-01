package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Prohor722/go_rest_api_1/internal/config"
	"github.com/Prohor722/go_rest_api_1/internal/http/handlers/student"
	"github.com/Prohor722/go_rest_api_1/internal/storage/sqlite"
)

// import "fmt"

func main() {
	// load config
	cfg := config.MustLoad()


	// database setup
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initilized", slog.String("env", cfg.Env))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students/", student.GetStudents(storage))
	router.HandleFunc("PUT /api/students/{id}", student.UpdateStudent(storage))
	
	//Delete
	router.HandleFunc("DELETE /api/students/{id}", student.DeleteStudent(storage))
	


	// setup server
	server := http.Server {
		Addr: cfg.Address,
		Handler: router,
	}


	// fmt.Println("Server Started!!")
	// fmt.Printf("Serving on: %s",cfg.Address)
	slog.Info("Server started!",slog.String("Address",cfg.Address))

	err2 := server.ListenAndServe()

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func(){
		if err2 != nil {
			log.Fatal("Fail to start server !")
		}
	}()

	<-done

	slog.Info("shutting down the server..")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer cancel()

	e := server.Shutdown(ctx)

	if e != nil {
		slog.Error("Failed to shutdown server", slog.String("error:",e.Error()))
	}

	slog.Info("Server shutdown successfully !")

}