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
	pricedParcel := &PricedParcel{
		Parcel: parcel,
	}
	if parcel.D1 < 10 && parcel.D2 < 10 && parcel.D3 < 10 {
		pricedParcel.Cost = 3.00
		return pricedParcel
	}
	if parcel.D1 < 50 && parcel.D2 < 50 && parcel.D3 < 50 {
		pricedParcel.Cost = 8.00
		return pricedParcel
	}
	if parcel.D1 < 100 && parcel.D2 < 100 && parcel.D3 < 100 {
		pricedParcel.Cost = 15.00
		return pricedParcel
	}

	pricedParcel.Cost = 25.00
	return pricedParcel
}
