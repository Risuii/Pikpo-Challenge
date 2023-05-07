package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/Risuii/Server/handler"
	"github.com/Risuii/Server/helpers/constant"
	"github.com/Risuii/Server/repository"
	"github.com/Risuii/Server/usecase"
	"github.com/Risuii/config"
	pb "github.com/Risuii/proto"
)

func main() {
	// connect to db
	cfg := config.New()

	// db := postgres.Open(cfg.Database.DSN)
	db, err := sql.Open("postgres", cfg.Database.DSN)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Connect to database")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		log.Fatalf("failed to liste: %v", err)
	}

	activityRepo := repository.NewActivityRepository(db, constant.TableActivity)
	activityUseCase := usecase.NewActivityUsecase(activityRepo)
	activityHandler := handler.NewActivityService(activityUseCase)

	grpcServer := grpc.NewServer()

	pb.RegisterActivityServiceServer(grpcServer, activityHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error when server grpc", err.Error())
	}

	log.Println("Server live !")
}
