package parcel

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestParcelPricingBySize(t *testing.T) {

    type testCase struct {
        name         string
        parcel       Parcel
        expectedCost float64
        expectedClassification string
    }

    var testCases = []testCase{
        {
            name: "all dimensions small",
            parcel: Parcel{
                d1: 2,
                d2: 3,
                d3: 4,
            },
            expectedCost: 3.00,
            expectedClassification: "small",
        },
        {
            name: "all dimensions medium",
            parcel: Parcel{
                d1: 25,
                d2: 36,
                d3: 42,
            },
            expectedCost: 8.00,
            expectedClassification: "medium",
        },
        {
            name: "all dimensions large",
            parcel: Parcel{
                d1: 66,
                d2: 77,
                d3: 88,
            },
            expectedCost: 15.00,
            expectedClassification: "large",
        },
        {
            name: "all dimensions extra large",
            parcel: Parcel{
                d1: 100,
                d2: 100,
                d3: 100,
            },
            expectedCost: 25.00,
            expectedClassification: "extraLarge",
        },
        {
            name: "mixed dimensions medium",
            parcel: Parcel{
                d1: 1,
                d2: 1,
                d3: 42,
            },
            expectedCost: 8.00,
            expectedClassification: "medium",
        },
        {
            name: "mixed dimensions large",
            parcel: Parcel{
                d1: 1,
                d2: 1,
                d3: 88,
            },
            expectedCost: 15.00,
            expectedClassification: "large",
        },
        {
            name: "mixed dimensions extra large",
            parcel: Parcel{
                d1: 1,
                d2: 1,
                d3: 100,
            },
            expectedCost: 25.00,
            expectedClassification: "extraLarge",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            cost, classification := tc.parcel.CostDueToSize()
            assert.Equal(t, tc.expectedCost, cost)
            assert.Equal(t, tc.expectedClassification, classification)
        })
    }
}


func TestParcelPricingByWeight(t *testing.T) {

    type testCase struct {
        name                    string
        pricedParcel            PricedParcel
        expectedCostDueToWeight float64
    }

    var testCases = []testCase{
        {
            name: "small not overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    d1:     2,
                    d2:     3,
                    d3:     4,
                    weight: smallParcelWeightLimit,
                },
                Classification: "small",
            },
            expectedCostDueToWeight: 0,
        },
        {
            name: "medium not overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    d1:     30,
                    d2:     30,
                    d3:     30,
                    weight: mediumParcelWeightLimit,
                },
                Classification: "medium",
            },
            expectedCostDueToWeight: 0,
        },
        {
            name: "large not overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    d1:     90,
                    d2:     90,
                    d3:     90,
                    weight: largeParcelWeightLimit,
                },
                Classification: "large",
            },
            expectedCostDueToWeight: 0,
        },
        {
            name: "extra large not overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    d1:     200,
                    d2:     300,
                    d3:     400,
                    weight: extraLargeParcelWeightLimit,
                },
                Classification: "extraLarge",
            },
            expectedCostDueToWeight: 0,
        },
        {
            name: "small overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    d1:     2,
                    d2:     3,
                    d3:     4,
                    weight: 5,
                },
                Classification: "small",
            },
            expectedCostDueToWeight: 8,
        },
        {
            name: "medium overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    d1:     30,
                    d2:     30,
                    d3:     30,
                    weight: 5.1,
                },
                Classification: "medium",
            },
            expectedCostDueToWeight: 6,
        },
        {
            name: "large overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    d1:     90,
                    d2:     90,
                    d3:     90,
                    weight: 8.9,
                },
                Classification: "large",
            },
            expectedCostDueToWeight: 6,
        },
        {
            name: "extra large overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    d1:     200,
                    d2:     300,
                    d3:     400,
                    weight: 14.2,
                },
                Classification: "extraLarge",
            },
            expectedCostDueToWeight: 10,
        },
        {
            name: "extra heavy parcel not overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    d1:     200,
                    d2:     300,
                    d3:     400,
                    weight: 50,
                },
                Classification: "extraHeavy",
            },
            expectedCostDueToWeight: 50,
        },
        {
            name: "extra heavy parcel overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    d1:     200,
                    d2:     300,
                    d3:     400,
                    weight: 50.4,
                },
                Classification: "extraHeavy",
            },
            expectedCostDueToWeight: 51,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            cost := tc.pricedParcel.CostDueToWeight()
            assert.Equal(t, tc.expectedCostDueToWeight, cost)
        })
    }
}