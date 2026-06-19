package logics

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(NewTaskName string, NewTaskDescription string) Task {

	return Task{
		Title:       NewTaskName,
		Description: NewTaskDescription,
		Completed:   false,

		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

}

func (t *Task) Complete() {
	CompletedAtTime := time.Now()

	t.Completed = true
	t.CompletedAt = &CompletedAtTime
}

func (t *Task) Uncomplete() {

	t.Completed = false
	t.CompletedAt = nil
}
