package main

import (
    "fmt"
    "github.com/jtbonhomme/goscales/scales"
)

func main() {
    s := scales.New("C major")
    s.FromDegrees("1 2 3 4 5 6 7M")
    res, err := s.Transpose("C")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s scale: %v\n", s.Name, res)
}
