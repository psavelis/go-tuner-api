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

type findNoteUseCase struct {
	repository ports.NoteRepository
}

func NewFindNoteUseCase(repository ports.NoteRepository) ports.FindNoteByFrequency {
	return &findNoteUseCase{repository: repository}
}

func (uc *findNoteUseCase) Execute(frequencyInHz float64) (entity.StandardNote, error) {
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

	note.GaugePercent = 50 + ((frequencyInHz - exact) / (max - exact)) * 50

	if (frequencyInHz > exact) {
		note.GaugePercent = 50 - ((exact - frequencyInHz) / (exact - min)) * 50
	}

	if (note.GaugePercent > 100) {
		note.GaugePercent -= 100
	}

	fmt.Println(frequencyInHz, "Hz =>", strings.Split(note.Name, "/")[0], "Gauge =>", note.GaugePercent, "%")
 
	return note, nil
}

func getFrequency(key int) float64 {
	return temperingHz * math.Pow(2, (float64(key)-49)/12)
}

func getKey(frequency float64) int {
	return int(12*math.Log2(frequency/temperingHz) + 49)
}
