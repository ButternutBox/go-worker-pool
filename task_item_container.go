package dispatcher

// TaskItemContainer is an interface for a generic TaskItem
type TaskItemContainer interface {
	TaskItem() TaskItemContainer
}
