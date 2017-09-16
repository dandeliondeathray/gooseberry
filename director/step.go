package director

// CommitID is the hash of a Git commit
type CommitID string

// Context has immutable data on the context of the job.
type Context struct {
	CommitID CommitID
}

// ExecutionContext maintains data on what steps have been executed.
type ExecutionContext struct {
	executedDependencies map[Dependency]bool
}

// HasDependencyBeenExecuted answers if a dependency has already been executed in this
// execution trigger.
func (s *ExecutionContext) HasDependencyBeenExecuted(dep Dependency) bool {
	_, ok := s.executedDependencies[dep]
	return ok
}

// DependencyExecuted is called by a dependency when it has been asked to be executed.
func (s *ExecutionContext) DependencyExecuted(dep Dependency) {
	s.executedDependencies[dep] = true
}

// MakeExecutionContext creates an execution context.
func MakeExecutionContext() *ExecutionContext {
	return &ExecutionContext{make(map[Dependency]bool)}
}

// Work is the director's interface to the work a Step should perform.
type Work interface {
	Schedule()
}

// Dependency is an abstract dependency of a Step.
// Note that Step implements this interface.
type Dependency interface {
	Execute(executionContext *ExecutionContext) *OngoingStep
}

// OngoingStep is the state of a step when it waits for a dependency or work to complete.
type OngoingStep struct {
	work                Work
	ongoingDependencies []Dependency
}

func (s *OngoingStep) removeOngoingDependency(dependency Dependency) {
	newDependencies := make([]Dependency, 0, len(s.ongoingDependencies))
	for i := range s.ongoingDependencies {
		if s.ongoingDependencies[i] != dependency {
			newDependencies = append(newDependencies, s.ongoingDependencies[i])
		}
	}
	s.ongoingDependencies = newDependencies
}

// DependencyComplete is called when the dependency has completed with some result.
func (s *OngoingStep) DependencyComplete(dependency Dependency, result Result) {
	s.removeOngoingDependency(dependency)
	if len(s.ongoingDependencies) == 0 {
		s.work.Schedule()
	}
}

// Step is the smallest element of a pipeline.
type Step struct {
	work         Work
	dependencies []Dependency
}

// Execute schedules all dependencies, or schedules the work if there are no dependencies.
func (s *Step) Execute(executionContext *ExecutionContext) *OngoingStep {
	executionContext.DependencyExecuted(s)

	if len(s.dependencies) == 0 {
		s.work.Schedule()
		return &OngoingStep{work: s.work, ongoingDependencies: make([]Dependency, 0)}
	}

	for i := range s.dependencies {
		dep := s.dependencies[i]
		if !executionContext.HasDependencyBeenExecuted(dep) {
			dep.Execute(executionContext)
		}
	}
	return &OngoingStep{work: s.work, ongoingDependencies: s.dependencies}
}

// NewStep creates a new Step instance.
func NewStep(work Work, dependencies ...Dependency) *Step {
	return &Step{work: work, dependencies: dependencies}
}
