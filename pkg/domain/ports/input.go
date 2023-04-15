package ports

import "github.com/psavelis/go-tuner-api/pkg/domain/entity"

type FindNoteByFrequency interface {
	Find(frequencyInHz float64) (entity.StandardNote, error)
}
