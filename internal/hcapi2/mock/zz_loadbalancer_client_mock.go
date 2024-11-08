// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hetznercloud/cli/internal/hcapi2 (interfaces: LoadBalancerClient)
//
// Generated by this command:
//
//	mockgen -package mock -destination zz_loadbalancer_client_mock.go github.com/hetznercloud/cli/internal/hcapi2 LoadBalancerClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	net "net"
	reflect "reflect"

	hcloud "github.com/hetznercloud/hcloud-go/v2/hcloud"
	gomock "go.uber.org/mock/gomock"
)

// MockLoadBalancerClient is a mock of LoadBalancerClient interface.
type MockLoadBalancerClient struct {
	ctrl     *gomock.Controller
	recorder *MockLoadBalancerClientMockRecorder
	isgomock struct{}
}

// MockLoadBalancerClientMockRecorder is the mock recorder for MockLoadBalancerClient.
type MockLoadBalancerClientMockRecorder struct {
	mock *MockLoadBalancerClient
}

// NewMockLoadBalancerClient creates a new mock instance.
func NewMockLoadBalancerClient(ctrl *gomock.Controller) *MockLoadBalancerClient {
	mock := &MockLoadBalancerClient{ctrl: ctrl}
	mock.recorder = &MockLoadBalancerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoadBalancerClient) EXPECT() *MockLoadBalancerClientMockRecorder {
	return m.recorder
}

// AddIPTarget mocks base method.
func (m *MockLoadBalancerClient) AddIPTarget(ctx context.Context, loadBalancer *hcloud.LoadBalancer, opts hcloud.LoadBalancerAddIPTargetOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIPTarget", ctx, loadBalancer, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddIPTarget indicates an expected call of AddIPTarget.
func (mr *MockLoadBalancerClientMockRecorder) AddIPTarget(ctx, loadBalancer, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIPTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).AddIPTarget), ctx, loadBalancer, opts)
}

// AddLabelSelectorTarget mocks base method.
func (m *MockLoadBalancerClient) AddLabelSelectorTarget(ctx context.Context, loadBalancer *hcloud.LoadBalancer, opts hcloud.LoadBalancerAddLabelSelectorTargetOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLabelSelectorTarget", ctx, loadBalancer, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddLabelSelectorTarget indicates an expected call of AddLabelSelectorTarget.
func (mr *MockLoadBalancerClientMockRecorder) AddLabelSelectorTarget(ctx, loadBalancer, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLabelSelectorTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).AddLabelSelectorTarget), ctx, loadBalancer, opts)
}

// AddServerTarget mocks base method.
func (m *MockLoadBalancerClient) AddServerTarget(ctx context.Context, loadBalancer *hcloud.LoadBalancer, opts hcloud.LoadBalancerAddServerTargetOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddServerTarget", ctx, loadBalancer, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddServerTarget indicates an expected call of AddServerTarget.
func (mr *MockLoadBalancerClientMockRecorder) AddServerTarget(ctx, loadBalancer, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddServerTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).AddServerTarget), ctx, loadBalancer, opts)
}

// AddService mocks base method.
func (m *MockLoadBalancerClient) AddService(ctx context.Context, loadBalancer *hcloud.LoadBalancer, opts hcloud.LoadBalancerAddServiceOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddService", ctx, loadBalancer, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddService indicates an expected call of AddService.
func (mr *MockLoadBalancerClientMockRecorder) AddService(ctx, loadBalancer, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddService", reflect.TypeOf((*MockLoadBalancerClient)(nil).AddService), ctx, loadBalancer, opts)
}

// All mocks base method.
func (m *MockLoadBalancerClient) All(ctx context.Context) ([]*hcloud.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx)
	ret0, _ := ret[0].([]*hcloud.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockLoadBalancerClientMockRecorder) All(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockLoadBalancerClient)(nil).All), ctx)
}

// AllWithOpts mocks base method.
func (m *MockLoadBalancerClient) AllWithOpts(ctx context.Context, opts hcloud.LoadBalancerListOpts) ([]*hcloud.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllWithOpts", ctx, opts)
	ret0, _ := ret[0].([]*hcloud.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllWithOpts indicates an expected call of AllWithOpts.
func (mr *MockLoadBalancerClientMockRecorder) AllWithOpts(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllWithOpts", reflect.TypeOf((*MockLoadBalancerClient)(nil).AllWithOpts), ctx, opts)
}

// AttachToNetwork mocks base method.
func (m *MockLoadBalancerClient) AttachToNetwork(ctx context.Context, loadBalancer *hcloud.LoadBalancer, opts hcloud.LoadBalancerAttachToNetworkOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachToNetwork", ctx, loadBalancer, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AttachToNetwork indicates an expected call of AttachToNetwork.
func (mr *MockLoadBalancerClientMockRecorder) AttachToNetwork(ctx, loadBalancer, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachToNetwork", reflect.TypeOf((*MockLoadBalancerClient)(nil).AttachToNetwork), ctx, loadBalancer, opts)
}

// ChangeAlgorithm mocks base method.
func (m *MockLoadBalancerClient) ChangeAlgorithm(ctx context.Context, loadBalancer *hcloud.LoadBalancer, opts hcloud.LoadBalancerChangeAlgorithmOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeAlgorithm", ctx, loadBalancer, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeAlgorithm indicates an expected call of ChangeAlgorithm.
func (mr *MockLoadBalancerClientMockRecorder) ChangeAlgorithm(ctx, loadBalancer, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeAlgorithm", reflect.TypeOf((*MockLoadBalancerClient)(nil).ChangeAlgorithm), ctx, loadBalancer, opts)
}

// ChangeDNSPtr mocks base method.
func (m *MockLoadBalancerClient) ChangeDNSPtr(ctx context.Context, lb *hcloud.LoadBalancer, ip string, ptr *string) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeDNSPtr", ctx, lb, ip, ptr)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeDNSPtr indicates an expected call of ChangeDNSPtr.
func (mr *MockLoadBalancerClientMockRecorder) ChangeDNSPtr(ctx, lb, ip, ptr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeDNSPtr", reflect.TypeOf((*MockLoadBalancerClient)(nil).ChangeDNSPtr), ctx, lb, ip, ptr)
}

// ChangeProtection mocks base method.
func (m *MockLoadBalancerClient) ChangeProtection(ctx context.Context, loadBalancer *hcloud.LoadBalancer, opts hcloud.LoadBalancerChangeProtectionOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeProtection", ctx, loadBalancer, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeProtection indicates an expected call of ChangeProtection.
func (mr *MockLoadBalancerClientMockRecorder) ChangeProtection(ctx, loadBalancer, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeProtection", reflect.TypeOf((*MockLoadBalancerClient)(nil).ChangeProtection), ctx, loadBalancer, opts)
}

// ChangeType mocks base method.
func (m *MockLoadBalancerClient) ChangeType(ctx context.Context, loadBalancer *hcloud.LoadBalancer, opts hcloud.LoadBalancerChangeTypeOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeType", ctx, loadBalancer, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeType indicates an expected call of ChangeType.
func (mr *MockLoadBalancerClientMockRecorder) ChangeType(ctx, loadBalancer, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeType", reflect.TypeOf((*MockLoadBalancerClient)(nil).ChangeType), ctx, loadBalancer, opts)
}

// Create mocks base method.
func (m *MockLoadBalancerClient) Create(ctx context.Context, opts hcloud.LoadBalancerCreateOpts) (hcloud.LoadBalancerCreateResult, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, opts)
	ret0, _ := ret[0].(hcloud.LoadBalancerCreateResult)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockLoadBalancerClientMockRecorder) Create(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLoadBalancerClient)(nil).Create), ctx, opts)
}

// Delete mocks base method.
func (m *MockLoadBalancerClient) Delete(ctx context.Context, loadBalancer *hcloud.LoadBalancer) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, loadBalancer)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockLoadBalancerClientMockRecorder) Delete(ctx, loadBalancer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLoadBalancerClient)(nil).Delete), ctx, loadBalancer)
}

// DeleteService mocks base method.
func (m *MockLoadBalancerClient) DeleteService(ctx context.Context, loadBalancer *hcloud.LoadBalancer, listenPort int) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", ctx, loadBalancer, listenPort)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockLoadBalancerClientMockRecorder) DeleteService(ctx, loadBalancer, listenPort any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockLoadBalancerClient)(nil).DeleteService), ctx, loadBalancer, listenPort)
}

// DetachFromNetwork mocks base method.
func (m *MockLoadBalancerClient) DetachFromNetwork(ctx context.Context, loadBalancer *hcloud.LoadBalancer, opts hcloud.LoadBalancerDetachFromNetworkOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachFromNetwork", ctx, loadBalancer, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DetachFromNetwork indicates an expected call of DetachFromNetwork.
func (mr *MockLoadBalancerClientMockRecorder) DetachFromNetwork(ctx, loadBalancer, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachFromNetwork", reflect.TypeOf((*MockLoadBalancerClient)(nil).DetachFromNetwork), ctx, loadBalancer, opts)
}

// DisablePublicInterface mocks base method.
func (m *MockLoadBalancerClient) DisablePublicInterface(ctx context.Context, loadBalancer *hcloud.LoadBalancer) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisablePublicInterface", ctx, loadBalancer)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DisablePublicInterface indicates an expected call of DisablePublicInterface.
func (mr *MockLoadBalancerClientMockRecorder) DisablePublicInterface(ctx, loadBalancer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisablePublicInterface", reflect.TypeOf((*MockLoadBalancerClient)(nil).DisablePublicInterface), ctx, loadBalancer)
}

// EnablePublicInterface mocks base method.
func (m *MockLoadBalancerClient) EnablePublicInterface(ctx context.Context, loadBalancer *hcloud.LoadBalancer) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePublicInterface", ctx, loadBalancer)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EnablePublicInterface indicates an expected call of EnablePublicInterface.
func (mr *MockLoadBalancerClientMockRecorder) EnablePublicInterface(ctx, loadBalancer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePublicInterface", reflect.TypeOf((*MockLoadBalancerClient)(nil).EnablePublicInterface), ctx, loadBalancer)
}

// Get mocks base method.
func (m *MockLoadBalancerClient) Get(ctx context.Context, idOrName string) (*hcloud.LoadBalancer, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, idOrName)
	ret0, _ := ret[0].(*hcloud.LoadBalancer)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockLoadBalancerClientMockRecorder) Get(ctx, idOrName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLoadBalancerClient)(nil).Get), ctx, idOrName)
}

// GetByID mocks base method.
func (m *MockLoadBalancerClient) GetByID(ctx context.Context, id int64) (*hcloud.LoadBalancer, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*hcloud.LoadBalancer)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockLoadBalancerClientMockRecorder) GetByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockLoadBalancerClient)(nil).GetByID), ctx, id)
}

// GetByName mocks base method.
func (m *MockLoadBalancerClient) GetByName(ctx context.Context, name string) (*hcloud.LoadBalancer, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(*hcloud.LoadBalancer)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByName indicates an expected call of GetByName.
func (mr *MockLoadBalancerClientMockRecorder) GetByName(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockLoadBalancerClient)(nil).GetByName), ctx, name)
}

// GetMetrics mocks base method.
func (m *MockLoadBalancerClient) GetMetrics(ctx context.Context, lb *hcloud.LoadBalancer, opts hcloud.LoadBalancerGetMetricsOpts) (*hcloud.LoadBalancerMetrics, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetrics", ctx, lb, opts)
	ret0, _ := ret[0].(*hcloud.LoadBalancerMetrics)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMetrics indicates an expected call of GetMetrics.
func (mr *MockLoadBalancerClientMockRecorder) GetMetrics(ctx, lb, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetrics", reflect.TypeOf((*MockLoadBalancerClient)(nil).GetMetrics), ctx, lb, opts)
}

// LabelKeys mocks base method.
func (m *MockLoadBalancerClient) LabelKeys(arg0 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LabelKeys", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// LabelKeys indicates an expected call of LabelKeys.
func (mr *MockLoadBalancerClientMockRecorder) LabelKeys(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LabelKeys", reflect.TypeOf((*MockLoadBalancerClient)(nil).LabelKeys), arg0)
}

// List mocks base method.
func (m *MockLoadBalancerClient) List(ctx context.Context, opts hcloud.LoadBalancerListOpts) ([]*hcloud.LoadBalancer, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, opts)
	ret0, _ := ret[0].([]*hcloud.LoadBalancer)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockLoadBalancerClientMockRecorder) List(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLoadBalancerClient)(nil).List), ctx, opts)
}

// LoadBalancerName mocks base method.
func (m *MockLoadBalancerClient) LoadBalancerName(id int64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadBalancerName", id)
	ret0, _ := ret[0].(string)
	return ret0
}

// LoadBalancerName indicates an expected call of LoadBalancerName.
func (mr *MockLoadBalancerClientMockRecorder) LoadBalancerName(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadBalancerName", reflect.TypeOf((*MockLoadBalancerClient)(nil).LoadBalancerName), id)
}

// Names mocks base method.
func (m *MockLoadBalancerClient) Names() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Names")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Names indicates an expected call of Names.
func (mr *MockLoadBalancerClientMockRecorder) Names() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Names", reflect.TypeOf((*MockLoadBalancerClient)(nil).Names))
}

// RemoveIPTarget mocks base method.
func (m *MockLoadBalancerClient) RemoveIPTarget(ctx context.Context, loadBalancer *hcloud.LoadBalancer, ip net.IP) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIPTarget", ctx, loadBalancer, ip)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RemoveIPTarget indicates an expected call of RemoveIPTarget.
func (mr *MockLoadBalancerClientMockRecorder) RemoveIPTarget(ctx, loadBalancer, ip any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIPTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).RemoveIPTarget), ctx, loadBalancer, ip)
}

// RemoveLabelSelectorTarget mocks base method.
func (m *MockLoadBalancerClient) RemoveLabelSelectorTarget(ctx context.Context, loadBalancer *hcloud.LoadBalancer, labelSelector string) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLabelSelectorTarget", ctx, loadBalancer, labelSelector)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RemoveLabelSelectorTarget indicates an expected call of RemoveLabelSelectorTarget.
func (mr *MockLoadBalancerClientMockRecorder) RemoveLabelSelectorTarget(ctx, loadBalancer, labelSelector any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLabelSelectorTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).RemoveLabelSelectorTarget), ctx, loadBalancer, labelSelector)
}

// RemoveServerTarget mocks base method.
func (m *MockLoadBalancerClient) RemoveServerTarget(ctx context.Context, loadBalancer *hcloud.LoadBalancer, server *hcloud.Server) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveServerTarget", ctx, loadBalancer, server)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RemoveServerTarget indicates an expected call of RemoveServerTarget.
func (mr *MockLoadBalancerClientMockRecorder) RemoveServerTarget(ctx, loadBalancer, server any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveServerTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).RemoveServerTarget), ctx, loadBalancer, server)
}

// Update mocks base method.
func (m *MockLoadBalancerClient) Update(ctx context.Context, loadBalancer *hcloud.LoadBalancer, opts hcloud.LoadBalancerUpdateOpts) (*hcloud.LoadBalancer, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, loadBalancer, opts)
	ret0, _ := ret[0].(*hcloud.LoadBalancer)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockLoadBalancerClientMockRecorder) Update(ctx, loadBalancer, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLoadBalancerClient)(nil).Update), ctx, loadBalancer, opts)
}

// UpdateService mocks base method.
func (m *MockLoadBalancerClient) UpdateService(ctx context.Context, loadBalancer *hcloud.LoadBalancer, listenPort int, opts hcloud.LoadBalancerUpdateServiceOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", ctx, loadBalancer, listenPort, opts)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockLoadBalancerClientMockRecorder) UpdateService(ctx, loadBalancer, listenPort, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockLoadBalancerClient)(nil).UpdateService), ctx, loadBalancer, listenPort, opts)
}
