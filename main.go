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

type TempoRequest struct {
	Title string `header:"x-title"`
	Body  struct {
		Tempo int    `json:"tempo" required:"true" doc:"Tempo in bpm"`
		Title string `json:"title" required:"false"`
		URN   string `json:"urn" pattern:"^skybert.net:largo:"`
	}
}

type TempoResponse struct {
	TraceId string `header: x-trace-id`
	Body    struct {
		Title   string `json:"title,omitempty" required:"false"`
		Tempo   int    `json:"tempo" example:"40" doc:"How fast is largo music? (bpm)"`
		IsLargo bool   `json:"largo"`
	}
}

// type LargoInfoResponse struct {
// 	Body struct {
// 		TempoMin int `json:"tempo_min"`
// 		TempoMax int `json:"tempo_max"`
// 	}
// }

func handleTempo(ctx context.Context, req *TempoRequest) (*TempoResponse, error) {
	resp := &TempoResponse{}
	resp.TraceId = "largo-" + strconv.Itoa(rand.Int())
	title := req.Title
	if title == "" {
		title = req.Body.Title
	}
	resp.Body.Title = title
	resp.Body.IsLargo = (req.Body.Tempo >= 40 && req.Body.Tempo <= 66)
	resp.Body.Tempo = req.Body.Tempo
	return resp, nil
}

// func handleDefReq(ctx context.Context, req *struct{}) (*LargoInfoResponse, error) {
// 	resp := &LargoInfoResponse{}
// 	resp.Body.TempoMin = 40
// 	resp.Body.TempoMax = 66

// 	return resp, nil
// }

func main() {
	router := http.NewServeMux()
	humaConf := huma.DefaultConfig("Largo", "42.0.0")
	humaConf.DocsPath = "/swagger-ui"

	api := humago.New(router, humaConf)
	huma.Post(api, "/largo", handleTempo)
	// huma.Get(api, "/largo", handleDefReq)

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
