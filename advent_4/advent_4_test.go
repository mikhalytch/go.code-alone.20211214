package main

import (
	"reflect"
	"testing"
)

func TestReadAdvent4File(t *testing.T) {
	t.Run("read sample", func(t *testing.T) {
		filename := "advent_4.sample"
		want := advent4File{4, []int{1, 2, 1, 2}}
		got := readAdvent4File(filename)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("read real file", func(t *testing.T) {
		filename := "advent_4.test.txt"
		wantLengths := 212850
		got := readAdvent4File(filename)

		assertNumbers(t, got.lengthsAmt, wantLengths)
		assertNumbers(t, len(got.lengths), wantLengths)
	})
}

func assertNumbers(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("Got %v lenghts, want %v", got, want)
	}
}
