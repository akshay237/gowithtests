package clockface

import (
	"testing"
	"time"
)

func TestSecondHandAtMidNight(t *testing.T) {
	tm := time.Date(1997, time.January, 1, 0, 0, 0, 0, time.UTC)
	want := Point{
		X: 150,
		Y: 150 - 90,
	}
	got := SecondHand(tm)
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
