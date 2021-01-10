// Code generated by MockGen. DO NOT EDIT.
// Source: bounty.go

// Package bugcrowd is a generated GoMock package.
package bugcrowd

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBountyAPI is a mock of BountyAPI interface
type MockBountyAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBountyAPIMockRecorder
}

// MockBountyAPIMockRecorder is the mock recorder for MockBountyAPI
type MockBountyAPIMockRecorder struct {
	mock *MockBountyAPI
}

// NewMockBountyAPI creates a new mock instance
func NewMockBountyAPI(ctrl *gomock.Controller) *MockBountyAPI {
	mock := &MockBountyAPI{ctrl: ctrl}
	mock.recorder = &MockBountyAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBountyAPI) EXPECT() *MockBountyAPIMockRecorder {
	return m.recorder
}

// GetBounties mocks base method
func (m *MockBountyAPI) GetBounties(requestConfig GetBountiesRequestConfig) (GetBountiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBounties", requestConfig)
	ret0, _ := ret[0].(GetBountiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBounties indicates an expected call of GetBounties
func (mr *MockBountyAPIMockRecorder) GetBounties(requestConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBounties", reflect.TypeOf((*MockBountyAPI)(nil).GetBounties), requestConfig)
}

// RetrieveBounty mocks base method
func (m *MockBountyAPI) RetrieveBounty(uuid string) (RetrieveBountyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveBounty", uuid)
	ret0, _ := ret[0].(RetrieveBountyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveBounty indicates an expected call of RetrieveBounty
func (mr *MockBountyAPIMockRecorder) RetrieveBounty(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveBounty", reflect.TypeOf((*MockBountyAPI)(nil).RetrieveBounty), uuid)
}
