// Public domain.

// LCG implements a linear congruential generator.
package lcg

import "math"

var (
	lcga, lcgm uint64
	invLcgm    float64
)

func init() {
	lcga = uint64(math.Pow(13, 13))
	lcgm = 1 << 59
	invLcgm = 1 / float64(lcgm)
	lcgm--
}

// Rand generates numbers like the generator of math/rand but implements
// only a little bit of the math/rand.Rand interface.
type Rand struct {
	uint64
}

// New returns a new generator.
//
// It does not follow the Rand interface of math/rand.
func New() *Rand {
	return &Rand{3}
}

// Float64 generates a float64 in the range (0,1).
//
// It implements Rand.Float64 of math/rand.
func (r *Rand) Float64() float64 {
	r.uint64 = r.uint64 * lcga & lcgm
	return float64(r.uint64) * invLcgm
}

// Seed seeds the generator.
//
// It implements Rand.Seed of math/rand.
func (r *Rand) Seed(s int64) {
	r.uint64 = uint64(s)
}
