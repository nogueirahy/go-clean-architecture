package vtexapi

import (
	"github.com/nogueirahy/go-clean-architecture/infra/vtex_api/product/model"
)

type IVtexApiProduct interface {
	SearchProduct(params model.VtexApiSearchProductRequest) model.VtexApiSearchProductResponse
}
