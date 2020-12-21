package parcel

import (
	"math"
)

const (
	NumParcelFields = 5

	// dimensions are in cm
	smallParcelMaxDimension = 10
	mediumParcelMaxDimension = 50
	largeParcelMaxDimension = 100

	smallParcelCost      = 3.00
	mediumParcelCost     = 8.00
	largeParcelCost      = 15.00
	extraLargeParcelCost = 25.00
	heavyParcelCost = 50.00

	// weights are in Kg
	smallParcelWeightLimit = float64(1)
	mediumParcelWeightLimit = float64(3)
	largeParcelWeightLimit = float64(6)
	extraLargeParcelWeightLimit = float64(10)

	heavyParcelWeightLimit = 50
)

type Parcel struct {
	Id     string
	d1     int64
	d2     int64
	d3     int64
	weight float64
}

type PricedParcel struct {
	Parcel *Parcel
	Cost   float64
	Classification string
}

func NewParcel(id string, d1 int64, d2 int64, d3 int64, weight float64) *Parcel {
	return &Parcel{
		Id:     id,
		d1:     d1,
		d2:     d2,
		d3:     d3,
		weight: weight,
	}
}

func (p Parcel) CostDueToSize() (float64, string) {
	if p.d1 < smallParcelMaxDimension && p.d2 < smallParcelMaxDimension && p.d3 < smallParcelMaxDimension {
		return smallParcelCost, "small"
	}
	if p.d1 < mediumParcelMaxDimension && p.d2 < mediumParcelMaxDimension && p.d3 < mediumParcelMaxDimension {
		return mediumParcelCost, "medium"
	}
	if p.d1 < largeParcelMaxDimension && p.d2 < largeParcelMaxDimension && p.d3 < largeParcelMaxDimension {
		return largeParcelCost, "large"
	}
	return extraLargeParcelCost, "extraLarge"
}

func (p Parcel) IsExtraHeavy() bool {
    if p.weight >= heavyParcelWeightLimit {
        return true
    }
    return false
}

func (p PricedParcel) CostDueToWeight() float64 {
    // round up because any amount over is over
    overWeightCharge := float64(0)
	switch p.Classification {
	case "small":
        overWeightCharge = math.Ceil(p.Parcel.weight - smallParcelWeightLimit) * 2
	case "medium":
        overWeightCharge = math.Ceil(p.Parcel.weight - mediumParcelWeightLimit) * 2
	case "large":
        overWeightCharge = math.Ceil(p.Parcel.weight - largeParcelWeightLimit) * 2
	case "extraLarge":
        overWeightCharge = math.Ceil(p.Parcel.weight - extraLargeParcelWeightLimit) * 2
    case "extraHeavy":
        overWeightCharge = heavyParcelCost + math.Ceil(p.Parcel.weight - heavyParcelWeightLimit)

	//default: I've not included this as we set the classification so unlikely too be unknown
	}

	if overWeightCharge > 0 {
		return overWeightCharge
	}

    return 0
}
