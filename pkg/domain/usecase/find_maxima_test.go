package usecase_test

import (
	"testing"

	"github.com/psavelis/go-tuner-api/pkg/domain/usecase"
)

func TestFindMaxima(t *testing.T) {
	findMaximaUseCase := usecase.NewFindMaximaUseCase()

	res, err := findMaximaUseCase.Execute(44000, []float64{0.1, 0.2, -0.3, -0.4, 0.5})

	if res > 14667 || res < 14666 {
		t.Errorf("Expected frequency to be 14666.666~ , got %f", res)
	}

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
