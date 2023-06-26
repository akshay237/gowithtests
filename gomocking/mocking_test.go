package gomocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// Mocking is replace real things you inject with a pretend version that you can control and inspect in your tests
// Refactoring means that the code changes but the behaviour should remains same.
// Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful but that means
// a tighter coupling between your test code and the implementation.
func TestCountDown(t *testing.T) {
	// 	buffer := &bytes.Buffer{}
	// 	spy := &SpySleeper{}
	// 	CountDown(buffer, spy)

	// 	got := buffer.String()
	// 	want := `3
	// 2
	// 1
	// Go!`

	// if got != want {
	// 	t.Errorf("got %q but want %q ", got, want)
	// }

	// if spy.Calls != 3 {
	// 	t.Errorf("not enough calls to sleeper want 3 but got %d", spy.Calls)
	// }

	t.Run(" print sleep after every operation", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		track := &SpyOperationsTrack{}
		CountDown(buffer, track)

		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q but want %q ", got, want)
		}
	})

	t.Run("print before ever write", func(t *testing.T) {

		track := &SpyOperationsTrack{}
		CountDown(track, track)
		got := track.calls
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted calls %v got %v", want, got)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("test configurable sleeper ", func(t *testing.T) {
		timeSleep := 5 * time.Second

		spyTime := &SpyTime{}
		configurableTime := &ConfigurableSleeper{timeSleep, spyTime.Sleep}
		configurableTime.Sleep()

		if spyTime.durationSlept != timeSleep {
			t.Errorf("should have slept for %v but slept for %v", timeSleep, spyTime.durationSlept)
		}
	})
}
