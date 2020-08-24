# Local Plane

Plugin for running services locally with AIS Service Discovery library.

This is a series of mock adapters, which can be used with ais-service-discovery-go. Which allows you to communicate between services locally. Or mock them entirely.

## Use with mocked response
```golang
package main

import (
	"log"

	discovery "github.com/peak-ai/ais-service-discovery-go"
	"github.com/peak-ai/ais-service-discovery-go/pkg/types"
	"github.com/peak-ai/local-plane/backend"
	"github.com/peak-ai/local-plane/config"
)

func main() {

  // Configure a mock endpoint for your service to call
	config := config.Config{
		"test.service->queue": config.Service{
			Type: "queue",
			Resolve: config.Resolver{
				MockResponse: []byte("testing 123"),
			},
		},
	}

	d := discovery.NewDiscovery(backend.WithMockedBackend(config))

	id, err := d.Queue("test.service->queue", types.Request{})
	if err != nil {
		log.Panic(err)
	}

	log.Println("id: ", id)
}

```


## Use with a redirected endpoint

```golang
package main

import (
	"log"

	discovery "github.com/peak-ai/ais-service-discovery-go"
	"github.com/peak-ai/ais-service-discovery-go/pkg/types"
	"github.com/peak-ai/local-plane/backend"
	"github.com/peak-ai/local-plane/config"
)

func main() {

  // Configure a mock endpoint for your service to call
	config := config.Config{
		"test.service->queue": config.Service{
			Type: "queue",
			Resolve: config.Resolver{
				Type:         "http-post",
				Endpoint:     "http://localhost:8080",
			},
		},
	}

  // Calls http://localhost:8080 instead
  d := discovery.NewDiscovery(backend.WithRedirectedBackend(config))

  id, err := d.Queue("test.service->queue", types.Request{})
  if err != nil {
    log.Panic(err)
  }

  log.Println("Response: ", id)
}

```
