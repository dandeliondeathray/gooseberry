package director

// CommitID is the hash of a Git commit
type CommitID string

// Context has immutable data on the context of the job.
type Context struct {
	CommitID CommitID
}

// Work is the director's interface to the work a Step should perform.
type Work interface {
	Schedule()
}

// Dependency is an abstract dependency of a Step.
// Note that Step implements this interface.
type Dependency interface {
	Execute()
}

// OngoingStep is the state of a step when it waits for a dependency or work to complete.
type OngoingStep struct {
	work Work
}

// DependencyComplete is called when the dependency has completed with some result.
func (s *OngoingStep) DependencyComplete(dependency Dependency, result Result) {
	s.work.Schedule()
}

// Step is the smallest element of a pipeline.
type Step struct {
	work         Work
	dependencies []Dependency
}

// Execute schedules all dependencies, or schedules the work if there are no dependencies.
func (s *Step) Execute() *OngoingStep {
	if len(s.dependencies) == 0 {
		s.work.Schedule()
	} else {
		for i := range s.dependencies {
			s.dependencies[i].Execute()
		}
	}
	return &OngoingStep{work: s.work}
}

// NewStep creates a new Step instance.
func NewStep(work Work, dependencies ...Dependency) *Step {
	return &Step{work: work, dependencies: dependencies}
}
