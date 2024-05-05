package selection

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/Sermelyan/alg/utils"
)

func TestSelectionInts(t *testing.T) {
	for i := 0; i < 500; i++ {
		n := rand.Intn(10000)
		s := utils.GenRandomSlice(n, utils.RandomNumber(rand.Int))
		Selection(s)
		if !slices.IsSorted(s) {
			t.Error("expected slice to be sorted", s)
		}
	}
}

func TestSelectionFloats(t *testing.T) {
	for i := 0; i < 500; i++ {
		n := rand.Intn(10000)
		s := utils.GenRandomSlice(n, utils.RandomNumber(rand.NormFloat64))
		Selection(s)
		if !slices.IsSorted(s) {
			t.Error("expected slice to be sorted", s)
		}
	}
}

func Benchmark(b *testing.B) {
	const n = 10000
	sI := utils.GenRandomSlice(n, utils.RandomNumber(rand.Int))
	sF := utils.GenRandomSlice(n, utils.RandomNumber(rand.NormFloat64))
	b.ResetTimer()

	b.Run("insert ints", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Selection(sI)
			b.StopTimer()
			rand.Shuffle(n, func(i, j int) {
				sI[i], sI[j] = sI[j], sI[i]
			})
			b.StartTimer()
		}
	})
	b.Run("slices.sort ints", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slices.Sort(sI)
			b.StopTimer()
			rand.Shuffle(n, func(i, j int) {
				sI[i], sI[j] = sI[j], sI[i]
			})
			b.StartTimer()
		}
	})
	b.Run("insert floats", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Selection(sF)
			b.StopTimer()
			rand.Shuffle(n, func(i, j int) {
				sF[i], sF[j] = sF[j], sF[i]
			})
			b.StartTimer()
		}
	})
	b.Run("slices.sort floats", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slices.Sort(sF)
			b.StopTimer()
			rand.Shuffle(n, func(i, j int) {
				sF[i], sF[j] = sF[j], sF[i]
			})
			b.StartTimer()
		}
	})
}
