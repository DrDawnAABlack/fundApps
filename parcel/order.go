package parcel


type Order struct {
    PricedParcels map[string]*PricedParcel
    CostOfShipping float64
    CostOfSpeedyShipping float64
}

func NewOrder() *Order {
    return &Order{
        PricedParcels: map[string]*PricedParcel{},
    }
}