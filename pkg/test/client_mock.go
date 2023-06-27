// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package test

import (
	"context"
	"sync"

	"github.com/diwise/iot-device-mgmt/pkg/client"
)

// Ensure, that DeviceManagementClientMock does implement DeviceManagementClient.
// If this is not the case, regenerate this file with moq.
var _ client.DeviceManagementClient = &DeviceManagementClientMock{}

// DeviceManagementClientMock is a mock implementation of DeviceManagementClient.
//
//	func TestSomethingThatUsesDeviceManagementClient(t *testing.T) {
//
//		// make and configure a mocked DeviceManagementClient
//		mockedDeviceManagementClient := &DeviceManagementClientMock{
//			CloseFunc: func(ctx context.Context)  {
//				panic("mock out the Close method")
//			},
//			FindDeviceFromDevEUIFunc: func(ctx context.Context, devEUI string) (client.Device, error) {
//				panic("mock out the FindDeviceFromDevEUI method")
//			},
//			FindDeviceFromInternalIDFunc: func(ctx context.Context, deviceID string) (client.Device, error) {
//				panic("mock out the FindDeviceFromInternalID method")
//			},
//		}
//
//		// use mockedDeviceManagementClient in code that requires DeviceManagementClient
//		// and then make assertions.
//
//	}
type DeviceManagementClientMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func(ctx context.Context)

	// FindDeviceFromDevEUIFunc mocks the FindDeviceFromDevEUI method.
	FindDeviceFromDevEUIFunc func(ctx context.Context, devEUI string) (client.Device, error)

	// FindDeviceFromInternalIDFunc mocks the FindDeviceFromInternalID method.
	FindDeviceFromInternalIDFunc func(ctx context.Context, deviceID string) (client.Device, error)

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// FindDeviceFromDevEUI holds details about calls to the FindDeviceFromDevEUI method.
		FindDeviceFromDevEUI []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DevEUI is the devEUI argument value.
			DevEUI string
		}
		// FindDeviceFromInternalID holds details about calls to the FindDeviceFromInternalID method.
		FindDeviceFromInternalID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeviceID is the deviceID argument value.
			DeviceID string
		}
	}
	lockClose                    sync.RWMutex
	lockFindDeviceFromDevEUI     sync.RWMutex
	lockFindDeviceFromInternalID sync.RWMutex
}

// Close calls CloseFunc.
func (mock *DeviceManagementClientMock) Close(ctx context.Context) {
	if mock.CloseFunc == nil {
		panic("DeviceManagementClientMock.CloseFunc: method is nil but DeviceManagementClient.Close was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	mock.CloseFunc(ctx)
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//
//	len(mockedDeviceManagementClient.CloseCalls())
func (mock *DeviceManagementClientMock) CloseCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// FindDeviceFromDevEUI calls FindDeviceFromDevEUIFunc.
func (mock *DeviceManagementClientMock) FindDeviceFromDevEUI(ctx context.Context, devEUI string) (client.Device, error) {
	if mock.FindDeviceFromDevEUIFunc == nil {
		panic("DeviceManagementClientMock.FindDeviceFromDevEUIFunc: method is nil but DeviceManagementClient.FindDeviceFromDevEUI was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		DevEUI string
	}{
		Ctx:    ctx,
		DevEUI: devEUI,
	}
	mock.lockFindDeviceFromDevEUI.Lock()
	mock.calls.FindDeviceFromDevEUI = append(mock.calls.FindDeviceFromDevEUI, callInfo)
	mock.lockFindDeviceFromDevEUI.Unlock()
	return mock.FindDeviceFromDevEUIFunc(ctx, devEUI)
}

// FindDeviceFromDevEUICalls gets all the calls that were made to FindDeviceFromDevEUI.
// Check the length with:
//
//	len(mockedDeviceManagementClient.FindDeviceFromDevEUICalls())
func (mock *DeviceManagementClientMock) FindDeviceFromDevEUICalls() []struct {
	Ctx    context.Context
	DevEUI string
} {
	var calls []struct {
		Ctx    context.Context
		DevEUI string
	}
	mock.lockFindDeviceFromDevEUI.RLock()
	calls = mock.calls.FindDeviceFromDevEUI
	mock.lockFindDeviceFromDevEUI.RUnlock()
	return calls
}

// FindDeviceFromInternalID calls FindDeviceFromInternalIDFunc.
func (mock *DeviceManagementClientMock) FindDeviceFromInternalID(ctx context.Context, deviceID string) (client.Device, error) {
	if mock.FindDeviceFromInternalIDFunc == nil {
		panic("DeviceManagementClientMock.FindDeviceFromInternalIDFunc: method is nil but DeviceManagementClient.FindDeviceFromInternalID was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		DeviceID string
	}{
		Ctx:      ctx,
		DeviceID: deviceID,
	}
	mock.lockFindDeviceFromInternalID.Lock()
	mock.calls.FindDeviceFromInternalID = append(mock.calls.FindDeviceFromInternalID, callInfo)
	mock.lockFindDeviceFromInternalID.Unlock()
	return mock.FindDeviceFromInternalIDFunc(ctx, deviceID)
}

// FindDeviceFromInternalIDCalls gets all the calls that were made to FindDeviceFromInternalID.
// Check the length with:
//
//	len(mockedDeviceManagementClient.FindDeviceFromInternalIDCalls())
func (mock *DeviceManagementClientMock) FindDeviceFromInternalIDCalls() []struct {
	Ctx      context.Context
	DeviceID string
} {
	var calls []struct {
		Ctx      context.Context
		DeviceID string
	}
	mock.lockFindDeviceFromInternalID.RLock()
	calls = mock.calls.FindDeviceFromInternalID
	mock.lockFindDeviceFromInternalID.RUnlock()
	return calls
}
