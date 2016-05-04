package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type HelloWorldService interface {
	HelloWorld(string) (string, error)
}

type helloWorldService struct{}

func (helloWorldService) HelloWorld(s string) (string, error) {
	if s == "" {
		return "Hello, World", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func main() {
	ctx := context.Background()
	svc := helloWorldService{}

	helloWorldHandler := httptransport.NewServer(
		ctx,
		makeHelloWorldEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	http.Handle("/helloWorld", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func makeHelloWorldEndpoint(svc HelloWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(helloWorldRequest)
		v, err := svc.HelloWorld(req.S)
		if err != nil {
			return helloWorldResponse{v, err.Error()}, nil
		}
		return helloWorldResponse{v, ""}, nil
	}
}

func decodeHelloWorldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request helloWorldRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type helloWorldRequest struct {
	S string `json:"s"`
}

type helloWorldResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}
