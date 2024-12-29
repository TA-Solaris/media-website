package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/TA-Solaris/common"
	pb "github.com/TA-Solaris/common/api"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	mediaServiceAddr = "localhost:2000" // TODO Change to not be hardcoded
)

func main() {
	conn, err := grpc.Dial(mediaServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials())) // TODO Use something not going to be deprecated
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Println("Dialing media service at", mediaServiceAddr)
	mediaServiceClient := pb.NewMediaServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(mediaServiceClient)
	handler.RegisterRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("Failed to start the HTTP server")
	}
}
