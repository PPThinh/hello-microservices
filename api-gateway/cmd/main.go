package main

import (
	hellopb "api-gateway/proto/hello"
	userpb "api-gateway/proto/user"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net/http"
	"strconv"
)

func main() {
	helloConn, err := grpc.Dial("[::]:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial hello: %v", err)
	}
	defer helloConn.Close()
	helloClient := hellopb.NewHelloServiceClient(helloConn)

	userConn, err := grpc.Dial("[::]:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial user: %v", err)
	}
	defer userConn.Close()
	userClient := userpb.NewUserServiceClient(userConn)

	http.HandleFunc("/hello-user", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		helloResp, err := helloClient.GetHello(context.Background(), &emptypb.Empty{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userReq := &userpb.UserRequest{Id: uint32(id)}
		userResp, err := userClient.GetUser(context.Background(), userReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("%s %s", helloResp.Message, userResp.Name)
		fmt.Fprintln(w, message)
	})

	log.Println("API Gateway listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
