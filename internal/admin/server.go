package admin

import (
	"github.com/google/wire"
)

// ProviderSet is admin providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer)
