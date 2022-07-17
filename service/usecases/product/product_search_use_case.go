package service

import usecases "github.com/nogueirahy/go-clean-architecture/domain/usecases/product"

type ProductSearchUseCase struct {
	searchProduct usecases.ISearchProductUseCase
}

func (p *ProductSearchUseCase) execute(textSearch string) {
	p.searchProduct.Execute(textSearch)
}
