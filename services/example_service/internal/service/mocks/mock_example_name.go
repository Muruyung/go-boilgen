// Code generated by MockGen. DO NOT EDIT.
// Source: /home/robi/ngoding/pribadi/go-boilgen/services/example_service/domain/usecase/example_name.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	usecase "github.com/Muruyung/go-boilgen/services/example_service/domain/usecase"
	goutils "github.com/Muruyung/go-utilities"
	gomock "github.com/golang/mock/gomock"
)

// MockExampleNameUseCase is a mock of ExampleNameUseCase interface.
type MockExampleNameUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockExampleNameUseCaseMockRecorder
}

// MockExampleNameUseCaseMockRecorder is the mock recorder for MockExampleNameUseCase.
type MockExampleNameUseCaseMockRecorder struct {
	mock *MockExampleNameUseCase
}

// NewMockExampleNameUseCase creates a new mock instance.
func NewMockExampleNameUseCase(ctrl *gomock.Controller) *MockExampleNameUseCase {
	mock := &MockExampleNameUseCase{ctrl: ctrl}
	mock.recorder = &MockExampleNameUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExampleNameUseCase) EXPECT() *MockExampleNameUseCaseMockRecorder {
	return m.recorder
}

// CreateExampleName mocks base method.
func (m *MockExampleNameUseCase) CreateExampleName(ctx context.Context, dto usecase.DTOExampleName) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExampleName", ctx, dto)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExampleName indicates an expected call of CreateExampleName.
func (mr *MockExampleNameUseCaseMockRecorder) CreateExampleName(ctx, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExampleName", reflect.TypeOf((*MockExampleNameUseCase)(nil).CreateExampleName), ctx, dto)
}

// DeleteExampleName mocks base method.
func (m *MockExampleNameUseCase) DeleteExampleName(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExampleName", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExampleName indicates an expected call of DeleteExampleName.
func (mr *MockExampleNameUseCaseMockRecorder) DeleteExampleName(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExampleName", reflect.TypeOf((*MockExampleNameUseCase)(nil).DeleteExampleName), ctx, id)
}

// GetExampleNameByID mocks base method.
func (m *MockExampleNameUseCase) GetExampleNameByID(ctx context.Context, id int) (*entity.ExampleName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExampleNameByID", ctx, id)
	ret0, _ := ret[0].(*entity.ExampleName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExampleNameByID indicates an expected call of GetExampleNameByID.
func (mr *MockExampleNameUseCaseMockRecorder) GetExampleNameByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExampleNameByID", reflect.TypeOf((*MockExampleNameUseCase)(nil).GetExampleNameByID), ctx, id)
}

// GetListExampleName mocks base method.
func (m *MockExampleNameUseCase) GetListExampleName(ctx context.Context, request *goutils.RequestOption) ([]*entity.ExampleName, *goutils.MetaResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListExampleName", ctx, request)
	ret0, _ := ret[0].([]*entity.ExampleName)
	ret1, _ := ret[1].(*goutils.MetaResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetListExampleName indicates an expected call of GetListExampleName.
func (mr *MockExampleNameUseCaseMockRecorder) GetListExampleName(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListExampleName", reflect.TypeOf((*MockExampleNameUseCase)(nil).GetListExampleName), ctx, request)
}

// UpdateExampleName mocks base method.
func (m *MockExampleNameUseCase) UpdateExampleName(ctx context.Context, id int, dto usecase.DTOExampleName) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExampleName", ctx, id, dto)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExampleName indicates an expected call of UpdateExampleName.
func (mr *MockExampleNameUseCaseMockRecorder) UpdateExampleName(ctx, id, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExampleName", reflect.TypeOf((*MockExampleNameUseCase)(nil).UpdateExampleName), ctx, id, dto)
}
