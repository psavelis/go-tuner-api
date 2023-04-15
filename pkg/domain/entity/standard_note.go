package domain

import "errors"

type StandardNote struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	KeyNumber    int            `json:"keyNumber"`
	Frequency    FrequencyRange `json:"frequency"`
	PitchPerfect float64        `json:"pitchPerfect"`
}

type FrequencyRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

var ErrNoteNotFound = errors.New("note not found")
