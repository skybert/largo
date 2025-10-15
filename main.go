package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

type TokenResponse struct {
	TraceId string `header: x-trace-id`
	Body    struct {
		AccessToken string `json:"access_token" example:"ey.." doc:"The keys to the mines of Moria" validate:"required"`
	}
}

type TokenRequest struct {
	Name string `header:"x-name" required:false`
	Body struct {
		Age int `json:"age"`
	}
}

func handleToken(ctx context.Context, input *TokenRequest) (*TokenResponse, error) {
	resp := &TokenResponse{}
	resp.TraceId = "largo-" + strconv.Itoa(rand.Int())
	resp.Body.AccessToken = "hi " + input.Name + ", you may be " + fmt.Sprint(input.Body.Age) + " old"
	return resp, nil
}

func main() {
	router := http.NewServeMux()
	humaConf := huma.DefaultConfig("Largo", "42.0.0")
	humaConf.DocsPath = "/swagger-ui"

	api := humago.New(router, humaConf)
	huma.Post(api, "/token", handleToken)

	port := "8888"
	fmt.Printf("Starting Largo on port %v\n", port)

	http.ListenAndServe("127.0.0.1:"+port, router)
}
