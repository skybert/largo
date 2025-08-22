package main

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

type TokenResponse struct {
	TraceId string `header:x-trace-id`
	Body    struct {
		AccessToken string `json:"access_token" example:"ey.." doc:"The keys to the mines of Moria"`
	}
}

func handleToken(ctx context.Context, input *struct {
	Name string `header:"x-name" required:false`
}) (*TokenResponse, error) {
	resp := &TokenResponse{}
	resp.TraceId = "largo-" + strconv.Itoa(rand.Int())
	resp.Body.AccessToken = "your access token, Sir"
	return resp, nil
}

func main() {
	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))
	// TODO: Register operations...

	// Register GET /greeting/{name} handler.
	huma.Post(api, "/token", handleToken)

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}
