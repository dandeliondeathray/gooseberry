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
