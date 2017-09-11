package director_test

import (
	"testing"

	"github.com/dandeliondeathray/gooseberry/director"
	"github.com/dandeliondeathray/gooseberry/director/director_mock"
	"github.com/golang/mock/gomock"
)

func makeStep(mockCtrl *gomock.Controller) (*director.Step, *director_mock.MockWork) {
	mockWork := director_mock.NewMockWork(mockCtrl)
	return director.NewStep(mockWork), mockWork
}

func TestExecuteStep_NoDependencies_WorkIsScheduled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	step, work := makeStep(mockCtrl)

	work.EXPECT().Schedule()

	step.Execute()
}
