package dispatcher

// Job represents the job to be run
type Job struct {
	TaskItemContainer TaskItemContainer
	RunTask           func(TaskItemContainer)
}
