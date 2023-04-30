package ports

import "github.com/psavelis/go-tuner-api/pkg/domain/entity"

type FindNoteByFrequency interface {
	Execute(frequencyInHz float64) (entity.StandardNote, error)
}

type FindMaximaBySampleRate interface {
	Execute(sampleRate float64, audioData []float64) (float64, error)
}

