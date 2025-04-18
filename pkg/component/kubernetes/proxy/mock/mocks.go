// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/component/kubernetes/proxy (interfaces: Interface)
//
// Generated by this command:
//
//	mockgen -package mock -destination=mocks.go github.com/gardener/gardener/pkg/component/kubernetes/proxy Interface
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	net "net"
	reflect "reflect"

	proxy "github.com/gardener/gardener/pkg/component/kubernetes/proxy"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// DeleteStaleResources mocks base method.
func (m *MockInterface) DeleteStaleResources(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStaleResources", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStaleResources indicates an expected call of DeleteStaleResources.
func (mr *MockInterfaceMockRecorder) DeleteStaleResources(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStaleResources", reflect.TypeOf((*MockInterface)(nil).DeleteStaleResources), arg0)
}

// Deploy mocks base method.
func (m *MockInterface) Deploy(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockInterfaceMockRecorder) Deploy(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockInterface)(nil).Deploy), ctx)
}

// Destroy mocks base method.
func (m *MockInterface) Destroy(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockInterfaceMockRecorder) Destroy(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockInterface)(nil).Destroy), ctx)
}

// SetKubeconfig mocks base method.
func (m *MockInterface) SetKubeconfig(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKubeconfig", arg0)
}

// SetKubeconfig indicates an expected call of SetKubeconfig.
func (mr *MockInterfaceMockRecorder) SetKubeconfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKubeconfig", reflect.TypeOf((*MockInterface)(nil).SetKubeconfig), arg0)
}

// SetPodNetworkCIDRs mocks base method.
func (m *MockInterface) SetPodNetworkCIDRs(arg0 []net.IPNet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPodNetworkCIDRs", arg0)
}

// SetPodNetworkCIDRs indicates an expected call of SetPodNetworkCIDRs.
func (mr *MockInterfaceMockRecorder) SetPodNetworkCIDRs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPodNetworkCIDRs", reflect.TypeOf((*MockInterface)(nil).SetPodNetworkCIDRs), arg0)
}

// SetWorkerPools mocks base method.
func (m *MockInterface) SetWorkerPools(arg0 []proxy.WorkerPool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetWorkerPools", arg0)
}

// SetWorkerPools indicates an expected call of SetWorkerPools.
func (mr *MockInterfaceMockRecorder) SetWorkerPools(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorkerPools", reflect.TypeOf((*MockInterface)(nil).SetWorkerPools), arg0)
}

// Wait mocks base method.
func (m *MockInterface) Wait(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockInterfaceMockRecorder) Wait(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockInterface)(nil).Wait), ctx)
}

// WaitCleanup mocks base method.
func (m *MockInterface) WaitCleanup(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitCleanup", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitCleanup indicates an expected call of WaitCleanup.
func (mr *MockInterfaceMockRecorder) WaitCleanup(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitCleanup", reflect.TypeOf((*MockInterface)(nil).WaitCleanup), ctx)
}

// WaitCleanupStaleResources mocks base method.
func (m *MockInterface) WaitCleanupStaleResources(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitCleanupStaleResources", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitCleanupStaleResources indicates an expected call of WaitCleanupStaleResources.
func (mr *MockInterfaceMockRecorder) WaitCleanupStaleResources(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitCleanupStaleResources", reflect.TypeOf((*MockInterface)(nil).WaitCleanupStaleResources), arg0)
}
