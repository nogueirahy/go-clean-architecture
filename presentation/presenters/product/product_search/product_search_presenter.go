package presentation

import (
	"github.com/nogueirahy/go-clean-architecture/domain/entities"
	domain "github.com/nogueirahy/go-clean-architecture/domain/usecases/product"
)

type ProductSearchPresenter struct {
	ProductSearchUseCase domain.IProducSearchtUseCase
}

func (p ProductSearchPresenter) Execute() entities.ProductEntity {
	return p.ProductSearchUseCase.Execute("lech")
}
