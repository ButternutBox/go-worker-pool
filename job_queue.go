package dispatcher

// NewJobQueue returns a buffered channel that we can send work requests on
func NewJobQueue(maxQueue int) chan Job {
	// return error
	if maxQueue < 0 {
		return nil
	}
	return make(chan Job, maxQueue)
}
