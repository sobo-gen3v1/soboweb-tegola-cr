package observability

import (
	"github.com/sobo-gen3v1/soboweb-tegola-cr/dict"
	"github.com/sobo-gen3v1/soboweb-tegola-cr/internal/observer"
)

var NullObserver observer.Null

func noneInit(dict.Dicter) (Interface, error) { return NullObserver, nil }
