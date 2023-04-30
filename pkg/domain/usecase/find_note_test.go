package usecase_test

import (
	"testing"

	"github.com/psavelis/go-tuner-api/pkg/domain/entity"
	"github.com/psavelis/go-tuner-api/pkg/domain/usecase"
)

func TestGetFrequency(t *testing.T) {
	noteRepository := New(entity.StandardNote{
		Name: "A4",
	}, nil)

	findNoteUseCase := usecase.NewFindNoteUseCase(noteRepository)

	standardNote, err := findNoteUseCase.Execute(440)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if standardNote.Name != "A4" {
		t.Errorf("Expected note name to be A4, got %s", standardNote.Name)
	}
}

type NotesRepositoryMock struct {
	note entity.StandardNote
	err  error
}

func New(noteMock entity.StandardNote, err error) *NotesRepositoryMock {
	return &NotesRepositoryMock{
		note: noteMock,
	}
}

func (r *NotesRepositoryMock) Find(key int) (entity.StandardNote, error) {
	return r.note, r.err
}
