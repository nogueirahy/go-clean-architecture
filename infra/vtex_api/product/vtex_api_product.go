package vtexapi

import (
	"github.com/nogueirahy/go-clean-architecture/infra/http"
	"github.com/nogueirahy/go-clean-architecture/infra/vtex_api/product/model"
)

type VtexApi struct {
	httpProvider *http.HttpAdapter
}

func (v *VtexApi) SearchProduct(params model.VtexApiSearchProductRequest) model.VtexApiSearchProductResponse {
	result := v.httpProvider.Get()
	response := model.VtexApiSearchProductResponse{Foo: result}
	return response
}
