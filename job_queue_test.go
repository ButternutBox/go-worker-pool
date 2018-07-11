package dispatcher

import "testing"

func TestNewJobQueueNegativeMaxQueue(t *testing.T) {
	queue := NewJobQueue(-1)
	if queue != nil {
		t.Log("Expected", queue, "not to equal", nil)
	}
}

func TestNewJobQueuePositiveMaxQueue(t *testing.T) {
	queue := NewJobQueue(10)
	if len(queue) != 10 {
		t.Log("Expected", queue, "to equal", 10)
	}
}
