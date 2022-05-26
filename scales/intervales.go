package scales

import (
    "fmt"
)

var alterations map[string]int = map[string]int{
    "m":  -1,
    "M":   1,
    "dim":  -1,
    "aug":  1,
}

var intervales map[string]int = map[string]int{
    "1": 0,
    "2": 2,
    "3m": 3,
    "3b": 3,
    "3": 4,
    "4": 5,
    "4#": 6,
    "5D": 6,
    "5": 7,
    "5A": 8,
    "6": 9,
    "7b": 10,
    "7m": 10,
    "7": 11,
    "7M": 11,
}

var fullScale []string = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "Bb", "B"}

func index(note string) (int, error) {
       for i, n := range fullScale {
                if note == n {
                        return i, nil
                }
        }
        return 0, fmt.Errorf("%s is not a valid note", note)
}

// Semitones return the number of semi tones that corresponds to an intervale.
// An intervale can be a non altered degre: 1, 2, 3, ..., or an altered degre:
// 3m, 5dim
// ** NOT IMPLEMENTED **
func (s *Scale) Semitones(intervale string) int {
    return 0
}
