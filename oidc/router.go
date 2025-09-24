package oidc

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
	TraceId string `header:"x-trace-id" doc:"Trace identifier for debugging"`
	Body    struct {
		AccessToken string `json:"access_token" example:"ey.." doc:"The keys to the mines of Moria"`
	}
}

type TokenRequest struct {
	Name string `header:"x-name" doc:"the name"`
}

func handleToken(ctx context.Context, input *TokenRequest) (*TokenResponse, error) {
	resp := &TokenResponse{}
	resp.TraceId = "largo-" + strconv.Itoa(rand.Int())
	resp.Body.AccessToken = "your access token, Sir"
	return resp, nil
}

func CreateRouter() *http.ServeMux {
	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))
	huma.Post(api, "/token", handleToken)
	return router
}
