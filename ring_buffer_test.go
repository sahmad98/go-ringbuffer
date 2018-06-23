package ringbuffer

import "testing"

func TestNewRingBuffer(t *testing.T) {
	rb := NewRingBuffer(5)
	if rb.Size != 5 {
		t.Errorf("NewRingBuffer creationg failed")
	}
}

func TestWritingBuffer(t *testing.T) {
	rb := NewRingBuffer(2)
	rb.Write(1)
	if x := rb.Read(); x != 1 {
		t.Errorf("Error Manipulating RingBuffer, Added 1, Got %d\n", x)
	}

	x := rb.Read()
	if x != nil {
		t.Errorf("Error Manipulating RingBuffer, No Element in Buffer, Found 1")
	}
}

func TestMultipleValues(t *testing.T) {
	rb := NewRingBuffer(2)
	rb.Write(2)
	rb.Write(3)
	rb.Write(4)

	a := rb.Read()
	b := rb.Read()

	if a != 4 || b != 3 {
		t.Errorf("Error Manipulating RingBuffer")
	}
}
