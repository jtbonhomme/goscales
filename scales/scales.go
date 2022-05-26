package scales

import (
    "strings"
    "fmt"
)

// Scale define a musical scale, built with degrees
type Scale struct {
    Name string
    intervales []int
}

// New creates a new Scale.
func New() *Scale {
    return &Scale{
        intervales: []int{},
    }
}

// FromNotes build a scale from a string which lists notes of the scale.
// A scale MUST contains at most 7 notes.
// ** NOT IMPLEMENTED **
// ex.: 'C D E F G A B'
func (s *Scale) FromNotes(name string, notes string) error {
    s.Name = name
    s.intervales = []int{}
    splitNotes := strings.Split(notes, " ")
    if len(splitNotes) > 7 || len(splitNotes) == 0 {
        return fmt.Errorf("incorrect number of notes in %s", notes)
    }
    return nil
}

// FromDegrees build a scale from a string which lists degrees of the scale.
// A scale must contains at most 7 degrees.
// ex.: '1 2 3m 4 5b 6 7M'
func (s *Scale) FromDegrees(name string, degrees string) error {
    s.Name = name
    s.intervales = []int{}
    splitDegrees := strings.Split(degrees, " ")
    if len(splitDegrees) > 7 || len(splitDegrees) == 0 {
        return fmt.Errorf("incorrect number of degrees in %s", degrees)
    }
    for _, d := range splitDegrees {
        i, ok := intervales[d]
        if !ok {
            return fmt.Errorf("unknown degree in %s ()", d)
        }
        s.intervales = append(s.intervales, i)
    }
    return nil
}

// FromDict build a scale from a string which correspond to one scale name from the dictionary.
// ex.: 'spanish'
func (s *Scale) FromDict(name string) error {
    s.Name = name
    s.intervales = []int{}
    degrees, ok := dictionary[name]
    if !ok {
        return fmt.Errorf("this scale is not referenced: %s", name)
    }
    splitDegrees := strings.Split(degrees, " ")
    for _, d := range splitDegrees {
        i, ok := intervales[d]
        if !ok {
            return fmt.Errorf("unknown degree in %s ()", d)
        }
        s.intervales = append(s.intervales, i)
    }
    return nil
}

// Note returns the note corresponding to a given degree from a scale.
// ** NOT IMPLEMENTED **
func (s *Scale) Note(degree string) ([]string, error) {
    res := []string{}
    return res, nil
}

// Degree returns the degree corresponding to a note from a scale.
// ** NOT IMPLEMENTED **
func (s *Scale) Degree(degree string) ([]string, error) {
    res := []string{}
    return res, nil
}


// Transpose returns the scale starting from the given tonic.
func (s *Scale) Transpose(tonic string) ([]string, error) {
    res := []string{}
    base, err := index(tonic)
    if err != nil {
        return res, fmt.Errorf("unknown note: %s", tonic)
    }

    for _, intervale := range s.intervales {
        res = append(res, fullScale[(base+intervale)%12])
    }

    return res, nil
}
