package model

import (
	"github.com/cassa10/arq2-tp1/src/domain/util"
)

type Product struct {
	Name        string
	Description string
}

func (p Product) String() string {
	return util.ParseStruct("Product", p)
}
