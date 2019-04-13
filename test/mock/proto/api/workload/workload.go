// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/spire/api/workload (interfaces: SpiffeWorkloadAPIClient,SpiffeWorkloadAPIServer,SpiffeWorkloadAPI_FetchX509SVIDClient,SpiffeWorkloadAPI_FetchX509SVIDServer,SpiffeWorkloadAPI_FetchJWTBundlesServer)

// Package mock_workload is a generated GoMock package.
package mock_workload

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	workload "github.com/spiffe/spire/proto/spire/api/workload"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockSpiffeWorkloadAPIClient is a mock of SpiffeWorkloadAPIClient interface
type MockSpiffeWorkloadAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockSpiffeWorkloadAPIClientMockRecorder
}

// MockSpiffeWorkloadAPIClientMockRecorder is the mock recorder for MockSpiffeWorkloadAPIClient
type MockSpiffeWorkloadAPIClientMockRecorder struct {
	mock *MockSpiffeWorkloadAPIClient
}

// NewMockSpiffeWorkloadAPIClient creates a new mock instance
func NewMockSpiffeWorkloadAPIClient(ctrl *gomock.Controller) *MockSpiffeWorkloadAPIClient {
	mock := &MockSpiffeWorkloadAPIClient{ctrl: ctrl}
	mock.recorder = &MockSpiffeWorkloadAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpiffeWorkloadAPIClient) EXPECT() *MockSpiffeWorkloadAPIClientMockRecorder {
	return m.recorder
}

// FetchJWTBundles mocks base method
func (m *MockSpiffeWorkloadAPIClient) FetchJWTBundles(arg0 context.Context, arg1 *workload.JWTBundlesRequest, arg2 ...grpc.CallOption) (workload.SpiffeWorkloadAPI_FetchJWTBundlesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchJWTBundles", varargs...)
	ret0, _ := ret[0].(workload.SpiffeWorkloadAPI_FetchJWTBundlesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJWTBundles indicates an expected call of FetchJWTBundles
func (mr *MockSpiffeWorkloadAPIClientMockRecorder) FetchJWTBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJWTBundles", reflect.TypeOf((*MockSpiffeWorkloadAPIClient)(nil).FetchJWTBundles), varargs...)
}

// FetchJWTSVID mocks base method
func (m *MockSpiffeWorkloadAPIClient) FetchJWTSVID(arg0 context.Context, arg1 *workload.JWTSVIDRequest, arg2 ...grpc.CallOption) (*workload.JWTSVIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchJWTSVID", varargs...)
	ret0, _ := ret[0].(*workload.JWTSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJWTSVID indicates an expected call of FetchJWTSVID
func (mr *MockSpiffeWorkloadAPIClientMockRecorder) FetchJWTSVID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJWTSVID", reflect.TypeOf((*MockSpiffeWorkloadAPIClient)(nil).FetchJWTSVID), varargs...)
}

// FetchX509SVID mocks base method
func (m *MockSpiffeWorkloadAPIClient) FetchX509SVID(arg0 context.Context, arg1 *workload.X509SVIDRequest, arg2 ...grpc.CallOption) (workload.SpiffeWorkloadAPI_FetchX509SVIDClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchX509SVID", varargs...)
	ret0, _ := ret[0].(workload.SpiffeWorkloadAPI_FetchX509SVIDClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchX509SVID indicates an expected call of FetchX509SVID
func (mr *MockSpiffeWorkloadAPIClientMockRecorder) FetchX509SVID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchX509SVID", reflect.TypeOf((*MockSpiffeWorkloadAPIClient)(nil).FetchX509SVID), varargs...)
}

// ValidateJWTSVID mocks base method
func (m *MockSpiffeWorkloadAPIClient) ValidateJWTSVID(arg0 context.Context, arg1 *workload.ValidateJWTSVIDRequest, arg2 ...grpc.CallOption) (*workload.ValidateJWTSVIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateJWTSVID", varargs...)
	ret0, _ := ret[0].(*workload.ValidateJWTSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateJWTSVID indicates an expected call of ValidateJWTSVID
func (mr *MockSpiffeWorkloadAPIClientMockRecorder) ValidateJWTSVID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateJWTSVID", reflect.TypeOf((*MockSpiffeWorkloadAPIClient)(nil).ValidateJWTSVID), varargs...)
}

// MockSpiffeWorkloadAPIServer is a mock of SpiffeWorkloadAPIServer interface
type MockSpiffeWorkloadAPIServer struct {
	ctrl     *gomock.Controller
	recorder *MockSpiffeWorkloadAPIServerMockRecorder
}

// MockSpiffeWorkloadAPIServerMockRecorder is the mock recorder for MockSpiffeWorkloadAPIServer
type MockSpiffeWorkloadAPIServerMockRecorder struct {
	mock *MockSpiffeWorkloadAPIServer
}

// NewMockSpiffeWorkloadAPIServer creates a new mock instance
func NewMockSpiffeWorkloadAPIServer(ctrl *gomock.Controller) *MockSpiffeWorkloadAPIServer {
	mock := &MockSpiffeWorkloadAPIServer{ctrl: ctrl}
	mock.recorder = &MockSpiffeWorkloadAPIServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpiffeWorkloadAPIServer) EXPECT() *MockSpiffeWorkloadAPIServerMockRecorder {
	return m.recorder
}

// FetchJWTBundles mocks base method
func (m *MockSpiffeWorkloadAPIServer) FetchJWTBundles(arg0 *workload.JWTBundlesRequest, arg1 workload.SpiffeWorkloadAPI_FetchJWTBundlesServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJWTBundles", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchJWTBundles indicates an expected call of FetchJWTBundles
func (mr *MockSpiffeWorkloadAPIServerMockRecorder) FetchJWTBundles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJWTBundles", reflect.TypeOf((*MockSpiffeWorkloadAPIServer)(nil).FetchJWTBundles), arg0, arg1)
}

// FetchJWTSVID mocks base method
func (m *MockSpiffeWorkloadAPIServer) FetchJWTSVID(arg0 context.Context, arg1 *workload.JWTSVIDRequest) (*workload.JWTSVIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJWTSVID", arg0, arg1)
	ret0, _ := ret[0].(*workload.JWTSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJWTSVID indicates an expected call of FetchJWTSVID
func (mr *MockSpiffeWorkloadAPIServerMockRecorder) FetchJWTSVID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJWTSVID", reflect.TypeOf((*MockSpiffeWorkloadAPIServer)(nil).FetchJWTSVID), arg0, arg1)
}

// FetchX509SVID mocks base method
func (m *MockSpiffeWorkloadAPIServer) FetchX509SVID(arg0 *workload.X509SVIDRequest, arg1 workload.SpiffeWorkloadAPI_FetchX509SVIDServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchX509SVID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchX509SVID indicates an expected call of FetchX509SVID
func (mr *MockSpiffeWorkloadAPIServerMockRecorder) FetchX509SVID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchX509SVID", reflect.TypeOf((*MockSpiffeWorkloadAPIServer)(nil).FetchX509SVID), arg0, arg1)
}

// ValidateJWTSVID mocks base method
func (m *MockSpiffeWorkloadAPIServer) ValidateJWTSVID(arg0 context.Context, arg1 *workload.ValidateJWTSVIDRequest) (*workload.ValidateJWTSVIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateJWTSVID", arg0, arg1)
	ret0, _ := ret[0].(*workload.ValidateJWTSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateJWTSVID indicates an expected call of ValidateJWTSVID
func (mr *MockSpiffeWorkloadAPIServerMockRecorder) ValidateJWTSVID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateJWTSVID", reflect.TypeOf((*MockSpiffeWorkloadAPIServer)(nil).ValidateJWTSVID), arg0, arg1)
}

// MockSpiffeWorkloadAPI_FetchX509SVIDClient is a mock of SpiffeWorkloadAPI_FetchX509SVIDClient interface
type MockSpiffeWorkloadAPI_FetchX509SVIDClient struct {
	ctrl     *gomock.Controller
	recorder *MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder
}

// MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder is the mock recorder for MockSpiffeWorkloadAPI_FetchX509SVIDClient
type MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder struct {
	mock *MockSpiffeWorkloadAPI_FetchX509SVIDClient
}

// NewMockSpiffeWorkloadAPI_FetchX509SVIDClient creates a new mock instance
func NewMockSpiffeWorkloadAPI_FetchX509SVIDClient(ctrl *gomock.Controller) *MockSpiffeWorkloadAPI_FetchX509SVIDClient {
	mock := &MockSpiffeWorkloadAPI_FetchX509SVIDClient{ctrl: ctrl}
	mock.recorder = &MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDClient) EXPECT() *MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDClient)(nil).Context))
}

// Header mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDClient)(nil).Header))
}

// Recv mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDClient) Recv() (*workload.X509SVIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*workload.X509SVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDClient)(nil).Trailer))
}

// MockSpiffeWorkloadAPI_FetchX509SVIDServer is a mock of SpiffeWorkloadAPI_FetchX509SVIDServer interface
type MockSpiffeWorkloadAPI_FetchX509SVIDServer struct {
	ctrl     *gomock.Controller
	recorder *MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder
}

// MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder is the mock recorder for MockSpiffeWorkloadAPI_FetchX509SVIDServer
type MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder struct {
	mock *MockSpiffeWorkloadAPI_FetchX509SVIDServer
}

// NewMockSpiffeWorkloadAPI_FetchX509SVIDServer creates a new mock instance
func NewMockSpiffeWorkloadAPI_FetchX509SVIDServer(ctrl *gomock.Controller) *MockSpiffeWorkloadAPI_FetchX509SVIDServer {
	mock := &MockSpiffeWorkloadAPI_FetchX509SVIDServer{ctrl: ctrl}
	mock.recorder = &MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDServer) EXPECT() *MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDServer) Send(arg0 *workload.X509SVIDResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockSpiffeWorkloadAPI_FetchX509SVIDServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockSpiffeWorkloadAPI_FetchX509SVIDServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchX509SVIDServer)(nil).SetTrailer), arg0)
}

// MockSpiffeWorkloadAPI_FetchJWTBundlesServer is a mock of SpiffeWorkloadAPI_FetchJWTBundlesServer interface
type MockSpiffeWorkloadAPI_FetchJWTBundlesServer struct {
	ctrl     *gomock.Controller
	recorder *MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder
}

// MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder is the mock recorder for MockSpiffeWorkloadAPI_FetchJWTBundlesServer
type MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder struct {
	mock *MockSpiffeWorkloadAPI_FetchJWTBundlesServer
}

// NewMockSpiffeWorkloadAPI_FetchJWTBundlesServer creates a new mock instance
func NewMockSpiffeWorkloadAPI_FetchJWTBundlesServer(ctrl *gomock.Controller) *MockSpiffeWorkloadAPI_FetchJWTBundlesServer {
	mock := &MockSpiffeWorkloadAPI_FetchJWTBundlesServer{ctrl: ctrl}
	mock.recorder = &MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpiffeWorkloadAPI_FetchJWTBundlesServer) EXPECT() *MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockSpiffeWorkloadAPI_FetchJWTBundlesServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchJWTBundlesServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockSpiffeWorkloadAPI_FetchJWTBundlesServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchJWTBundlesServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockSpiffeWorkloadAPI_FetchJWTBundlesServer) Send(arg0 *workload.JWTBundlesResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchJWTBundlesServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockSpiffeWorkloadAPI_FetchJWTBundlesServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchJWTBundlesServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockSpiffeWorkloadAPI_FetchJWTBundlesServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchJWTBundlesServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockSpiffeWorkloadAPI_FetchJWTBundlesServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchJWTBundlesServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockSpiffeWorkloadAPI_FetchJWTBundlesServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockSpiffeWorkloadAPI_FetchJWTBundlesServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockSpiffeWorkloadAPI_FetchJWTBundlesServer)(nil).SetTrailer), arg0)
}