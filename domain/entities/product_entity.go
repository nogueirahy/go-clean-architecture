package entities

type ProductEntity struct {
	Id                   string  `json:"productId"`
	ProductName          string  `json:"productName"`
	PromotionalTag       string  `json:"promotionalTag"`
	ClusterTag           string  `json:"clusterTag"`
	PriceWithoutDiscount float64 `json:"priceWithoutDiscount"`
	Price                float64 `json:"price"`
	AvailableQuantity    int     `json:"availableQuantity"`
	IsAvailable          bool    `json:"isAvailable"`
	ImageUrl             string  `json:"imageUrl"`
	MeasurementUnit      string  `json:"measurementUnit"`
	UnitMultiplier       float64 `json:"unitMultiplier"`
}

func (p *ProductEntity) SomeBusinessLogic() bool {
	return p.Id == "1"
}
