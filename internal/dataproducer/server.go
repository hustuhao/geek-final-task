package dataproducer

import (
	"github.com/google/wire"
)

// ProviderSet is geek-final-task providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer)
