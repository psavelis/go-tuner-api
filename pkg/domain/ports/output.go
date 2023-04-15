package ports

import domain "github.com/psavelis/go-tuner-api/pkg/domain/entity"

type NoteRepository interface {
	Find(key int) (domain.StandardNote, error)
}
