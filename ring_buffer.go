// Simple RingBuffer pacakge
// Ringbuffer is non blocking for readers and writers, writers will
// overwrite older data in a circular fashion. Readers will read
// from the current position and update it.
package ringbuffer


// RingBuffer Structure 
type RingBuffer struct {
	Size      int            // Size of the Ringbuffer 
	Container []interface{}  // Array container of objects 
	Reader    int            // Reader position
	Writer    int            // Writer Position
}

// Create a new RingBuffer of initial size "size"
// Returns a pointer to the new RingBuffer
func NewRingBuffer(size int) *RingBuffer {
	rb := new(RingBuffer)
	rb.Size = size
	rb.Container = make([]interface{}, size)
	rb.Reader = 0
	rb.Writer = 0
	return rb
}


// Write object into the RingBuffer
func (r *RingBuffer) Write(v interface{}) {
	r.Container[r.Writer] = v
	r.Writer = (r.Writer + 1) % r.Size
}

// Seek position of the reader by delta, delta can be positive as
// well as negative
func (r *RingBuffer) seekReader(delta int) {
	r.Reader = (r.Reader + delta) % r.Size
}


// Read single object from the RingBuffer
func (r *RingBuffer) Read() interface{} {
	defer r.seekReader(1)
	return r.Container[r.Reader]
}
