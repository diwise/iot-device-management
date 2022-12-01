// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package test

import (
	"sync"
	
	"github.com/diwise/iot-device-mgmt/pkg/client"
)

// Ensure, that DeviceMock does implement Device.
// If this is not the case, regenerate this file with moq.
var _ client.Device = &DeviceMock{}

// DeviceMock is a mock implementation of Device.
//
//	func TestSomethingThatUsesDevice(t *testing.T) {
//
//		// make and configure a mocked Device
//		mockedDevice := &DeviceMock{
//			EnvironmentFunc: func() string {
//				panic("mock out the Environment method")
//			},
//			IDFunc: func() string {
//				panic("mock out the ID method")
//			},
//			IsActiveFunc: func() bool {
//				panic("mock out the IsActive method")
//			},
//			LatitudeFunc: func() float64 {
//				panic("mock out the Latitude method")
//			},
//			LongitudeFunc: func() float64 {
//				panic("mock out the Longitude method")
//			},
//			SensorTypeFunc: func() string {
//				panic("mock out the SensorType method")
//			},
//			TenantFunc: func() string {
//				panic("mock out the Tenant method")
//			},
//			TypesFunc: func() []string {
//				panic("mock out the Types method")
//			},
//		}
//
//		// use mockedDevice in code that requires Device
//		// and then make assertions.
//
//	}
type DeviceMock struct {
	// EnvironmentFunc mocks the Environment method.
	EnvironmentFunc func() string

	// IDFunc mocks the ID method.
	IDFunc func() string

	// IsActiveFunc mocks the IsActive method.
	IsActiveFunc func() bool

	// LatitudeFunc mocks the Latitude method.
	LatitudeFunc func() float64

	// LongitudeFunc mocks the Longitude method.
	LongitudeFunc func() float64

	// SensorTypeFunc mocks the SensorType method.
	SensorTypeFunc func() string

	// TenantFunc mocks the Tenant method.
	TenantFunc func() string

	// TypesFunc mocks the Types method.
	TypesFunc func() []string

	// calls tracks calls to the methods.
	calls struct {
		// Environment holds details about calls to the Environment method.
		Environment []struct {
		}
		// ID holds details about calls to the ID method.
		ID []struct {
		}
		// IsActive holds details about calls to the IsActive method.
		IsActive []struct {
		}
		// Latitude holds details about calls to the Latitude method.
		Latitude []struct {
		}
		// Longitude holds details about calls to the Longitude method.
		Longitude []struct {
		}
		// SensorType holds details about calls to the SensorType method.
		SensorType []struct {
		}
		// Tenant holds details about calls to the Tenant method.
		Tenant []struct {
		}
		// Types holds details about calls to the Types method.
		Types []struct {
		}
	}
	lockEnvironment sync.RWMutex
	lockID          sync.RWMutex
	lockIsActive    sync.RWMutex
	lockLatitude    sync.RWMutex
	lockLongitude   sync.RWMutex
	lockSensorType  sync.RWMutex
	lockTenant      sync.RWMutex
	lockTypes       sync.RWMutex
}

// Environment calls EnvironmentFunc.
func (mock *DeviceMock) Environment() string {
	if mock.EnvironmentFunc == nil {
		panic("DeviceMock.EnvironmentFunc: method is nil but Device.Environment was just called")
	}
	callInfo := struct {
	}{}
	mock.lockEnvironment.Lock()
	mock.calls.Environment = append(mock.calls.Environment, callInfo)
	mock.lockEnvironment.Unlock()
	return mock.EnvironmentFunc()
}

// EnvironmentCalls gets all the calls that were made to Environment.
// Check the length with:
//
//	len(mockedDevice.EnvironmentCalls())
func (mock *DeviceMock) EnvironmentCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockEnvironment.RLock()
	calls = mock.calls.Environment
	mock.lockEnvironment.RUnlock()
	return calls
}

// ID calls IDFunc.
func (mock *DeviceMock) ID() string {
	if mock.IDFunc == nil {
		panic("DeviceMock.IDFunc: method is nil but Device.ID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockID.Lock()
	mock.calls.ID = append(mock.calls.ID, callInfo)
	mock.lockID.Unlock()
	return mock.IDFunc()
}

// IDCalls gets all the calls that were made to ID.
// Check the length with:
//
//	len(mockedDevice.IDCalls())
func (mock *DeviceMock) IDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockID.RLock()
	calls = mock.calls.ID
	mock.lockID.RUnlock()
	return calls
}

// IsActive calls IsActiveFunc.
func (mock *DeviceMock) IsActive() bool {
	if mock.IsActiveFunc == nil {
		panic("DeviceMock.IsActiveFunc: method is nil but Device.IsActive was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsActive.Lock()
	mock.calls.IsActive = append(mock.calls.IsActive, callInfo)
	mock.lockIsActive.Unlock()
	return mock.IsActiveFunc()
}

// IsActiveCalls gets all the calls that were made to IsActive.
// Check the length with:
//
//	len(mockedDevice.IsActiveCalls())
func (mock *DeviceMock) IsActiveCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsActive.RLock()
	calls = mock.calls.IsActive
	mock.lockIsActive.RUnlock()
	return calls
}

// Latitude calls LatitudeFunc.
func (mock *DeviceMock) Latitude() float64 {
	if mock.LatitudeFunc == nil {
		panic("DeviceMock.LatitudeFunc: method is nil but Device.Latitude was just called")
	}
	callInfo := struct {
	}{}
	mock.lockLatitude.Lock()
	mock.calls.Latitude = append(mock.calls.Latitude, callInfo)
	mock.lockLatitude.Unlock()
	return mock.LatitudeFunc()
}

// LatitudeCalls gets all the calls that were made to Latitude.
// Check the length with:
//
//	len(mockedDevice.LatitudeCalls())
func (mock *DeviceMock) LatitudeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockLatitude.RLock()
	calls = mock.calls.Latitude
	mock.lockLatitude.RUnlock()
	return calls
}

// Longitude calls LongitudeFunc.
func (mock *DeviceMock) Longitude() float64 {
	if mock.LongitudeFunc == nil {
		panic("DeviceMock.LongitudeFunc: method is nil but Device.Longitude was just called")
	}
	callInfo := struct {
	}{}
	mock.lockLongitude.Lock()
	mock.calls.Longitude = append(mock.calls.Longitude, callInfo)
	mock.lockLongitude.Unlock()
	return mock.LongitudeFunc()
}

// LongitudeCalls gets all the calls that were made to Longitude.
// Check the length with:
//
//	len(mockedDevice.LongitudeCalls())
func (mock *DeviceMock) LongitudeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockLongitude.RLock()
	calls = mock.calls.Longitude
	mock.lockLongitude.RUnlock()
	return calls
}

// SensorType calls SensorTypeFunc.
func (mock *DeviceMock) SensorType() string {
	if mock.SensorTypeFunc == nil {
		panic("DeviceMock.SensorTypeFunc: method is nil but Device.SensorType was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSensorType.Lock()
	mock.calls.SensorType = append(mock.calls.SensorType, callInfo)
	mock.lockSensorType.Unlock()
	return mock.SensorTypeFunc()
}

// SensorTypeCalls gets all the calls that were made to SensorType.
// Check the length with:
//
//	len(mockedDevice.SensorTypeCalls())
func (mock *DeviceMock) SensorTypeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSensorType.RLock()
	calls = mock.calls.SensorType
	mock.lockSensorType.RUnlock()
	return calls
}

// Tenant calls TenantFunc.
func (mock *DeviceMock) Tenant() string {
	if mock.TenantFunc == nil {
		panic("DeviceMock.TenantFunc: method is nil but Device.Tenant was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTenant.Lock()
	mock.calls.Tenant = append(mock.calls.Tenant, callInfo)
	mock.lockTenant.Unlock()
	return mock.TenantFunc()
}

// TenantCalls gets all the calls that were made to Tenant.
// Check the length with:
//
//	len(mockedDevice.TenantCalls())
func (mock *DeviceMock) TenantCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTenant.RLock()
	calls = mock.calls.Tenant
	mock.lockTenant.RUnlock()
	return calls
}

// Types calls TypesFunc.
func (mock *DeviceMock) Types() []string {
	if mock.TypesFunc == nil {
		panic("DeviceMock.TypesFunc: method is nil but Device.Types was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTypes.Lock()
	mock.calls.Types = append(mock.calls.Types, callInfo)
	mock.lockTypes.Unlock()
	return mock.TypesFunc()
}

// TypesCalls gets all the calls that were made to Types.
// Check the length with:
//
//	len(mockedDevice.TypesCalls())
func (mock *DeviceMock) TypesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTypes.RLock()
	calls = mock.calls.Types
	mock.lockTypes.RUnlock()
	return calls
}
