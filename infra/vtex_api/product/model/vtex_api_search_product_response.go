package model

import "github.com/nogueirahy/go-clean-architecture/domain/entities"

type Translator struct{}

func (t *Translator) ToEntity() entities.ProductEntity {
	data := entities.ProductEntity{Id: "1", ProductName: "Leche"}
	return data
}

type VtexApiSearchProductResponse struct {
	Translator *Translator
	Foo        float64
}
