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
                D1: 2,
                D2: 3,
                D3: 4,
            },
            expectedCost: 3.00,
            expectedClassification: "small",
        },
        {
            name: "all dimensions medium",
            parcel: Parcel{
                D1: 25,
                D2: 36,
                D3: 42,
            },
            expectedCost: 8.00,
            expectedClassification: "medium",
        },
        {
            name: "all dimensions large",
            parcel: Parcel{
                D1: 66,
                D2: 77,
                D3: 88,
            },
            expectedCost: 15.00,
            expectedClassification: "large",
        },
        {
            name: "all dimensions extra large",
            parcel: Parcel{
                D1: 100,
                D2: 100,
                D3: 100,
            },
            expectedCost: 25.00,
            expectedClassification: "extraLarge",
        },
        {
            name: "mixed dimensions medium",
            parcel: Parcel{
                D1: 1,
                D2: 1,
                D3: 42,
            },
            expectedCost: 8.00,
            expectedClassification: "medium",
        },
        {
            name: "mixed dimensions large",
            parcel: Parcel{
                D1: 1,
                D2: 1,
                D3: 88,
            },
            expectedCost: 15.00,
            expectedClassification: "large",
        },
        {
            name: "mixed dimensions extra large",
            parcel: Parcel{
                D1: 1,
                D2: 1,
                D3: 100,
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
        name         string
        pricedParcel       PricedParcel
        expectedExtraWeightCost float64
    }

    var testCases = []testCase{
        {
            name: "small not overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    D1:     2,
                    D2:     3,
                    D3:     4,
                    Weight: smallParcelWeightLimit,
                },
                Classification: "small",
            },
            expectedExtraWeightCost: 0,
        },
        {
            name: "medium not overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    D1:     30,
                    D2:     30,
                    D3:     30,
                    Weight: mediumParcelWeightLimit,
                },
                Classification: "medium",
            },
            expectedExtraWeightCost: 0,
        },
        {
            name: "large not overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    D1:     90,
                    D2:     90,
                    D3:     90,
                    Weight: largeParcelWeightLimit,
                },
                Classification: "large",
            },
            expectedExtraWeightCost: 0,
        },
        {
            name: "extra large not overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    D1:     200,
                    D2:     300,
                    D3:     400,
                    Weight: extraLargeParcelWeightLimit,
                },
                Classification: "extraLarge",
            },
            expectedExtraWeightCost: 0,
        },
        {
            name: "small overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    D1:     2,
                    D2:     3,
                    D3:     4,
                    Weight: 5,
                },
                Classification: "small",
            },
            expectedExtraWeightCost: 8,
        },
        {
            name: "medium overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    D1:     30,
                    D2:     30,
                    D3:     30,
                    Weight: 5.1,
                },
                Classification: "medium",
            },
            expectedExtraWeightCost: 6,
        },
        {
            name: "large overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    D1:     90,
                    D2:     90,
                    D3:     90,
                    Weight: 8.9,
                },
                Classification: "large",
            },
            expectedExtraWeightCost: 6,
        },
        {
            name: "extra large overweight",
            pricedParcel: PricedParcel{
                Parcel: &Parcel{
                    D1:     200,
                    D2:     300,
                    D3:     400,
                    Weight: 14.2,
                },
                Classification: "extraLarge",
            },
            expectedExtraWeightCost: 10,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            cost := tc.pricedParcel.CostDueToWeight()
            assert.Equal(t, tc.expectedExtraWeightCost, cost)
        })
    }
}