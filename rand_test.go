// Public domain.

package lcg_test

import (
	"testing"

	"github.com/soniakeys/lcg"
)

func TestLCG(t *testing.T) {
	r := lcg.New()
	for _, want := range []float64{
		.0015762136730836773,
		.38537207475475027,
		.6771517073363665,
	} {
		if got := r.Float64(); got != want {
			t.Fatalf("got %f, want %f\n", got, want)
		}
	}
	r.Seed(4)
	for _, want := range []float64{
		.0021016182307782363,
		.18049609967300034,
		.9028689431151553,
	} {
		if got := r.Float64(); got != want {
			t.Fatalf("got %f, want %f\n", got, want)
		}
	}
}
