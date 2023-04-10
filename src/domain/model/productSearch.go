package model

import (
	"github.com/cassa10/arq2-tp1/src/domain/util"
)

type ProductSearchFilter struct {
	Name     string   `json:"name"`
	Category string   `json:"category"`
	PriceMin *float64 `json:"priceMin"`
	PriceMax *float64 `json:"priceMax"`
}

func NewProductSearchFilter(name, category string, priceMin, priceMax *float64) ProductSearchFilter {
	return ProductSearchFilter{
		Name:     name,
		Category: category,
		PriceMin: priceMin,
		PriceMax: priceMax,
	}
}

func (f *ProductSearchFilter) ContainsAnyPriceFilter() bool {
	return f.PriceMin != nil || f.PriceMax != nil
}

func (f *ProductSearchFilter) GetPriceMinOrDefault() float64 {
	if f.PriceMin == nil {
		return 0
	}
	return *f.PriceMin
}

func (f *ProductSearchFilter) GetPriceMaxOrDefault() float64 {
	if f.PriceMax == nil {
		return 999999999999999999
	}
	return *f.PriceMax
}

func (f *ProductSearchFilter) String() string {
	return util.ParseStruct("ProductSearchFilter", f)
}
