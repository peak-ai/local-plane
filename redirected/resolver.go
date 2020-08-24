package redirected

import (
	"github.com/peak-ai/ais-service-discovery-go/pkg/types"
	"github.com/peak-ai/local-plane/config"
)

type Resolver interface {
	Resolve(resolver config.Resolver) (*types.Response, error)
}
