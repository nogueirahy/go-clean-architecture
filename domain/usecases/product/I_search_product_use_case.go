package usecases

import "github.com/nogueirahy/go-clean-architecture/domain/entities"

type ISearchProductUseCase interface {
	Execute(textSearch string) (entities.ProductEntity, error)
}
