package gosync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increment the counter 3 times", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		got := counter.Value()
		want := 3
		asserCounter(t, got, want)
	})

	t.Run("run using wait group concurrently", func(t *testing.T) {
		countWanted := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(countWanted)

		for i := 0; i < countWanted; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()
		got := counter.Value()
		asserCounter(t, got, countWanted)
	})
}

func asserCounter(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
