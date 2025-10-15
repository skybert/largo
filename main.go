package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"go.yaml.in/yaml/v3"
)

type TokenResponse struct {
	TraceId string `header: x-trace-id`
	Body    struct {
		AccessToken string `json:"access_token" example:"ey.." doc:"The keys to the mines of Moria"`
	}
}

type TokenRequest struct {
	Name string `header:"x-name"`
	Body struct {
		Name string `json:"name" required:"false"`
		Age  int    `json:"age" minimum:"18"`
		URN  string `json:"urn" pattern:"^skybert.net:foo:"`
	}
}

func handleToken(ctx context.Context, input *TokenRequest) (*TokenResponse, error) {
	resp := &TokenResponse{}
	resp.TraceId = "largo-" + strconv.Itoa(rand.Int())
	name := input.Name
	if name == "" {
		name = input.Body.Name
	}
	resp.Body.AccessToken = "hi " + name +
		", you may be " + fmt.Sprint(input.Body.Age) + " old" +
		" urn: " + input.Body.URN
	return resp, nil
}

func main() {
	router := http.NewServeMux()
	humaConf := huma.DefaultConfig("Largo", "42.0.0")
	humaConf.DocsPath = "/swagger-ui"

	api := humago.New(router, humaConf)
	huma.Post(api, "/token", handleToken)

	writeOpenAPISpecToFile(api)

	port := "8888"
	fmt.Printf("Starting Largo on port %v\n", port)

	http.ListenAndServe("127.0.0.1:"+port, router)
}

func writeOpenAPISpecToFile(api huma.API) {
	openAPISpec := api.OpenAPI()
	yamlBytes, err := yaml.Marshal(openAPISpec)
	if err != nil {
		panic(err)
	}

	os.WriteFile("openapi.yaml", yamlBytes, 0644)
}
