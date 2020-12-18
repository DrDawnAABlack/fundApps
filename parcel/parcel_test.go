package parcel

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestParcelPricingBySize(t *testing.T) {

    type testCase struct {
        name         string
        parcel       *Parcel
        expectedCost float64
        expectedClassification string
    }

    var testCases = []testCase{
        {
            name: "all dimensions small",
            parcel: &Parcel{
                Id: "small_parcel",
                D1: 2,
                D2: 3,
                D3: 4,
            },
            expectedCost: 3.00,
            expectedClassification: "small",
        },
        {
            name: "all dimensions medium",
            parcel: &Parcel{
                Id: "medium_parcel",
                D1: 25,
                D2: 36,
                D3: 42,
            },
            expectedCost: 8.00,
            expectedClassification: "medium",
        },
        {
            name: "all dimensions large",
            parcel: &Parcel{
                Id: "large_parcel",
                D1: 66,
                D2: 77,
                D3: 88,
            },
            expectedCost: 15.00,
            expectedClassification: "large",
        },
        {
            name: "all dimensions extra large",
            parcel: &Parcel{
                Id: "extra_large_parcel",
                D1: 100,
                D2: 100,
                D3: 100,
            },
            expectedCost: 25.00,
            expectedClassification: "extraLarge",
        },
        {
            name: "mixed dimensions medium",
            parcel: &Parcel{
                Id: "medium_parcel",
                D1: 1,
                D2: 1,
                D3: 42,
            },
            expectedCost: 8.00,
            expectedClassification: "medium",
        },
        {
            name: "mixed dimensions large",
            parcel: &Parcel{
                Id: "large_parcel",
                D1: 1,
                D2: 1,
                D3: 88,
            },
            expectedCost: 15.00,
            expectedClassification: "large",
        },
        {
            name: "mixed dimensions extra large",
            parcel: &Parcel{
                Id: "extra_large_parcel",
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
            cost, classification := CostDueToSize(tc.parcel)
            assert.Equal(t, tc.expectedCost, cost)
            assert.Equal(t, tc.expectedClassification, classification)
        })
    }
}
