// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/cloud/cloud.go

// Package mock_cloud is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	efs "github.com/aws/aws-sdk-go/service/efs"
	gomock "github.com/golang/mock/gomock"
	cloud "github.com/kubernetes-sigs/aws-efs-csi-driver/pkg/cloud"
)

// MockEfs is a mock of Efs interface.
type MockEfs struct {
	ctrl     *gomock.Controller
	recorder *MockEfsMockRecorder
}

// MockEfsMockRecorder is the mock recorder for MockEfs.
type MockEfsMockRecorder struct {
	mock *MockEfs
}

// NewMockEfs creates a new mock instance.
func NewMockEfs(ctrl *gomock.Controller) *MockEfs {
	mock := &MockEfs{ctrl: ctrl}
	mock.recorder = &MockEfsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEfs) EXPECT() *MockEfsMockRecorder {
	return m.recorder
}

// CreateAccessPointWithContext mocks base method.
func (m *MockEfs) CreateAccessPointWithContext(arg0 aws.Context, arg1 *efs.CreateAccessPointInput, arg2 ...request.Option) (*efs.CreateAccessPointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAccessPointWithContext", varargs...)
	ret0, _ := ret[0].(*efs.CreateAccessPointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessPointWithContext indicates an expected call of CreateAccessPointWithContext.
func (mr *MockEfsMockRecorder) CreateAccessPointWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessPointWithContext", reflect.TypeOf((*MockEfs)(nil).CreateAccessPointWithContext), varargs...)
}

// DeleteAccessPointWithContext mocks base method.
func (m *MockEfs) DeleteAccessPointWithContext(arg0 aws.Context, arg1 *efs.DeleteAccessPointInput, arg2 ...request.Option) (*efs.DeleteAccessPointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAccessPointWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DeleteAccessPointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccessPointWithContext indicates an expected call of DeleteAccessPointWithContext.
func (mr *MockEfsMockRecorder) DeleteAccessPointWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessPointWithContext", reflect.TypeOf((*MockEfs)(nil).DeleteAccessPointWithContext), varargs...)
}

// DescribeAccessPointsWithContext mocks base method.
func (m *MockEfs) DescribeAccessPointsWithContext(arg0 aws.Context, arg1 *efs.DescribeAccessPointsInput, arg2 ...request.Option) (*efs.DescribeAccessPointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccessPointsWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DescribeAccessPointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccessPointsWithContext indicates an expected call of DescribeAccessPointsWithContext.
func (mr *MockEfsMockRecorder) DescribeAccessPointsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccessPointsWithContext", reflect.TypeOf((*MockEfs)(nil).DescribeAccessPointsWithContext), varargs...)
}

// DescribeFileSystemsWithContext mocks base method.
func (m *MockEfs) DescribeFileSystemsWithContext(arg0 aws.Context, arg1 *efs.DescribeFileSystemsInput, arg2 ...request.Option) (*efs.DescribeFileSystemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFileSystemsWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DescribeFileSystemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFileSystemsWithContext indicates an expected call of DescribeFileSystemsWithContext.
func (mr *MockEfsMockRecorder) DescribeFileSystemsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystemsWithContext", reflect.TypeOf((*MockEfs)(nil).DescribeFileSystemsWithContext), varargs...)
}

// DescribeMountTargetsWithContext mocks base method.
func (m *MockEfs) DescribeMountTargetsWithContext(arg0 aws.Context, arg1 *efs.DescribeMountTargetsInput, arg2 ...request.Option) (*efs.DescribeMountTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMountTargetsWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DescribeMountTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMountTargetsWithContext indicates an expected call of DescribeMountTargetsWithContext.
func (mr *MockEfsMockRecorder) DescribeMountTargetsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMountTargetsWithContext", reflect.TypeOf((*MockEfs)(nil).DescribeMountTargetsWithContext), varargs...)
}

// MockCloud is a mock of Cloud interface.
type MockCloud struct {
	ctrl     *gomock.Controller
	recorder *MockCloudMockRecorder
}

// MockCloudMockRecorder is the mock recorder for MockCloud.
type MockCloudMockRecorder struct {
	mock *MockCloud
}

// NewMockCloud creates a new mock instance.
func NewMockCloud(ctrl *gomock.Controller) *MockCloud {
	mock := &MockCloud{ctrl: ctrl}
	mock.recorder = &MockCloudMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloud) EXPECT() *MockCloudMockRecorder {
	return m.recorder
}

// CreateAccessPoint mocks base method.
func (m *MockCloud) CreateAccessPoint(ctx context.Context, clientToken string, accessPointOpts *cloud.AccessPointOptions) (*cloud.AccessPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessPoint", ctx, clientToken, accessPointOpts)
	ret0, _ := ret[0].(*cloud.AccessPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessPoint indicates an expected call of CreateAccessPoint.
func (mr *MockCloudMockRecorder) CreateAccessPoint(ctx, clientToken, accessPointOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessPoint", reflect.TypeOf((*MockCloud)(nil).CreateAccessPoint), ctx, clientToken, accessPointOpts)
}

// DeleteAccessPoint mocks base method.
func (m *MockCloud) DeleteAccessPoint(ctx context.Context, accessPointId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessPoint", ctx, accessPointId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessPoint indicates an expected call of DeleteAccessPoint.
func (mr *MockCloudMockRecorder) DeleteAccessPoint(ctx, accessPointId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessPoint", reflect.TypeOf((*MockCloud)(nil).DeleteAccessPoint), ctx, accessPointId)
}

// DescribeAccessPoint mocks base method.
func (m *MockCloud) DescribeAccessPoint(ctx context.Context, accessPointId string) (*cloud.AccessPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAccessPoint", ctx, accessPointId)
	ret0, _ := ret[0].(*cloud.AccessPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccessPoint indicates an expected call of DescribeAccessPoint.
func (mr *MockCloudMockRecorder) DescribeAccessPoint(ctx, accessPointId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccessPoint", reflect.TypeOf((*MockCloud)(nil).DescribeAccessPoint), ctx, accessPointId)
}

// DescribeFileSystem mocks base method.
func (m *MockCloud) DescribeFileSystem(ctx context.Context, fileSystemId string) (*cloud.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeFileSystem", ctx, fileSystemId)
	ret0, _ := ret[0].(*cloud.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFileSystem indicates an expected call of DescribeFileSystem.
func (mr *MockCloudMockRecorder) DescribeFileSystem(ctx, fileSystemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystem", reflect.TypeOf((*MockCloud)(nil).DescribeFileSystem), ctx, fileSystemId)
}

// DescribeMountTargets mocks base method.
func (m *MockCloud) DescribeMountTargets(ctx context.Context, fileSystemId, az string) (*cloud.MountTarget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeMountTargets", ctx, fileSystemId, az)
	ret0, _ := ret[0].(*cloud.MountTarget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMountTargets indicates an expected call of DescribeMountTargets.
func (mr *MockCloudMockRecorder) DescribeMountTargets(ctx, fileSystemId, az interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMountTargets", reflect.TypeOf((*MockCloud)(nil).DescribeMountTargets), ctx, fileSystemId, az)
}

// FindAccessPointByClientToken mocks base method.
func (m *MockCloud) FindAccessPointByClientToken(ctx context.Context, clientToken, fileSystemId string) (*cloud.AccessPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAccessPointByClientToken", ctx, clientToken, fileSystemId)
	ret0, _ := ret[0].(*cloud.AccessPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAccessPointByClientToken indicates an expected call of FindAccessPointByClientToken.
func (mr *MockCloudMockRecorder) FindAccessPointByClientToken(ctx, clientToken, fileSystemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAccessPointByClientToken", reflect.TypeOf((*MockCloud)(nil).FindAccessPointByClientToken), ctx, clientToken, fileSystemId)
}

// GetMetadata mocks base method.
func (m *MockCloud) GetMetadata() cloud.MetadataService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata")
	ret0, _ := ret[0].(cloud.MetadataService)
	return ret0
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockCloudMockRecorder) GetMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockCloud)(nil).GetMetadata))
}

// ListAccessPoints mocks base method.
func (m *MockCloud) ListAccessPoints(ctx context.Context, fileSystemId string) ([]*cloud.AccessPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessPoints", ctx, fileSystemId)
	ret0, _ := ret[0].([]*cloud.AccessPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessPoints indicates an expected call of ListAccessPoints.
func (mr *MockCloudMockRecorder) ListAccessPoints(ctx, fileSystemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessPoints", reflect.TypeOf((*MockCloud)(nil).ListAccessPoints), ctx, fileSystemId)
}
