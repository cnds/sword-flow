package flow

// Task is user define, workflow item
type Task interface {
}

// ArcGurand Arc is the connection from Transition to Place
type ArcGurand interface {
	Permit(task *Task) (bool, error)
}

// TransitionTrigger trigger for on task enabled
type TransitionTrigger interface {
	OnTaskEnabled(task *Task) error
}
