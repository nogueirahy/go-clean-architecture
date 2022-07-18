package main

import (
	"fmt"

	vtexapi "github.com/nogueirahy/go-clean-architecture/infra/vtex_api/product"
	presentation "github.com/nogueirahy/go-clean-architecture/presentation/presenters/product/product_search"
	service "github.com/nogueirahy/go-clean-architecture/service/usecases/product"
)

func main() {
	vtexApi := vtexapi.VtexApi{}
	productSearchService := service.ProductSearchUseCase{VtexApi: &vtexApi}
	productSearchController := presentation.ProductSearchPresenter{ProductSearchUseCase: &productSearchService}

	productEntity := productSearchController.Execute()

	fmt.Println(productEntity.ProductName)
	fmt.Println(productEntity.SomeBusinessLogic())
}
