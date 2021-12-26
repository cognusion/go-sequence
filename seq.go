// Package sequence provides a performant goro-safe incrementing sequencer,
// with randomized HashID-encoding of the values to work as a font of IDs for errors,
// requests, etc.
package sequence

import (
	"github.com/speps/go-hashids/v2"

	"crypto/rand"
	"sync/atomic"
)

// Seq is a sequencer. It should never be used by default, and instead obtained via a
// New function
type Seq struct {
	hashid  *hashids.HashID
	counter int64
}

// New returns an initialized Seq suitable for most cases
func New(start int) *Seq {
	return NewWithHashIDLength(start, 7)
}

// NewWithHashIDLength returns an initilized Seq with the minimum hashID length set as specified
func NewWithHashIDLength(start, hashIDlength int) *Seq {
	b := make([]byte, 16)
	rand.Read(b)
	hd := hashids.NewData()
	hd.Salt = string(b)
	hd.MinLength = hashIDlength
	hashid, _ := hashids.NewWithData(hd)

	return &Seq{
		counter: int64(start),
		hashid:  hashid,
	}
}

// Next returns the next int in the sequence
func (s *Seq) Next() int {
	val := int(atomic.AddInt64(&s.counter, 1))
	return val
}

// NextHashID returns the hashID-encoded value of the next int in the sequence
func (s *Seq) NextHashID() string {
	val := s.Next()
	hash, _ := s.hashid.Encode([]int{val}) // Yup, skipping errors like a dummy
	return hash
}
