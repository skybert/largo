package main

import (
	"net/http"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/skybert/largo/oidc"
)

func main() {
	http.ListenAndServe("127.0.0.1:8888", oidc.CreateRouter())
}
