// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/topic/topic_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/mfsyahrz/bareksaproject/internal/domain/entity"
)

// MockTopicService is a mock of TopicService interface.
type MockTopicService struct {
	ctrl     *gomock.Controller
	recorder *MockTopicServiceMockRecorder
}

// MockTopicServiceMockRecorder is the mock recorder for MockTopicService.
type MockTopicServiceMockRecorder struct {
	mock *MockTopicService
}

// NewMockTopicService creates a new mock instance.
func NewMockTopicService(ctrl *gomock.Controller) *MockTopicService {
	mock := &MockTopicService{ctrl: ctrl}
	mock.recorder = &MockTopicServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopicService) EXPECT() *MockTopicServiceMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockTopicService) FindAll(ctx context.Context) ([]entity.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]entity.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockTopicServiceMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTopicService)(nil).FindAll), ctx)
}
