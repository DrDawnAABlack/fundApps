package main

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "./parcel"
)

func TestParser(t *testing.T) {
    type testCase struct {
        name           string
        input          string
        expectedParcel *parcel.Parcel
        expectedError  error
        speedyShipping bool
    }

    var testCases = []testCase{
        {
            name:  "happy parse",
            input: "98453 | 43 | 15 | 32425 ",
            expectedParcel: &parcel.Parcel{
                Id: "98453",
                D1: 43,
                D2: 15,
                D3: 32425,
            },
            expectedError: nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            parcel, err := parseDimensions(tc.input)
            assert.Equal(t, tc.expectedParcel, parcel)
            assert.Equal(t, tc.expectedError, err)
        })
    }
}
