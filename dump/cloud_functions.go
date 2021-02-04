package dump

import (
	"context"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	bank_v1 "github.com/jmichiels/cloud-functions-tests/pkg/bank/protobuf/bank/v1"
	"google.golang.org/grpc"
	"net/http"
)

type server struct {
	bank_v1.UnimplementedBankServiceServer
}

func (s server) CreateClient(ctx context.Context, request *bank_v1.CreateClientRequest) (*bank_v1.CreateClientResponse, error) {
	fmt.Println("CreateClient")
	return &bank_v1.CreateClientResponse{}, nil
}

func (s server) GetClients(ctx context.Context, request *bank_v1.GetClientsRequest) (*bank_v1.GetClientsResponse, error) {
	fmt.Println("GetClients")
	return &bank_v1.GetClientsResponse{}, nil
}

//func main() {
//	// Create a listener on TCP port
//	lis, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		log.Fatalln("Failed to listen:", err)
//	}
//	// Create a gRPC server object
//	grpcServer := grpc.NewServer()
//	bakServiceServer := &server{}
//
//	// Attach the Greeter service to the server
//	bank_v1.RegisterBankServiceServer(grpcServer, bakServiceServer)
//
//	mux := runtime.NewServeMux()
//	if err := bank_v1.RegisterBankServiceHandlerServer(context.Background(), mux, bakServiceServer); err != nil {
//		log.Fatalln(err)
//	}
//
//	// Serve gRPC Server
//	log.Println("Serving gRPC on 0.0.0.0:8080")
//	log.Fatal(grpcServer.Serve(lis))
//}

var httpHandler http.Handler

// Runs once when the instance starts.
func init() {
	// Create a gRPC server object
	grpcServer := grpc.NewServer()
	bankServiceServer := &server{}
	// Attach the Greeter service to the server
	bank_v1.RegisterBankServiceServer(grpcServer, bankServiceServer)

	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	httpHandler = wrappedGrpc

	//conf := &firebase.Config{
	//	ProjectID: "cloud-functions-tests-45711",
	//}
	//// Initialize the Firebase client.
	//app, err := firebase.NewApp(context.Background(), conf)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//// Initialize the HTTP handler.
	//httpHandler, err = internal.NewHttpHandler(app)
	//if err != nil {
	//	log.Fatalln(err)
	//}
}

// Cloud function that acts as a single entry point to handle all HTTP requests.
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpHandler.ServeHTTP(w, r)
}
