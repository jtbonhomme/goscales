package main

import (
    "fmt"
    "github.com/jtbonhomme/goscales/scales"
)

func main() {
    s := scales.New()

    fmt.Printf("All scales: %v\n", scales.Scales())

    s.FromDegrees("major", "1 2 3 4 5 6 7M")
    res, err := s.Transpose("C")
    if err != nil {
        panic(err)
    }
    fmt.Printf("C %s scale: %v\n", s.Name, res)

    s.FromDict("dorian")
    res, err = s.Transpose("C")
    if err != nil {
        panic(err)
    }
    fmt.Printf("C %s scale: %v\n", s.Name, res)

    res, err = s.Transpose("D")
    if err != nil {
        panic(err)
    }
    fmt.Printf("D %s scale: %v\n", s.Name, res)
}
