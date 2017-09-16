package director

// EventPublisher publishes step events by some means.
type EventPublisher interface {
	WorkScheduled(work Work)
	WorkCompleted(work Work)
	DependencyExecuted(dependency Dependency)
	DependencyCompleted(dependency Dependency)
}
