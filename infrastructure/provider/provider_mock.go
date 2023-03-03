package provider

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for Mock
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// GetFilms mocks base method.
func (m *MockProvider) GetFilms(ctx context.Context, page int) (*GetFilmsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilms", ctx, page)
	ret0, _ := ret[0].(*GetFilmsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilms indicates an expected call of GetFilms.
func (mr *MockProviderMockRecorder) GetFilms(ctx, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilms", reflect.TypeOf((*MockProvider)(nil).GetFilms), ctx, page)
}

// GetPlanets mocks base method.
func (m *MockProvider) GetPlanets(ctx context.Context, page int) (*GetPlanetsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlanets", ctx, page)
	ret0, _ := ret[0].(*GetPlanetsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlanets indicates an expected call of GetPlanets.
func (mr *MockProviderMockRecorder) GetPlanets(ctx, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlanets", reflect.TypeOf((*MockProvider)(nil).GetPlanets), ctx, page)
}
