package main

import (
	"os"
	"context"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"github.com/develamit/go-grpc-gateway/session-service/pkg/armstatus"

	pb "github.com/develamit/go-grpc-gateway/session-service/api/proto/v1"
)

const (
        gRPCport = ":4101"
        restport = ":4102"
)


type server struct{
	pb.UnimplementedSessionServer
}

func NewServer() *server {
	return &server{}
}
// Get
func (s *server) SessionGet(ctx context.Context, in *pb.SessionRequest) (*pb.SessionReply, error) {
	log.Info("[GET] Received: ", in.GetName())
	return &pb.SessionReply{Message: in.Name + " world (GET)"}, nil
}

func (s *server) SessionPost(ctx context.Context, in *pb.SessionRequest) (*pb.SessionReply, error) {
	log.Info("[POST] Received: ", in.GetName())
	return &pb.SessionReply{Message: in.Name + " world (POST)"}, nil
}

func main() {
	// Setup logrus
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	log.SetOutput(os.Stdout)

        file, err := os.OpenFile("logs/session_service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err == nil {
	    log.SetOutput(file)
        } else {
            log.Info("Failed to log to file, using default stderr")
        }

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", gRPCport)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Session service to the server
	pb.RegisterSessionServer(s, &server{})
	// Serve gRPC server
	log.Info("server listening at: ", lis.Addr())
	//log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		gRPCport,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal("Failed to dial server: ", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Session
	err = pb.RegisterSessionHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register gateway: ", err)
	}

	gwServer := &http.Server{
		Addr:    restport,
		Handler: gwmux,
	}

	log.Info("Serving gRPC-Gateway on http://0.0.0.0", restport)
	log.Fatal(gwServer.ListenAndServe())

	// write the status
	s := armstatus.AStatus{"running"}
	armstatus.WriteStatus(s)
}

