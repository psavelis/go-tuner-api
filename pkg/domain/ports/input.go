package ports

import domain "github.com/psavelis/go-tuner-api/pkg/domain/entity"

type NoteService interface {
	Find(frequencyInHz float64) (domain.StandardNote, error)
}
