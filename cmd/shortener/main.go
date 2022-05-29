package main

import (
	"fmt"
	"log"
	"net"

	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/pkg"
	"github.com/spinel/go-musthave-shortener/internal/app/pkg/pb"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
	"github.com/spinel/go-musthave-shortener/internal/app/service"
	"google.golang.org/grpc"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

//syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT
func main() {
	fmt.Printf("Build version: %s\n", pkg.CheckNA(buildVersion))
	fmt.Printf("Build date: %s\n", pkg.CheckNA(buildDate))
	fmt.Printf("Build commit: %s\n\n", pkg.CheckNA(buildCommit))

	cfg := config.NewConfig()
	if err := cfg.Validate(); err != nil {
		log.Fatal("config validation failed: ", err)
	}
	if cfg.Config != "" {
		cfg, _ = pkg.ReadJsonFile(cfg.Config)
	}

	repo, err := repository.NewStorage(cfg)
	if err != nil {
		log.Fatal("storage init failed:", err)
	}

	s := service.Server{
		Repo: repo.EntityPg,
	}

	lis, err := net.Listen("tcp", ":50060")
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", ":50060")

	grpcServer := grpc.NewServer()

	pb.RegisterStorageServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
