package authreq

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testConfig struct {
	Handler      string `json:"handler"`
	AuthEndpoint string `json:"auth_endpoint"`
}

var testNext caddyhttp.HandlerFunc = func(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func TestMiddleware(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}))
	defer ts.Close()

	tc := testConfig{
		Handler:      "authreq",
		AuthEndpoint: ts.URL,
	}

	payload, err := json.Marshal(tc)
	require.NoError(t, err)

	ctx, cancel := caddy.NewContext(caddy.Context{Context: context.Background()})
	defer cancel()

	loaded, err := ctx.LoadModuleInline("handler", "http.handlers", payload)
	require.NoError(t, err)

	mod, ok := loaded.(*Middleware)
	require.True(t, ok)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	assert.NoError(t, mod.ServeHTTP(rec, req, testNext))
}
