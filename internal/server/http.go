package server

import (
	v1 "bibirt-api/api/user/v1"
	"bibirt-api/internal/conf"
	"bibirt-api/internal/service"

	"github.com/gorilla/handlers"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	cros_headers = []string{
		"access-control-allow-origin",
		"content-type",
		"x-md-global-appid",
		"x-md-golbal-appkey",
	}
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, auth *service.AuthService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders(cros_headers),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST"}),
		)),
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterAuthHTTPServer(srv, auth)
	return srv
}
