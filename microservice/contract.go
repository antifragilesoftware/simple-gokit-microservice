package microservice

import (
	"encoding/json"
	"net/http"
	"golang.org/x/net/context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func AddServices() {
	ctx := context.Background()
	svc := helloWorldService{}

	helloWorldHandler := httptransport.NewServer(
		ctx,
		makeHelloWorldEndpoint(svc),
		decodeHelloWorldRequest,
		encodeResponse,
	)

	http.Handle("/", helloWorldHandler)
}

func makeHelloWorldEndpoint(svc HelloWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.HelloWorld()
		if err != nil {
			return helloWorldResponse{v, err.Error()}, nil
		}
		return helloWorldResponse{v, ""}, nil
	}
}

func decodeHelloWorldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request helloWorldRequest
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
