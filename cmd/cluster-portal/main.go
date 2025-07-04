package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mdzahid786/cluster-portal/internal/config"
	"github.com/mdzahid786/cluster-portal/internal/http/handler/cluster"
	"github.com/mdzahid786/cluster-portal/internal/storage/mysql"
	"github.com/mdzahid786/cluster-portal/middleware"
)

func main() {
	// config setup
	cfg := config.MustLoad()
	users := cfg.Users
	// database setup
	storage, err := mysql.New(cfg)
	if err!=nil{
		log.Fatal(err)
	}
	// router setup
	router := http.NewServeMux()

	router.Handle("GET /api/login", middleware.AuthMiddleware(users, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := middleware.GetAuthenticatedUser(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		json.NewEncoder(w).Encode(user)
	})))
	router.Handle("POST /api/cluster", middleware.AuthMiddleware(users, http.HandlerFunc(cluster.New(storage))))
	//router.HandleFunc("POST /api/cluster", cluster.New(storage))
	//router.HandleFunc("GET /api/cluster/{id}", cluster.GetByID(storage))
	router.Handle("GET /api/cluster/{id}", middleware.AuthMiddleware(users, http.HandlerFunc(cluster.GetByID(storage))))
	//router.HandleFunc("GET /api/clusters/", cluster.GetClusters(storage))
	router.Handle("GET /api/clusters/", middleware.AuthMiddleware(users, http.HandlerFunc(cluster.GetClusters(storage))))
	
	adminHandler := middleware.AuthMiddleware(users,
    	middleware.AdminOnly(http.HandlerFunc(cluster.UpdateCluster(storage))),
	)
	router.Handle("PUT /api/cluster/{id}", adminHandler)
	//router.HandleFunc("PUT /api/cluster/{id}", cluster.UpdateCluster(storage))
	
	corsRouter := middleware.CORS(router)
	// server setup
	server := http.Server{
		Addr: cfg.HTTPServer.Addr,
		Handler: corsRouter,
	}
	slog.Info("Sever started", slog.String("Address", cfg.HTTPServer.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err.Error())
			log.Fatal("Failed to start server")
		}
	}()
	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}