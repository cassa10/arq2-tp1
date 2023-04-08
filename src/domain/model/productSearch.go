package model

type ProductSearchFilter struct {
	Name     string      `json:"name"`
	Category string      `json:"category"`
	Price    PriceFilter `json:"price"`
}

type PriceFilter struct {
	Value     float64 `json:"value"`
	IsStrict  bool    `json:"isStrict"`
	IsGreater bool    `json:"isGreater"`
}

func (f *ProductSearchFilter) GetPriceValue() float64 {
	return f.Price.Value
}

func (f *ProductSearchFilter) IsStrictPrice() bool {
	return f.Price.IsStrict
}

func (f *ProductSearchFilter) IsGreaterPrice() bool {
	return f.Price.IsGreater
}
