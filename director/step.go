package director

// Context has immutable data on the context of the job.
type Context struct{}

// Work is the director's interface to the work a Step should perform.
type Work interface {
	Schedule()
}

// Dependency is an abstract dependency of a Step.
// Note that Step implements this interface.
type Dependency interface {
	Execute()
}

// Step is the smallest element of a pipeline.
type Step struct {
	work         Work
	dependencies []Dependency
}

// Execute schedules all dependencies, or schedules the work if there are no dependencies.
func (s *Step) Execute() {
	if len(s.dependencies) == 0 {
		s.work.Schedule()
	} else {
		for i := range s.dependencies {
			s.dependencies[i].Execute()
		}
	}
}

// NewStep creates a new Step instance.
func NewStep(work Work, dependencies ...Dependency) *Step {
	return &Step{work: work, dependencies: dependencies}
}
