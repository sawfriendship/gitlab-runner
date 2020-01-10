// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package volumes

import (
	context "context"

	container "github.com/docker/docker/api/types/container"

	io "io"

	mock "github.com/stretchr/testify/mock"

	network "github.com/docker/docker/api/types/network"

	types "github.com/docker/docker/api/types"
)

// mockContainerClient is an autogenerated mock type for the containerClient type
type mockContainerClient struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *mockContainerClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerAttach provides a mock function with given fields: ctx, _a1, options
func (_m *mockContainerClient) ContainerAttach(ctx context.Context, _a1 string, options types.ContainerAttachOptions) (types.HijackedResponse, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 types.HijackedResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerAttachOptions) types.HijackedResponse); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		r0 = ret.Get(0).(types.HijackedResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerAttachOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerCreate provides a mock function with given fields: ctx, config, hostConfig, networkingConfig, containerName
func (_m *mockContainerClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	ret := _m.Called(ctx, config, hostConfig, networkingConfig, containerName)

	var r0 container.ContainerCreateCreatedBody
	if rf, ok := ret.Get(0).(func(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, string) container.ContainerCreateCreatedBody); ok {
		r0 = rf(ctx, config, hostConfig, networkingConfig, containerName)
	} else {
		r0 = ret.Get(0).(container.ContainerCreateCreatedBody)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, string) error); ok {
		r1 = rf(ctx, config, hostConfig, networkingConfig, containerName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerExecAttach provides a mock function with given fields: ctx, execID, config
func (_m *mockContainerClient) ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error) {
	ret := _m.Called(ctx, execID, config)

	var r0 types.HijackedResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ExecStartCheck) types.HijackedResponse); ok {
		r0 = rf(ctx, execID, config)
	} else {
		r0 = ret.Get(0).(types.HijackedResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ExecStartCheck) error); ok {
		r1 = rf(ctx, execID, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerExecCreate provides a mock function with given fields: ctx, _a1, config
func (_m *mockContainerClient) ContainerExecCreate(ctx context.Context, _a1 string, config types.ExecConfig) (types.IDResponse, error) {
	ret := _m.Called(ctx, _a1, config)

	var r0 types.IDResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ExecConfig) types.IDResponse); ok {
		r0 = rf(ctx, _a1, config)
	} else {
		r0 = ret.Get(0).(types.IDResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ExecConfig) error); ok {
		r1 = rf(ctx, _a1, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerInspect provides a mock function with given fields: ctx, containerID
func (_m *mockContainerClient) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error) {
	ret := _m.Called(ctx, containerID)

	var r0 types.ContainerJSON
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ContainerJSON); ok {
		r0 = rf(ctx, containerID)
	} else {
		r0 = ret.Get(0).(types.ContainerJSON)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, containerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerKill provides a mock function with given fields: ctx, containerID, signal
func (_m *mockContainerClient) ContainerKill(ctx context.Context, containerID string, signal string) error {
	ret := _m.Called(ctx, containerID, signal)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, containerID, signal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerLogs provides a mock function with given fields: ctx, _a1, options
func (_m *mockContainerClient) ContainerLogs(ctx context.Context, _a1 string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerLogsOptions) io.ReadCloser); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerLogsOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerRemove provides a mock function with given fields: ctx, containerID, options
func (_m *mockContainerClient) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	ret := _m.Called(ctx, containerID, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerRemoveOptions) error); ok {
		r0 = rf(ctx, containerID, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerStart provides a mock function with given fields: ctx, containerID, options
func (_m *mockContainerClient) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	ret := _m.Called(ctx, containerID, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerStartOptions) error); ok {
		r0 = rf(ctx, containerID, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImageImportBlocking provides a mock function with given fields: ctx, source, ref, options
func (_m *mockContainerClient) ImageImportBlocking(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) error {
	ret := _m.Called(ctx, source, ref, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ImageImportSource, string, types.ImageImportOptions) error); ok {
		r0 = rf(ctx, source, ref, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImageInspectWithRaw provides a mock function with given fields: ctx, imageID
func (_m *mockContainerClient) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error) {
	ret := _m.Called(ctx, imageID)

	var r0 types.ImageInspect
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ImageInspect); ok {
		r0 = rf(ctx, imageID)
	} else {
		r0 = ret.Get(0).(types.ImageInspect)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, imageID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, imageID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ImagePullBlocking provides a mock function with given fields: ctx, ref, options
func (_m *mockContainerClient) ImagePullBlocking(ctx context.Context, ref string, options types.ImagePullOptions) error {
	ret := _m.Called(ctx, ref, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImagePullOptions) error); ok {
		r0 = rf(ctx, ref, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Info provides a mock function with given fields: ctx
func (_m *mockContainerClient) Info(ctx context.Context) (types.Info, error) {
	ret := _m.Called(ctx)

	var r0 types.Info
	if rf, ok := ret.Get(0).(func(context.Context) types.Info); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.Info)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LabelContainer provides a mock function with given fields: _a0, containerType, otherLabels
func (_m *mockContainerClient) LabelContainer(_a0 *container.Config, containerType string, otherLabels ...string) {
	_va := make([]interface{}, len(otherLabels))
	for _i := range otherLabels {
		_va[_i] = otherLabels[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, containerType)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// NetworkDisconnect provides a mock function with given fields: ctx, networkID, containerID, force
func (_m *mockContainerClient) NetworkDisconnect(ctx context.Context, networkID string, containerID string, force bool) error {
	ret := _m.Called(ctx, networkID, containerID, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, networkID, containerID, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NetworkList provides a mock function with given fields: ctx, options
func (_m *mockContainerClient) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error) {
	ret := _m.Called(ctx, options)

	var r0 []types.NetworkResource
	if rf, ok := ret.Get(0).(func(context.Context, types.NetworkListOptions) []types.NetworkResource); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.NetworkResource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.NetworkListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveContainer provides a mock function with given fields: ctx, id
func (_m *mockContainerClient) RemoveContainer(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitForContainer provides a mock function with given fields: id
func (_m *mockContainerClient) WaitForContainer(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
