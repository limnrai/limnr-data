package repos

import (
	"github.com/brianvoe/gofakeit/v6"
)

func GetFakeInt(min int, max int) int {
	return gofakeit.Number(min, max)
}

func GetFakeFloat(min float64, max float64) float64 {
	return gofakeit.Float64Range(min, max)
}

func GetFakeString(minLength int, maxLength int) string {
	return gofakeit.LetterN(uint(gofakeit.Number(minLength, maxLength)))
}
