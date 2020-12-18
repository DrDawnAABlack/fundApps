package parcel

import ()

const (
	NumParcelFields = 4

	smallParcelMaxDimension = 10
	mediumParcelMaxDimension = 50
	largeParcelMaxDimension = 100

	smallParcelCost      = 3.00
	mediumParcelCost     = 8.00
	largeParcelCost      = 15.00
	extraLargeParcelCost = 25.00

	smallParcelWeightLimit = 1
	mediumParcelWeightLimit = 3
	largeParcelWeightLimit = 6
	extraLargeParcelWeightLimit = 10
)

type Parcel struct {
	Id string
	D1 int64
	D2 int64
	D3 int64
	Weight int64
}

type PricedParcel struct {
	Parcel *Parcel
	Cost   float64
	Classification string
}

func NewParcel(id string, d1, d2, d3 int64) *Parcel {
	return &Parcel{
		Id: id,
		D1: d1,
		D2: d2,
		D3: d3,
	}
}

func CostDueToSize(parcel *Parcel) (float64, string) {
	if parcel.D1 < smallParcelMaxDimension && parcel.D2 < smallParcelMaxDimension && parcel.D3 < smallParcelMaxDimension {
		return smallParcelCost, "small"
	}
	if parcel.D1 < mediumParcelMaxDimension && parcel.D2 < mediumParcelMaxDimension && parcel.D3 < mediumParcelMaxDimension {
		return mediumParcelCost, "medium"
	}
	if parcel.D1 < largeParcelMaxDimension && parcel.D2 < largeParcelMaxDimension && parcel.D3 < largeParcelMaxDimension {
		return largeParcelCost, "large"
	}

	return extraLargeParcelCost, "extraLarge"
}

func CostDueToWeight(parcel *Parcel) float64 {

    return 0
}