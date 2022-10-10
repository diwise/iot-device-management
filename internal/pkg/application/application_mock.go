// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package application

import (
	"context"
	"sync"
	"time"
)

// Ensure, that DeviceManagementMock does implement DeviceManagement.
// If this is not the case, regenerate this file with moq.
var _ DeviceManagement = &DeviceManagementMock{}

// DeviceManagementMock is a mock implementation of DeviceManagement.
//
// 	func TestSomethingThatUsesDeviceManagement(t *testing.T) {
//
// 		// make and configure a mocked DeviceManagement
// 		mockedDeviceManagement := &DeviceManagementMock{
// 			CreateDeviceFunc: func(contextMoqParam context.Context, device Device) error {
// 				panic("mock out the CreateDevice method")
// 			},
// 			GetDeviceFunc: func(contextMoqParam context.Context, s string) (Device, error) {
// 				panic("mock out the GetDevice method")
// 			},
// 			GetDeviceFromEUIFunc: func(contextMoqParam context.Context, s string) (Device, error) {
// 				panic("mock out the GetDeviceFromEUI method")
// 			},
// 			ListAllDevicesFunc: func(contextMoqParam context.Context, strings []string) ([]Device, error) {
// 				panic("mock out the ListAllDevices method")
// 			},
// 			ListEnvironmentsFunc: func(contextMoqParam context.Context) ([]Environment, error) {
// 				panic("mock out the ListEnvironments method")
// 			},
// 			NotifyStatusFunc: func(ctx context.Context, deviceID string, message Status) error {
// 				panic("mock out the NotifyStatus method")
// 			},
// 			SetStatusIfChangedFunc: func(ctx context.Context, deviceID string, message Status) error {
// 				panic("mock out the SetStatusIfChanged method")
// 			},
// 			UpdateDeviceFunc: func(ctx context.Context, deviceID string, fields map[string]interface{}) (Device, error) {
// 				panic("mock out the UpdateDevice method")
// 			},
// 			UpdateLastObservedOnDeviceFunc: func(ctx context.Context, deviceID string, timestamp time.Time) error {
// 				panic("mock out the UpdateLastObservedOnDevice method")
// 			},
// 		}
//
// 		// use mockedDeviceManagement in code that requires DeviceManagement
// 		// and then make assertions.
//
// 	}
type DeviceManagementMock struct {
	// CreateDeviceFunc mocks the CreateDevice method.
	CreateDeviceFunc func(contextMoqParam context.Context, device Device) error

	// GetDeviceFunc mocks the GetDevice method.
	GetDeviceFunc func(contextMoqParam context.Context, s string) (Device, error)

	// GetDeviceFromEUIFunc mocks the GetDeviceFromEUI method.
	GetDeviceFromEUIFunc func(contextMoqParam context.Context, s string) (Device, error)

	// ListAllDevicesFunc mocks the ListAllDevices method.
	ListAllDevicesFunc func(contextMoqParam context.Context, strings []string) ([]Device, error)

	// ListEnvironmentsFunc mocks the ListEnvironments method.
	ListEnvironmentsFunc func(contextMoqParam context.Context) ([]Environment, error)

	// NotifyStatusFunc mocks the NotifyStatus method.
	NotifyStatusFunc func(ctx context.Context, deviceID string, message Status) error

	// SetStatusIfChangedFunc mocks the SetStatusIfChanged method.
	SetStatusIfChangedFunc func(ctx context.Context, deviceID string, message Status) error

	// UpdateDeviceFunc mocks the UpdateDevice method.
	UpdateDeviceFunc func(ctx context.Context, deviceID string, fields map[string]interface{}) (Device, error)

	// UpdateLastObservedOnDeviceFunc mocks the UpdateLastObservedOnDevice method.
	UpdateLastObservedOnDeviceFunc func(ctx context.Context, deviceID string, timestamp time.Time) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateDevice holds details about calls to the CreateDevice method.
		CreateDevice []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Device is the device argument value.
			Device Device
		}
		// GetDevice holds details about calls to the GetDevice method.
		GetDevice []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// S is the s argument value.
			S string
		}
		// GetDeviceFromEUI holds details about calls to the GetDeviceFromEUI method.
		GetDeviceFromEUI []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// S is the s argument value.
			S string
		}
		// ListAllDevices holds details about calls to the ListAllDevices method.
		ListAllDevices []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Strings is the strings argument value.
			Strings []string
		}
		// ListEnvironments holds details about calls to the ListEnvironments method.
		ListEnvironments []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
		// NotifyStatus holds details about calls to the NotifyStatus method.
		NotifyStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeviceID is the deviceID argument value.
			DeviceID string
			// Message is the message argument value.
			Message Status
		}
		// SetStatusIfChanged holds details about calls to the SetStatusIfChanged method.
		SetStatusIfChanged []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeviceID is the deviceID argument value.
			DeviceID string
			// Message is the message argument value.
			Message Status
		}
		// UpdateDevice holds details about calls to the UpdateDevice method.
		UpdateDevice []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeviceID is the deviceID argument value.
			DeviceID string
			// Fields is the fields argument value.
			Fields map[string]interface{}
		}
		// UpdateLastObservedOnDevice holds details about calls to the UpdateLastObservedOnDevice method.
		UpdateLastObservedOnDevice []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeviceID is the deviceID argument value.
			DeviceID string
			// Timestamp is the timestamp argument value.
			Timestamp time.Time
		}
	}
	lockCreateDevice               sync.RWMutex
	lockGetDevice                  sync.RWMutex
	lockGetDeviceFromEUI           sync.RWMutex
	lockListAllDevices             sync.RWMutex
	lockListEnvironments           sync.RWMutex
	lockNotifyStatus               sync.RWMutex
	lockSetStatusIfChanged         sync.RWMutex
	lockUpdateDevice               sync.RWMutex
	lockUpdateLastObservedOnDevice sync.RWMutex
}

// CreateDevice calls CreateDeviceFunc.
func (mock *DeviceManagementMock) CreateDevice(contextMoqParam context.Context, device Device) error {
	if mock.CreateDeviceFunc == nil {
		panic("DeviceManagementMock.CreateDeviceFunc: method is nil but DeviceManagement.CreateDevice was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Device          Device
	}{
		ContextMoqParam: contextMoqParam,
		Device:          device,
	}
	mock.lockCreateDevice.Lock()
	mock.calls.CreateDevice = append(mock.calls.CreateDevice, callInfo)
	mock.lockCreateDevice.Unlock()
	return mock.CreateDeviceFunc(contextMoqParam, device)
}

// CreateDeviceCalls gets all the calls that were made to CreateDevice.
// Check the length with:
//     len(mockedDeviceManagement.CreateDeviceCalls())
func (mock *DeviceManagementMock) CreateDeviceCalls() []struct {
	ContextMoqParam context.Context
	Device          Device
} {
	var calls []struct {
		ContextMoqParam context.Context
		Device          Device
	}
	mock.lockCreateDevice.RLock()
	calls = mock.calls.CreateDevice
	mock.lockCreateDevice.RUnlock()
	return calls
}

// GetDevice calls GetDeviceFunc.
func (mock *DeviceManagementMock) GetDevice(contextMoqParam context.Context, s string) (Device, error) {
	if mock.GetDeviceFunc == nil {
		panic("DeviceManagementMock.GetDeviceFunc: method is nil but DeviceManagement.GetDevice was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		S               string
	}{
		ContextMoqParam: contextMoqParam,
		S:               s,
	}
	mock.lockGetDevice.Lock()
	mock.calls.GetDevice = append(mock.calls.GetDevice, callInfo)
	mock.lockGetDevice.Unlock()
	return mock.GetDeviceFunc(contextMoqParam, s)
}

// GetDeviceCalls gets all the calls that were made to GetDevice.
// Check the length with:
//     len(mockedDeviceManagement.GetDeviceCalls())
func (mock *DeviceManagementMock) GetDeviceCalls() []struct {
	ContextMoqParam context.Context
	S               string
} {
	var calls []struct {
		ContextMoqParam context.Context
		S               string
	}
	mock.lockGetDevice.RLock()
	calls = mock.calls.GetDevice
	mock.lockGetDevice.RUnlock()
	return calls
}

// GetDeviceFromEUI calls GetDeviceFromEUIFunc.
func (mock *DeviceManagementMock) GetDeviceFromEUI(contextMoqParam context.Context, s string) (Device, error) {
	if mock.GetDeviceFromEUIFunc == nil {
		panic("DeviceManagementMock.GetDeviceFromEUIFunc: method is nil but DeviceManagement.GetDeviceFromEUI was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		S               string
	}{
		ContextMoqParam: contextMoqParam,
		S:               s,
	}
	mock.lockGetDeviceFromEUI.Lock()
	mock.calls.GetDeviceFromEUI = append(mock.calls.GetDeviceFromEUI, callInfo)
	mock.lockGetDeviceFromEUI.Unlock()
	return mock.GetDeviceFromEUIFunc(contextMoqParam, s)
}

// GetDeviceFromEUICalls gets all the calls that were made to GetDeviceFromEUI.
// Check the length with:
//     len(mockedDeviceManagement.GetDeviceFromEUICalls())
func (mock *DeviceManagementMock) GetDeviceFromEUICalls() []struct {
	ContextMoqParam context.Context
	S               string
} {
	var calls []struct {
		ContextMoqParam context.Context
		S               string
	}
	mock.lockGetDeviceFromEUI.RLock()
	calls = mock.calls.GetDeviceFromEUI
	mock.lockGetDeviceFromEUI.RUnlock()
	return calls
}

// ListAllDevices calls ListAllDevicesFunc.
func (mock *DeviceManagementMock) ListAllDevices(contextMoqParam context.Context, strings []string) ([]Device, error) {
	if mock.ListAllDevicesFunc == nil {
		panic("DeviceManagementMock.ListAllDevicesFunc: method is nil but DeviceManagement.ListAllDevices was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Strings         []string
	}{
		ContextMoqParam: contextMoqParam,
		Strings:         strings,
	}
	mock.lockListAllDevices.Lock()
	mock.calls.ListAllDevices = append(mock.calls.ListAllDevices, callInfo)
	mock.lockListAllDevices.Unlock()
	return mock.ListAllDevicesFunc(contextMoqParam, strings)
}

// ListAllDevicesCalls gets all the calls that were made to ListAllDevices.
// Check the length with:
//     len(mockedDeviceManagement.ListAllDevicesCalls())
func (mock *DeviceManagementMock) ListAllDevicesCalls() []struct {
	ContextMoqParam context.Context
	Strings         []string
} {
	var calls []struct {
		ContextMoqParam context.Context
		Strings         []string
	}
	mock.lockListAllDevices.RLock()
	calls = mock.calls.ListAllDevices
	mock.lockListAllDevices.RUnlock()
	return calls
}

// ListEnvironments calls ListEnvironmentsFunc.
func (mock *DeviceManagementMock) ListEnvironments(contextMoqParam context.Context) ([]Environment, error) {
	if mock.ListEnvironmentsFunc == nil {
		panic("DeviceManagementMock.ListEnvironmentsFunc: method is nil but DeviceManagement.ListEnvironments was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockListEnvironments.Lock()
	mock.calls.ListEnvironments = append(mock.calls.ListEnvironments, callInfo)
	mock.lockListEnvironments.Unlock()
	return mock.ListEnvironmentsFunc(contextMoqParam)
}

// ListEnvironmentsCalls gets all the calls that were made to ListEnvironments.
// Check the length with:
//     len(mockedDeviceManagement.ListEnvironmentsCalls())
func (mock *DeviceManagementMock) ListEnvironmentsCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockListEnvironments.RLock()
	calls = mock.calls.ListEnvironments
	mock.lockListEnvironments.RUnlock()
	return calls
}

// NotifyStatus calls NotifyStatusFunc.
func (mock *DeviceManagementMock) NotifyStatus(ctx context.Context, deviceID string, message Status) error {
	if mock.NotifyStatusFunc == nil {
		panic("DeviceManagementMock.NotifyStatusFunc: method is nil but DeviceManagement.NotifyStatus was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		DeviceID string
		Message  Status
	}{
		Ctx:      ctx,
		DeviceID: deviceID,
		Message:  message,
	}
	mock.lockNotifyStatus.Lock()
	mock.calls.NotifyStatus = append(mock.calls.NotifyStatus, callInfo)
	mock.lockNotifyStatus.Unlock()
	return mock.NotifyStatusFunc(ctx, deviceID, message)
}

// NotifyStatusCalls gets all the calls that were made to NotifyStatus.
// Check the length with:
//     len(mockedDeviceManagement.NotifyStatusCalls())
func (mock *DeviceManagementMock) NotifyStatusCalls() []struct {
	Ctx      context.Context
	DeviceID string
	Message  Status
} {
	var calls []struct {
		Ctx      context.Context
		DeviceID string
		Message  Status
	}
	mock.lockNotifyStatus.RLock()
	calls = mock.calls.NotifyStatus
	mock.lockNotifyStatus.RUnlock()
	return calls
}

// SetStatusIfChanged calls SetStatusIfChangedFunc.
func (mock *DeviceManagementMock) SetStatusIfChanged(ctx context.Context, deviceID string, message Status) error {
	if mock.SetStatusIfChangedFunc == nil {
		panic("DeviceManagementMock.SetStatusIfChangedFunc: method is nil but DeviceManagement.SetStatusIfChanged was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		DeviceID string
		Message  Status
	}{
		Ctx:      ctx,
		DeviceID: deviceID,
		Message:  message,
	}
	mock.lockSetStatusIfChanged.Lock()
	mock.calls.SetStatusIfChanged = append(mock.calls.SetStatusIfChanged, callInfo)
	mock.lockSetStatusIfChanged.Unlock()
	return mock.SetStatusIfChangedFunc(ctx, deviceID, message)
}

// SetStatusIfChangedCalls gets all the calls that were made to SetStatusIfChanged.
// Check the length with:
//     len(mockedDeviceManagement.SetStatusIfChangedCalls())
func (mock *DeviceManagementMock) SetStatusIfChangedCalls() []struct {
	Ctx      context.Context
	DeviceID string
	Message  Status
} {
	var calls []struct {
		Ctx      context.Context
		DeviceID string
		Message  Status
	}
	mock.lockSetStatusIfChanged.RLock()
	calls = mock.calls.SetStatusIfChanged
	mock.lockSetStatusIfChanged.RUnlock()
	return calls
}

// UpdateDevice calls UpdateDeviceFunc.
func (mock *DeviceManagementMock) UpdateDevice(ctx context.Context, deviceID string, fields map[string]interface{}) (Device, error) {
	if mock.UpdateDeviceFunc == nil {
		panic("DeviceManagementMock.UpdateDeviceFunc: method is nil but DeviceManagement.UpdateDevice was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		DeviceID string
		Fields   map[string]interface{}
	}{
		Ctx:      ctx,
		DeviceID: deviceID,
		Fields:   fields,
	}
	mock.lockUpdateDevice.Lock()
	mock.calls.UpdateDevice = append(mock.calls.UpdateDevice, callInfo)
	mock.lockUpdateDevice.Unlock()
	return mock.UpdateDeviceFunc(ctx, deviceID, fields)
}

// UpdateDeviceCalls gets all the calls that were made to UpdateDevice.
// Check the length with:
//     len(mockedDeviceManagement.UpdateDeviceCalls())
func (mock *DeviceManagementMock) UpdateDeviceCalls() []struct {
	Ctx      context.Context
	DeviceID string
	Fields   map[string]interface{}
} {
	var calls []struct {
		Ctx      context.Context
		DeviceID string
		Fields   map[string]interface{}
	}
	mock.lockUpdateDevice.RLock()
	calls = mock.calls.UpdateDevice
	mock.lockUpdateDevice.RUnlock()
	return calls
}

// UpdateLastObservedOnDevice calls UpdateLastObservedOnDeviceFunc.
func (mock *DeviceManagementMock) UpdateLastObservedOnDevice(ctx context.Context, deviceID string, timestamp time.Time) error {
	if mock.UpdateLastObservedOnDeviceFunc == nil {
		panic("DeviceManagementMock.UpdateLastObservedOnDeviceFunc: method is nil but DeviceManagement.UpdateLastObservedOnDevice was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		DeviceID  string
		Timestamp time.Time
	}{
		Ctx:       ctx,
		DeviceID:  deviceID,
		Timestamp: timestamp,
	}
	mock.lockUpdateLastObservedOnDevice.Lock()
	mock.calls.UpdateLastObservedOnDevice = append(mock.calls.UpdateLastObservedOnDevice, callInfo)
	mock.lockUpdateLastObservedOnDevice.Unlock()
	return mock.UpdateLastObservedOnDeviceFunc(ctx, deviceID, timestamp)
}

// UpdateLastObservedOnDeviceCalls gets all the calls that were made to UpdateLastObservedOnDevice.
// Check the length with:
//     len(mockedDeviceManagement.UpdateLastObservedOnDeviceCalls())
func (mock *DeviceManagementMock) UpdateLastObservedOnDeviceCalls() []struct {
	Ctx       context.Context
	DeviceID  string
	Timestamp time.Time
} {
	var calls []struct {
		Ctx       context.Context
		DeviceID  string
		Timestamp time.Time
	}
	mock.lockUpdateLastObservedOnDevice.RLock()
	calls = mock.calls.UpdateLastObservedOnDevice
	mock.lockUpdateLastObservedOnDevice.RUnlock()
	return calls
}
