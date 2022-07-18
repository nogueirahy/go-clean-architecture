package service

import (
	"github.com/nogueirahy/go-clean-architecture/domain/entities"
	vtexapi "github.com/nogueirahy/go-clean-architecture/infra/vtex_api/product"
	"github.com/nogueirahy/go-clean-architecture/infra/vtex_api/product/model"
)

type ProductSearchUseCase struct {
	VtexApi vtexapi.IVtexApiProduct
}

func (p ProductSearchUseCase) Execute(textSearch string) entities.ProductEntity {
	payload := model.VtexApiSearchProductRequest{TextSearch: textSearch}
	responseData := p.VtexApi.SearchProduct(payload).Translator.ToEntity()

	return responseData
}
