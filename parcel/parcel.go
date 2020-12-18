package parcel

import ()

const NumParcelFields = 4

type Parcel struct {
	Id string
	D1 int64
	D2 int64
	D3 int64
}

type PricedParcel struct {
	Parcel *Parcel
	Cost   float64
}

func New(id string, d1, d2, d3 int64) *Parcel {
	return &Parcel{
		Id: id,
		D1: d1,
		D2: d2,
		D3: d3,
	}
}

func CostDueToSize(parcel *Parcel) *PricedParcel {
	return &PricedParcel{}
}

