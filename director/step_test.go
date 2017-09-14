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

	step.Execute()
}

func TestExecuteStep_OneDependency_DependencyIsExecuted(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dependency := makeDependency(mockCtrl)
	step, _ := makeStep(mockCtrl, dependency)

	dependency.EXPECT().Execute()

	step.Execute()
}

func TestExecuteStep_TwoDependencies_BothDependenciesExecuted(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dependency1 := makeDependency(mockCtrl)
	dependency2 := makeDependency(mockCtrl)
	step, _ := makeStep(mockCtrl, dependency1, dependency2)

	dependency1.EXPECT().Execute()
	dependency2.EXPECT().Execute()

	step.Execute()
}

func TestDependencyCompletes_SingleDependencyIsSuccessful_WorkIsScheduled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dependency := makeDependency(mockCtrl)
	step, work := makeStep(mockCtrl, dependency)

	dependency.EXPECT().Execute().AnyTimes()
	work.EXPECT().Schedule()

	ongoingStep := step.Execute()
	ongoingStep.DependencyComplete(dependency, director.SuccessResult())
}
