package authreq

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	if err := caddy.RegisterModule(Middleware{}); err != nil {
		log.Fatal(fmt.Errorf("authreq: failed to register module: %w", err))
	}
	httpcaddyfile.RegisterHandlerDirective("authreq", parseCaddyfile)
}

type Middleware struct {
	AuthEndpoint string `json:"auth_endpoint,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		Name: "http.handlers.authreq",
		New:  func() caddy.Module { return new(Middleware) },
	}
}

// Provision implements caddy.Provisioner.
func (m *Middleware) Provision(ctx caddy.Context) error {
	return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	authReq, err := http.NewRequest(http.MethodGet, m.AuthEndpoint, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("authreq: failed to create request: %w", err)
	}

	authReq.Header = r.Header

	authResp, err := http.DefaultClient.Do(authReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("authreq: failed to call auth request: %w", err)
	}

	if authResp.StatusCode == http.StatusAccepted {
		return next.ServeHTTP(w, r)
	}

	w.WriteHeader(authResp.StatusCode)
	return fmt.Errorf("authreq: unexpected status: %s", authResp.Status)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.Args(&m.AuthEndpoint) {
			return fmt.Errorf("authreq: failed to unmarshal: %w", d.ArgErr())
		}
	}
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	if err := m.UnmarshalCaddyfile(h.Dispenser); err != nil {
		return nil, fmt.Errorf("authreq: failed to parse Caddy file: %w", err)
	}
	return m, nil
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)
