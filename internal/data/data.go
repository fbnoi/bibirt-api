package data

import (
	"bibirt-api/internal/conf"
	"bibirt-api/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
