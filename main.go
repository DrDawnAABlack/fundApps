package main

import(
    "bufio"
    "os"
    "errors"
    "fmt"
    "log"
    "io"
    "path/filepath"
    "strings"
    "strconv"

    "./parcel"
)

func main() {

    fileName := os.Args[1]
    intputFD, err := os.Open(fileName)
    if err != nil {
        log.Fatalln(err)
    }
    defer intputFD.Close()

    baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
    outputFile := baseName + "_invoice.txt"
    outputFD, err := os.Create(outputFile)
    if err != nil {
        log.Fatalln(err)
    }
    defer outputFD.Close()

    reader := bufio.NewReader(intputFD)
    writer := bufio.NewWriter(outputFD)

    order := parcel.NewOrder()
    totalCostWithoutDiscount := float64(0)
    speedyShipping := false
    //price each parcel
    for {
        row, err := reader.ReadString('\n')
        if err != nil && err != io.EOF {
            log.Fatalln(err)
        }
        if err == io.EOF {
            break
        }

        // assumed input format for each line
        // parcel_id | d1 | d2 | d3
        // where d represents a dimension
        newParcel, err := parseDimensions(row)
        if err != nil {
            speedyShipping, err = isSpeedyShipping(row)
            if speedyShipping{
                // assume last row
                break
            }
            log.Fatalln(err)
        }

        pricedParcel := parcel.PricedParcel{
            Parcel: newParcel,
        }

        if pricedParcel.Parcel.IsExtraHeavy() {
            pricedParcel.Classification = "extraHeavy"
            pricedParcel.Cost = pricedParcel.CostDueToWeight()
        } else {
            pricedParcel.Cost, pricedParcel.Classification = newParcel.CostDueToSize()
            pricedParcel.Cost = pricedParcel.Cost + pricedParcel.CostDueToWeight()
        }

        order.PricedParcels[pricedParcel.Parcel.Id] = &pricedParcel

        outputResults(fmt.Sprintf("%s | %.2f \n", pricedParcel.Parcel.Id, pricedParcel.Cost), writer)

        totalCostWithoutDiscount = totalCostWithoutDiscount + pricedParcel.Cost
    }

    order.CostOfShipping = totalCostWithoutDiscount

    // apply any discounts


    // output the final cost of the parcels
    outputResults(fmt.Sprintf("Total = %.2f \n", order.CostOfShipping), writer)
    if speedyShipping {
        order.CostOfSpeedyShipping = totalCostWithoutDiscount * 2
        outputResults(fmt.Sprintf("Total with speedy shipping = %.2f \n", order.CostOfSpeedyShipping), writer)
    }

    writer.Flush()
}

func outputResults(msg string, writer *bufio.Writer) {
    fmt.Println(msg)
    _, err := writer.WriteString(msg)
    if err != nil {
        log.Fatalln(err)
    }
}
func isSpeedyShipping(row string) (bool, error){
    fmt.Println(row)
    if row == "" {
        return false, nil
    }
    // would love to get rid of the \n here
    if row == "speedy\n" {
        return true, nil
    }
    return false, errors.New("invalid input format")
}

func parseDimensions(row string) (*parcel.Parcel, error) {
    vars := strings.Split(row, "|")
    if len(vars) != parcel.NumParcelFields {
        return nil, errors.New("not enough input fields to create a package")
    }

    id := vars[0]
    if id == "" {
        return nil, errors.New("parcel ID cannot be empty")
    }

    d1, err := strconv.Atoi(strings.TrimSpace(vars[1]))
    if err != nil {
        return nil, errors.New(fmt.Sprintf("error converting input [%v]: [%v]", vars, err))
    }
    d2, err := strconv.Atoi(strings.TrimSpace(vars[2]))
    if err != nil {
        return nil, errors.New(fmt.Sprintf("error converting input [%v]: [%v]", vars, err))
    }
    d3, err := strconv.Atoi(strings.TrimSpace(vars[3]))
    if err != nil {
        return nil, errors.New(fmt.Sprintf("error converting input [%v]: [%v]", vars, err))
    }
    if d1 <= 0 || d2 <= 0 || d3 <= 0 {
        return nil, errors.New(fmt.Sprintf("at least one dimension is invalid [%v]", vars))
    }
    weight, err := strconv.ParseFloat(strings.TrimSpace(vars[4]), 64)
    if weight <= 0 {
        return nil, errors.New(fmt.Sprintf("the weight is invalid [%s]", strings.TrimSpace(vars[4])))
    }
    parcel := parcel.NewParcel(strings.TrimSpace(vars[0]), int64(d1), int64(d2), int64(d3), weight)

    return parcel, nil
}