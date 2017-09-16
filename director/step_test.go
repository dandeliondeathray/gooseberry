package director_test

import (
	"testing"

	"github.com/dandeliondeathray/gooseberry/director"
	"github.com/dandeliondeathray/gooseberry/director/director_mock"
	"github.com/golang/mock/gomock"
)

func makeStep(mockCtrl *gomock.Controller, deps ...director.Dependency) (*director.Step, *director_mock.MockWork) {
	mockWork := director_mock.NewMockWork(mockCtrl)
	return director.NewStep(mockWork, deps...), mockWork
}

func makeDependency(mockCtrl *gomock.Controller) *director_mock.MockDependency {
	return director_mock.NewMockDependency(mockCtrl)
}

func TestExecuteStep_NoDependencies_WorkIsScheduled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	step, work := makeStep(mockCtrl)

	work.EXPECT().Schedule()

	step.Execute(director.MakeExecutionContext())
}

func TestExecuteStep_OneDependency_DependencyIsExecuted(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dependency := makeDependency(mockCtrl)
	step, _ := makeStep(mockCtrl, dependency)

	dependency.EXPECT().Execute(gomock.Any())

	step.Execute(director.MakeExecutionContext())
}

func TestExecuteStep_TwoDependencies_BothDependenciesExecuted(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dependency1 := makeDependency(mockCtrl)
	dependency2 := makeDependency(mockCtrl)
	step, _ := makeStep(mockCtrl, dependency1, dependency2)

	dependency1.EXPECT().Execute(gomock.Any())
	dependency2.EXPECT().Execute(gomock.Any())

	step.Execute(director.MakeExecutionContext())
}

func TestDependencyCompletes_SingleDependencyIsSuccessful_WorkIsScheduled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dependency := makeDependency(mockCtrl)
	step, work := makeStep(mockCtrl, dependency)

	dependency.EXPECT().Execute(gomock.Any()).AnyTimes()
	work.EXPECT().Schedule()

	ongoingStep := step.Execute(director.MakeExecutionContext())
	ongoingStep.DependencyComplete(dependency, director.SuccessResult())
}

func TestDependencyCompletes_NotAllDepenciesComplete_WorkIsNotScheduled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dependency1 := makeDependency(mockCtrl)
	dependency2 := makeDependency(mockCtrl)
	step, _ := makeStep(mockCtrl, dependency1, dependency2)

	dependency1.EXPECT().Execute(gomock.Any()).AnyTimes()
	dependency2.EXPECT().Execute(gomock.Any()).AnyTimes()

	ongoingStep := step.Execute(director.MakeExecutionContext())
	ongoingStep.DependencyComplete(dependency1, director.SuccessResult())
}

func TestDependencyCompletes_AllDepenciesComplete_WorkIsScheduled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dependency1 := makeDependency(mockCtrl)
	dependency2 := makeDependency(mockCtrl)
	step, work := makeStep(mockCtrl, dependency1, dependency2)

	dependency1.EXPECT().Execute(gomock.Any()).AnyTimes()
	dependency2.EXPECT().Execute(gomock.Any()).AnyTimes()
	work.EXPECT().Schedule()

	ongoingStep := step.Execute(director.MakeExecutionContext())
	ongoingStep.DependencyComplete(dependency1, director.SuccessResult())
	ongoingStep.DependencyComplete(dependency2, director.SuccessResult())
}

func TestDiamondDependency_TwoStepsHaveSameDependency_DependencyIsExecutedOnlyOne(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// If we execute step D, then ensure that the work for step A is only scheduled once.
	//            / --- Step B <--- \
	// Step A  <--                    --- Step D
	//            \ --- Step C <--- /

	stepA, workA := makeStep(mockCtrl)
	stepB, _ := makeStep(mockCtrl, stepA)
	stepC, _ := makeStep(mockCtrl, stepA)
	stepD, _ := makeStep(mockCtrl, stepB, stepC)

	workA.EXPECT().Schedule()

	stepD.Execute(director.MakeExecutionContext())
}
