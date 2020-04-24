package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"log"
	"net"
	"net/http"

	"github.com/cockroachdb/cmux"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	pb "nice/hello"
)

type helloServer struct {
}

func (s *helloServer) Echo(ctx context.Context, msg *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	log.Println(msg, ctx)
	return msg, nil
}

func grpcserver(l net.Listener) {
	print("hi grpc")
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &helloServer{})
	grpcServer.Serve(l)
}

func httpserver(l net.Listener) {
	print("hi httpserver")
	jsonOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{EmitDefaults: true})
	mux := runtime.NewServeMux(jsonOpt)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost:8080", opts...)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	err = pb.RegisterHelloServiceHandler(ctx, mux, conn)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	url_mux := http.NewServeMux()
	url_mux.Handle("/", http.Handler(mux))
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: url_mux,
	}

	srv.Serve(l)
}

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux.
	m := cmux.New(l)
	go func() { m.Serve() }()

	grpcL := m.Match(cmux.HTTP2())

	httpL := m.Match(cmux.HTTP1())

	go grpcserver(grpcL)
	go httpserver(httpL)

	select {}

}
