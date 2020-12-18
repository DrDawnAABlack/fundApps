package parcel

import (
	"math"
)

const (
	NumParcelFields = 5

	smallParcelMaxDimension = 10
	mediumParcelMaxDimension = 50
	largeParcelMaxDimension = 100

	smallParcelCost      = 3.00
	mediumParcelCost     = 8.00
	largeParcelCost      = 15.00
	extraLargeParcelCost = 25.00

	smallParcelWeightLimit = float64(1)
	mediumParcelWeightLimit = float64(3)
	largeParcelWeightLimit = float64(6)
	extraLargeParcelWeightLimit = float64(10)
)

type Parcel struct {
	Id string
	D1 int64
	D2 int64
	D3 int64
	Weight float64
}

type PricedParcel struct {
	Parcel *Parcel
	Cost   float64
	Classification string
}

func NewParcel(id string, d1 int64, d2 int64, d3 int64, weight float64) *Parcel {
	return &Parcel{
		Id: id,
		D1: d1,
		D2: d2,
		D3: d3,
		Weight: weight,
	}
}

func (p Parcel) CostDueToSize() (float64, string) {
	if p.D1 < smallParcelMaxDimension && p.D2 < smallParcelMaxDimension && p.D3 < smallParcelMaxDimension {
		return smallParcelCost, "small"
	}
	if p.D1 < mediumParcelMaxDimension && p.D2 < mediumParcelMaxDimension && p.D3 < mediumParcelMaxDimension {
		return mediumParcelCost, "medium"
	}
	if p.D1 < largeParcelMaxDimension && p.D2 < largeParcelMaxDimension && p.D3 < largeParcelMaxDimension {
		return largeParcelCost, "large"
	}

	return extraLargeParcelCost, "extraLarge"
}

func (p PricedParcel) CostDueToWeight() float64 {
	overLimitWeight := float64(0)
	switch p.Classification {
	case "small":
			overLimitWeight = p.Parcel.Weight - smallParcelWeightLimit
	case "medium":
		overLimitWeight = p.Parcel.Weight - mediumParcelWeightLimit
	case "large":
		overLimitWeight = p.Parcel.Weight - largeParcelWeightLimit
	case "extraLarge":
		overLimitWeight = p.Parcel.Weight - extraLargeParcelWeightLimit
	//default: I've not included this as we set the classification so unlikely too be unknown
	}

	overWeightCharge := float64(0)
	if overLimitWeight > 0 {
		// round up because any amount over is over
		overWeightCharge = math.Ceil(overLimitWeight) * 2
	}

    return overWeightCharge
}