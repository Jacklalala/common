package random

import (
	crand "crypto/rand"
	"encoding/hex"
	"io"
	"math"
	mrand "math/rand"
	"sync"
)

var grand *Rand

type Rand struct {
	sync.Mutex
	rand *mrand.Rand
}

func init() {
	grand = NewRand()
	grand.init()
}

func (r *Rand) init() {
	bz := Value(8)
	var seed uint64
	for i := 0; i < 8; i++ {
		seed |= uint64(bz[i])
		seed <<= 8
	}
	r.reset(int64(seed))
}

func (r *Rand) reset(seed int64) {
	r.rand = mrand.New(mrand.NewSource(seed))
}

func (r *Rand) Seed(seed int64) {
	r.Lock()
	r.reset(seed)
	r.Unlock()
}

func (r *Rand) Uint32() uint32 {
	r.Lock()
	u32 := r.rand.Uint32()
	r.Unlock()
	return u32
}

func (r *Rand) Uint64() uint64 {
	return uint64(r.Uint32())<<32 + uint64(r.Uint32())
}

func (r *Rand) Int() int {
	r.Lock()
	i := r.rand.Int()
	r.Unlock()
	return i
}

func (r *Rand) Int31() int32 {
	r.Lock()
	i31 := r.rand.Int31()
	r.Unlock()
	return i31
}

func (r *Rand) Int31n(n int32) int32 {
	r.Lock()
	i31n := r.rand.Int31n(n)
	r.Unlock()
	return i31n
}

func (r *Rand) Int63() int64 {
	r.Lock()
	i63 := r.rand.Int63()
	r.Unlock()
	return i63
}

func (r *Rand) Int63n(n int64) int64 {
	r.Lock()
	i63n := r.rand.Int63n(n)
	r.Unlock()
	return i63n
}

func (r *Rand) Float32() float32 {
	r.Lock()
	f32 := r.rand.Float32()
	r.Unlock()
	return f32
}

func (r *Rand) Float64() float64 {
	r.Lock()
	f64 := r.rand.Float64()
	r.Unlock()
	return f64
}

// Intn returns, as an int, a uniform pseudo-random number in the range [0, n).
// It panics if n <= 0.
func (r *Rand) Intn(n int) int {
	r.Lock()
	i := r.rand.Intn(n)
	r.Unlock()
	return i
}

// Perm returns a pseudo-random permutation of n integers in [0, n).
func (r *Rand) Perm(n int) []int {
	r.Lock()
	perm := r.rand.Perm(n)
	r.Unlock()
	return perm
}

func NewRand() *Rand {
	rand := &Rand{}
	rand.init()
	return rand
}

func RandInt() int {
	return grand.Int()
}

func RandFloat64() float64 {
	return grand.Float64()
}

func RandIntn(n int) int {
	return grand.Intn(n)
}

func RandPerm(n int) []int {
	return grand.Perm(n)
}

// ValueHex returns a hex encoded string that's floor(size/2) * 2 long.
// Note: ValueHex(24) gives 96 bits of randomness that
// are usually strong enough for most purposes.
func ValueHex(size int) string {
	return hex.EncodeToString(Value(size / 2))
}

func Value(size int) []byte {
	b := make([]byte, size)
	crand.Read(b)
	return b
}

// Returns a crand.Reader.
func Reader() io.Reader {
	return crand.Reader
}

// round returns x rounded to the nearest int64 (non-negative values only).
func Round(x float64) int64 {
	if _, frac := math.Modf(x); frac >= 0.5 {
		return int64(math.Ceil(x))
	}
	return int64(math.Floor(x))
}

func RandomCSPRNG(n int) []byte {
	buff := make([]byte, n)
	_, err := io.ReadFull(crand.Reader, buff)
	if err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	return buff
}
