package provider

import (
	"context"

	"github.com/sobo-gen3v1/soboweb-tegola-cr/dict"
)

type MVTTiler interface {
	Layerer

	// MVTForLayers will return a MVT byte array or an error for the given layer names.
	MVTForLayers(ctx context.Context, tile Tile, params Params, layers []Layer) ([]byte, error)
}

// MVTInitFunc initialize a provider given a config map. The init function should validate the config map, and report any errors. This is called by the For function.
type MVTInitFunc func(dicter dict.Dicter, maps []Map) (MVTTiler, error)
