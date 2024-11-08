// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hetznercloud/cli/internal/hcapi2 (interfaces: ImageClient)
//
// Generated by this command:
//
//	mockgen -package mock -destination zz_image_client_mock.go github.com/hetznercloud/cli/internal/hcapi2 ImageClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	hcloud "github.com/hetznercloud/hcloud-go/v2/hcloud"
	gomock "go.uber.org/mock/gomock"
)

// MockImageClient is a mock of ImageClient interface.
type MockImageClient struct {
	ctrl     *gomock.Controller
	recorder *MockImageClientMockRecorder
	isgomock struct{}
}

// MockImageClientMockRecorder is the mock recorder for MockImageClient.
type MockImageClientMockRecorder struct {
	mock *MockImageClient
}

// NewMockImageClient creates a new mock instance.
func NewMockImageClient(ctrl *gomock.Controller) *MockImageClient {
	mock := &MockImageClient{ctrl: ctrl}
	mock.recorder = &MockImageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageClient) EXPECT() *MockImageClientMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockImageClient) All(ctx context.Context) ([]*hcloud.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx)
	ret0, _ := ret[0].([]*hcloud.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockImageClientMockRecorder) All(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockImageClient)(nil).All), ctx)
}

// AllWithOpts mocks base method.
func (m *MockImageClient) AllWithOpts(ctx context.Context, opts hcloud.ImageListOpts) ([]*hcloud.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllWithOpts", ctx, opts)
	ret0, _ := ret[0].([]*hcloud.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllWithOpts indicates an expected call of AllWithOpts.
func (mr *MockImageClientMockRecorder) AllWithOpts(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllWithOpts", reflect.TypeOf((*MockImageClient)(nil).AllWithOpts), ctx, opts)
}

// ChangeProtection mocks base method.
func (m *MockImageClient) ChangeProtection(ctx context.Context, image *hcloud.Image, opts hcloud.ImageChangeProtectionOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeProtection", ctx, image, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeProtection indicates an expected call of ChangeProtection.
func (mr *MockImageClientMockRecorder) ChangeProtection(ctx, image, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeProtection", reflect.TypeOf((*MockImageClient)(nil).ChangeProtection), ctx, image, opts)
}

// Delete mocks base method.
func (m *MockImageClient) Delete(ctx context.Context, image *hcloud.Image) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, image)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockImageClientMockRecorder) Delete(ctx, image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockImageClient)(nil).Delete), ctx, image)
}

// Get mocks base method.
func (m *MockImageClient) Get(ctx context.Context, idOrName string) (*hcloud.Image, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, idOrName)
	ret0, _ := ret[0].(*hcloud.Image)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockImageClientMockRecorder) Get(ctx, idOrName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockImageClient)(nil).Get), ctx, idOrName)
}

// GetByID mocks base method.
func (m *MockImageClient) GetByID(ctx context.Context, id int64) (*hcloud.Image, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*hcloud.Image)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockImageClientMockRecorder) GetByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockImageClient)(nil).GetByID), ctx, id)
}

// GetByName mocks base method.
func (m *MockImageClient) GetByName(ctx context.Context, name string) (*hcloud.Image, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(*hcloud.Image)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByName indicates an expected call of GetByName.
func (mr *MockImageClientMockRecorder) GetByName(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockImageClient)(nil).GetByName), ctx, name)
}

// GetByNameAndArchitecture mocks base method.
func (m *MockImageClient) GetByNameAndArchitecture(ctx context.Context, name string, architecture hcloud.Architecture) (*hcloud.Image, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNameAndArchitecture", ctx, name, architecture)
	ret0, _ := ret[0].(*hcloud.Image)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByNameAndArchitecture indicates an expected call of GetByNameAndArchitecture.
func (mr *MockImageClientMockRecorder) GetByNameAndArchitecture(ctx, name, architecture any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNameAndArchitecture", reflect.TypeOf((*MockImageClient)(nil).GetByNameAndArchitecture), ctx, name, architecture)
}

// GetForArchitecture mocks base method.
func (m *MockImageClient) GetForArchitecture(ctx context.Context, idOrName string, architecture hcloud.Architecture) (*hcloud.Image, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForArchitecture", ctx, idOrName, architecture)
	ret0, _ := ret[0].(*hcloud.Image)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetForArchitecture indicates an expected call of GetForArchitecture.
func (mr *MockImageClientMockRecorder) GetForArchitecture(ctx, idOrName, architecture any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForArchitecture", reflect.TypeOf((*MockImageClient)(nil).GetForArchitecture), ctx, idOrName, architecture)
}

// LabelKeys mocks base method.
func (m *MockImageClient) LabelKeys(arg0 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LabelKeys", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// LabelKeys indicates an expected call of LabelKeys.
func (mr *MockImageClientMockRecorder) LabelKeys(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LabelKeys", reflect.TypeOf((*MockImageClient)(nil).LabelKeys), arg0)
}

// List mocks base method.
func (m *MockImageClient) List(ctx context.Context, opts hcloud.ImageListOpts) ([]*hcloud.Image, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, opts)
	ret0, _ := ret[0].([]*hcloud.Image)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockImageClientMockRecorder) List(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockImageClient)(nil).List), ctx, opts)
}

// Names mocks base method.
func (m *MockImageClient) Names() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Names")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Names indicates an expected call of Names.
func (mr *MockImageClientMockRecorder) Names() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Names", reflect.TypeOf((*MockImageClient)(nil).Names))
}

// Update mocks base method.
func (m *MockImageClient) Update(ctx context.Context, image *hcloud.Image, opts hcloud.ImageUpdateOpts) (*hcloud.Image, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, image, opts)
	ret0, _ := ret[0].(*hcloud.Image)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockImageClientMockRecorder) Update(ctx, image, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockImageClient)(nil).Update), ctx, image, opts)
}
