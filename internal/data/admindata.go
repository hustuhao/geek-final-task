package data

import (
	"github.com/google/wire"
)

var AdminProviderSet = wire.NewSet(NewAdminData, NewOrderRepo, NewEsClient, NewDiscovery, NewRegistry)
