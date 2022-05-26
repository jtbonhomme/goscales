package scales

import (
    "strings"
    "fmt"
)

// Scales is a collection of scales
type Scales []Scale

// Scale define a musical scale, built with degrees
type Scale struct {
    Name string
    intervales []int
}

// New creates a new Scale.
func New(name string) *Scale {
    return &Scale{
        Name: name,
        intervales: []int{},
    }
}

// FromNotes build a scale from a string which lists notes of the scale.
// A scale MUST contains 7 notes.
// ** NOT IMPLEMENTED **
// ex.: 'C D E F G A B'
func (s *Scale) FromNotes(notes string) error {
    splitNotes := strings.Split(notes, " ")
    if len(splitNotes) != 7 {
        return fmt.Errorf("incorrect number of notes in %s", notes)
    }
    return nil
}

// FromNotes build a scale from a string which lists degrees of the scale.
// A scale must contains 7 degrees.
// ex.: '1 2 3m 4 5b 6 7M'
func (s *Scale) FromDegrees(degrees string) error {
    splitDegrees := strings.Split(degrees, " ")
    if len(splitDegrees) != 7 {
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
