# ringbuffer

Simple RingBuffer pacakge Ringbuffer is non blocking for readers and writers,
writers will overwrite older data in a circular fashion. Readers will read from
the current position and update it.

## Usage

#### type RingBuffer

```go
type RingBuffer struct {
	Size      int           // Size of the Ringbuffer
	Container []interface{} // Array container of objects
	Reader    int           // Reader position
	Writer    int           // Writer Position
}
```

RingBuffer Structure

#### func  NewRingBuffer

```go
func NewRingBuffer(size int) *RingBuffer
```
Create a new RingBuffer of initial size "size" Returns a pointer to the new
RingBuffer

#### func (*RingBuffer) Read

```go
func (r *RingBuffer) Read() interface{}
```
Read single object from the RingBuffer

#### func (*RingBuffer) Write

```go
func (r *RingBuffer) Write(v interface{})
```
Write object into the RingBuffer
