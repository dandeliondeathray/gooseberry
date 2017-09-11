package director

// Context has immutable data on the context of the job.
type Context struct{}

// Work is the director's interface to the work a Step should perform.
type Work interface {
	Schedule()
}

// Step is the smallest element of a pipeline.
type Step struct {
	work Work
}

// Execute schedules all dependencies, or schedules the work if there are no dependencies.
func (s *Step) Execute() {
	s.work.Schedule()
}

// NewStep creates a new Step instance.
func NewStep(work Work) *Step {
	return &Step{work: work}
}
