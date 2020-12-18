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
        expectedError  string
    }

    var testCases = []testCase{
        {
            name:  "happy parse",
            input: "98453 | 43 | 15 | 32425 | 276.02",
            expectedParcel: parcel.NewParcel("98453", 43, 15, 32425, 276.02),
            expectedError: "",
        },
        {
           name:  "parse with invalid weight",
           input: "98453 | 43 | 15 | 32425 | -6",
           expectedParcel: nil,
           expectedError: "the weight is invalid [-6]",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            parcel, err := parseDimensions(tc.input)
            assert.Equal(t, tc.expectedParcel, parcel)
            if err != nil {
                assert.Contains(t, err.Error(), tc.expectedError)
            }
        })
    }
}
