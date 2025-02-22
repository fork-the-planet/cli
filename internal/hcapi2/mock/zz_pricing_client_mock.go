// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hetznercloud/cli/internal/hcapi2 (interfaces: PricingClient)
//
// Generated by this command:
//
//	mockgen -package mock -destination zz_pricing_client_mock.go github.com/hetznercloud/cli/internal/hcapi2 PricingClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	hcloud "github.com/hetznercloud/hcloud-go/v2/hcloud"
	gomock "go.uber.org/mock/gomock"
)

// MockPricingClient is a mock of PricingClient interface.
type MockPricingClient struct {
	ctrl     *gomock.Controller
	recorder *MockPricingClientMockRecorder
	isgomock struct{}
}

// MockPricingClientMockRecorder is the mock recorder for MockPricingClient.
type MockPricingClientMockRecorder struct {
	mock *MockPricingClient
}

// NewMockPricingClient creates a new mock instance.
func NewMockPricingClient(ctrl *gomock.Controller) *MockPricingClient {
	mock := &MockPricingClient{ctrl: ctrl}
	mock.recorder = &MockPricingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPricingClient) EXPECT() *MockPricingClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockPricingClient) Get(ctx context.Context) (hcloud.Pricing, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(hcloud.Pricing)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockPricingClientMockRecorder) Get(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPricingClient)(nil).Get), ctx)
}
