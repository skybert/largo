package oidc

import (
	"context"
	"strings"
	"testing"
)

func TestHandleToken(t *testing.T) {
	input := &TokenRequest{
		Name: "TestUser",
	}

	ctx := context.Background()
	resp, err := handleToken(ctx, input)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !strings.HasPrefix(resp.TraceId, "largo-") {
		t.Errorf("Expected TraceId to start with 'largo-', got: %s", resp.TraceId)
	}

	if resp.Body.AccessToken != "your access token, Sir" {
		t.Errorf("Unexpected access token: %s", resp.Body.AccessToken)
	}
}
