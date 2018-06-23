package ringbuffer

type RingBuffer struct {
	Size      int
	Container []interface{}
	Reader    int
	Writer    int
}

func NewRingBuffer(size int) *RingBuffer {
	rb := new(RingBuffer)
	rb.Size = size
	rb.Container = make([]interface{}, size)
	rb.Reader = 0
	rb.Writer = 0
	return rb
}

func (r *RingBuffer) Write(v interface{}) {
	r.Container[r.Writer] = v
	r.Writer = (r.Writer + 1) % r.Size
}

func (r *RingBuffer) seekReader(delta int) {
	r.Reader = (r.Reader + delta) % r.Size
}

func (r *RingBuffer) Read() interface{} {
	defer r.seekReader(1)
	return r.Container[r.Reader]
}
