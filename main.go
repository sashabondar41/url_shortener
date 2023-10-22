package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"urlshortener/internal/in_memory"
	"urlshortener/internal/repository"
	"urlshortener/internal/server"
	"urlshortener/pb"
)

const (
	repo = "repo"
	mem  = "mem"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
	stg, err := chooseStorage()
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterUrlShortenerServer(srv, server.New(stg))
	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	/*srv := server.New(stg)
	err = srv.Start(":8000")
	if err != nil {
		log.Fatalln(err.Error())
	}*/
}

func chooseStorage() (server.StorageInterface, error) {
	var storageType string
	flag.StringVar(&storageType, "type", repo, "Declaration of storage type")
	flag.Parse()
	fmt.Println("storage type: ", storageType)
	if storageType == repo {
		db, err := loadDatabase()
		if err != nil {
			return nil, err
		}
		return repository.New(db), nil
	} else if storageType == mem {
		return in_memory.New(), nil
	} else {
		return nil, fmt.Errorf("incorrect parameter \"type\"")
	}
}

func loadDatabase() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		host, username, password, databaseName, port)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("database opening failed: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}
	return db, nil
}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("loading env file failed: %w", err)
	}
	return nil
}
