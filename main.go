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

    totalCost := float64(0)
    //price each parcel
    for {
        row, err := reader.ReadString('\n')
        if err != nil && err != io.EOF {
            log.Fatalln(err)
        }
        if err == io.EOF {
            break
        }

        // assumed input format
        // parcel_id | d1 | d2 | d3
        // where d represents a dimension
        newParcel, err := ParseArgs(row)
        if err != nil {
            log.Fatalln(err)
        }

        pricedParcel := parcel.CostDueToSize(newParcel)

        pricedParcelOutput := fmt.Sprintf("%s | %.2f \n", pricedParcel.Parcel.Id, pricedParcel.Cost)
        fmt.Println(pricedParcelOutput)

        _, err = writer.WriteString(pricedParcelOutput)
        if err != nil {
            log.Fatalln(err)
        }

        totalCost = totalCost + pricedParcel.Cost
    }
    // output the final cost
    totalCostMsg := fmt.Sprintf("Total = %.2f \n", totalCost)
    fmt.Println(totalCostMsg)
    _, err = writer.WriteString(totalCostMsg)
    writer.Flush()
    if err != nil {
        log.Fatalln(err)
    }
}


func ParseArgs(args string) (*parcel.Parcel, error) {
    vars := strings.Split(args, "|")
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
        return nil, errors.New(fmt.Sprintf("at least one dimensio is invalid [%v]", args))
    }
    parcel := parcel.New(strings.TrimSpace(vars[0]), int64(d1), int64(d2), int64(d3))

    return parcel, nil
}