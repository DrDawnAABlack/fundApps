package parcel

import (
    "math"
)

type Order struct {
    PricedParcels map[string]*PricedParcel
    CostOfShipping float64
    CostOfSpeedyShipping float64
}

type DiscountParcelBatch struct {

}

const (
    numberOfPackagesInSmallPackageDeal = 3
)

func NewOrder() *Order {
    return &Order{
        PricedParcels: map[string]*PricedParcel{},
    }
}

func (r Order) applyDiscount() float64 {
    //discount := float64(0)

    //discountAfterSmallMania, pricedParcelsAfterSmallMania := smallMania(r.PricedParcels)
    //discount = discount + discountAfterSmallMania


    return 0
}

func smallMania(pricedParcels map[string]*PricedParcel) (float64, map[string]*PricedParcel) {
    cheapestParcelCost := []float64{}
    var cheapestParcelId []string

    for i:=0; i<numberOfPackagesInSmallPackageDeal; i++ {
        cheapestParcelCost[i] = math.MaxFloat64
    }

    for _, pp := range pricedParcels {
        if pp.Classification == "small" {
            cheapestParcelCost, cheapestParcelId = addParcelToDiscountCandidate(pp, cheapestParcelCost, cheapestParcelId )
        }
    }

    // remove discounted packages from the map
    for _, id := range cheapestParcelId {
        delete(pricedParcels, id)
    }

    return cheapestParcelCost[0], nil
}

func addParcelToDiscountCandidate(pricedParcel *PricedParcel, cheapestCosts []float64, ids []string) ([]float64, []string) {

    if pricedParcel.Cost <= cheapestCosts[0] {
        cheapestCosts[2] = cheapestCosts[1]
        ids[2] = ids[1]
        cheapestCosts[1] = cheapestCosts[0]
        ids[1] = ids[0]
        cheapestCosts[0] = pricedParcel.Cost
        ids[0] = pricedParcel.Parcel.Id
    }else if pricedParcel.Cost <= cheapestCosts[1] {
        cheapestCosts[2] = cheapestCosts[1]
        ids[2] = ids[1]
        cheapestCosts[1] = pricedParcel.Cost
        ids[1] = pricedParcel.Parcel.Id
    } else if pricedParcel.Cost <= cheapestCosts[2] {
        cheapestCosts[2] = pricedParcel.Cost
        ids[2] = pricedParcel.Parcel.Id
    }

    return cheapestCosts, ids
}