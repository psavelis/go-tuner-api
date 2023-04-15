package ports

import "github.com/psavelis/go-tuner-api/pkg/domain/entity"

type NoteRepository interface {
	Find(key int) (entity.StandardNote, error)
}
