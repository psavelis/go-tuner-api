package usecase

import (
	"fmt"
	"math"
	"strings"

	"github.com/psavelis/go-tuner-api/pkg/domain/entity"
	"github.com/psavelis/go-tuner-api/pkg/domain/ports"
)

// For more info about why I chose 432Hz as the tempering frequency, see: https://www.izotope.com/en/learn/tuning-standards-explained.html
const temperingHz = 432

type usecase struct {
	repository ports.NoteRepository
}

func New(repository ports.NoteRepository) ports.FindNoteByFrequency {
	return &usecase{repository: repository}
}

func getFrequency(key int) float64 {
	return temperingHz * math.Pow(2, (float64(key)-49)/12)
}

func getKey(frequency float64) int {
	return int(12*math.Log2(frequency/temperingHz) + 49)
}

func (uc *usecase) Execute(frequencyInHz float64) (entity.StandardNote, error) {
	key := getKey(frequencyInHz)

	note, err := uc.repository.Find(key)

	if err != nil {
		return entity.StandardNote{}, entity.ErrNoteNotFound
	}

	min := getFrequency(key - 1)
	max := getFrequency(key + 1)
	exact := getFrequency(key)

	boundary := max / min

	note.Frequency = entity.FrequencyRange{Min: min + boundary, Max: max - boundary}
	note.PitchPerfect = exact

	fmt.Println(frequencyInHz, "Hz =>", strings.Split(note.Name, "/")[0])

	return note, nil
}
