package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/assert"
)

func TestGetLargoDef(t *testing.T) {
	_, api := humatest.New(t)

	addRoutes(api)

	resp := api.Get("/largo")
	if !strings.Contains(resp.Body.String(), "Music with slow tempo") {
		t.Fatalf("Unexpected response: %s", resp.Body.String())
	}
}

func TestPostLargo(t *testing.T) {
	_, api := humatest.New(t)

	addRoutes(api)

	resp := api.Post("/largo", map[string]any{
		"tempo": 40,
	})

	assert.Equal(t, 200, resp.Code)

	var result *struct {
		Title   string `json:"title,omitempty" required:"false"`
		Tempo   int    `json:"tempo" example:"40" doc:"How fast is largo music? (bpm)"`
		IsLargo bool   `json:"largo"`
	}

	err := json.Unmarshal(resp.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Equal(t, true, result.IsLargo)
}

func TestPostLargoInvalidURN(t *testing.T) {
	_, api := humatest.New(t)

	addRoutes(api)

	resp := api.Post("/largo", map[string]any{
		"urn": "this:is:invalid",
	})

	if resp.Code != 422 {
		t.Fatalf("Unexpected status code: %v", resp)
	}
}
