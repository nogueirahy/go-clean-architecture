package domain

import "github.com/nogueirahy/go-clean-architecture/domain/entities"

type IProducSearchtUseCase interface {
	Execute(textSearch string) entities.ProductEntity
}
