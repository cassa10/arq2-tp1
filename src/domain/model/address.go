package model

import "github.com/cassa10/arq2-tp1/src/domain/util"

type Address struct {
	Street      string `json:"street"`
	Number      int    `json:"int"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Observation string `json:"observation"`
}

func (a Address) String() string {
	return util.ParseStruct("Address", a)
}
