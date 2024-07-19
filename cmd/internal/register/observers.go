package register

import (
	"github.com/sobo-gen3v1/soboweb-tegola-cr/dict"
	"github.com/sobo-gen3v1/soboweb-tegola-cr/internal/p"
	"github.com/sobo-gen3v1/soboweb-tegola-cr/observability"
)

func Observer(config dict.Dicter) (observability.Interface, error) {
	var oType = "none"
	if config != nil {
		oType, _ = config.String("type", p.String("none"))
	}
	return observability.For(oType, config)
}
