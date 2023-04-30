package usecase

import "github.com/psavelis/go-tuner-api/pkg/domain/ports"

type findMaximaUseCase struct {
}

func NewFindMaximaUseCase() ports.FindMaximaBySampleRate {
	return &findMaximaUseCase{}
}

func (u *findMaximaUseCase) Execute(sampleRate float64, audioData []float64) (float64, error) {
	signalCorrelation := map[int]float64{}

	max :=  map[int]int{}

	maxCount := 0
	fftLenght := len(audioData)

	for l := 0; l < fftLenght; l++ {
		signalCorrelation[l] = 0
		for i := 0; i < fftLenght-l; i++ {
			signalCorrelation[l] += audioData[i] * audioData[i+l]
		}
		if l > 1 {
			newMaxima := (signalCorrelation[l-2]-signalCorrelation[l-1]) < 0 &&
				(signalCorrelation[l-1]-signalCorrelation[l]) > 0

			if newMaxima {
				max[maxCount] = (l - 1)

				maxCount++

				if maxCount >= len(max) {
					break
				}
			}
		}
	}

	maximaMean := max[0]

	for i := 1; i < maxCount; i++ {
		maximaMean += max[i] - max[i-1]
	}

	maximaMean /= maxCount

	return sampleRate / float64(maximaMean), nil
}
