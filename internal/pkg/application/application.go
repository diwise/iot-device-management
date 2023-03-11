package application

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/diwise/iot-device-mgmt/internal/pkg/infrastructure/repositories/database"
	"github.com/diwise/iot-device-mgmt/pkg/types"
	"github.com/diwise/messaging-golang/pkg/messaging"
)

//go:generate moq -rm -out application_mock.go . App

type App interface {
	GetDevice(ctx context.Context, deviceID string) (types.Device, error)
	GetDeviceByEUI(ctx context.Context, eui string) (types.Device, error)
	GetDevices(ctx context.Context, tenants []string) ([]types.Device, error)
	CreateDevice(ctx context.Context, device types.Device) error
	UpdateDevice(ctx context.Context, deviceID string, fields map[string]interface{}) error

	HandleDeviceStatus(ctx context.Context, ds types.DeviceStatus) error
	SetStatus(ctx context.Context, deviceID string, message types.DeviceStatus) error

	GetTenants(ctx context.Context) ([]string, error)
	GetEnvironments(ctx context.Context) ([]types.Environment, error)
}

type app struct {
	store     database.Datastore
	messenger messaging.MsgContext
}

func New(s database.Datastore, m messaging.MsgContext) App {
	return &app{
		store:     s,
		messenger: m,
	}
}

func (a *app) HandleDeviceStatus(ctx context.Context, ds types.DeviceStatus) error {
	deviceID := ds.DeviceID
	timestamp, err := time.Parse(time.RFC3339Nano, ds.Timestamp)
	if err != nil {
		return fmt.Errorf("unable to parse timestamp from deviceStatus, %w", err)
	}

	err = a.store.UpdateLastObservedOnDevice(deviceID, timestamp)
	if err != nil {
		return fmt.Errorf("could not update last observed on device %s, %w", deviceID, err)
	}

	err = a.store.SetStatusIfChanged(MapStatus(ds))
	if err != nil {
		return fmt.Errorf("could not update status for device %s %w", deviceID, err)
	}

	d, err := a.store.GetDeviceFromID(deviceID)
	if err != nil {
		return err
	}

	return a.messenger.PublishOnTopic(ctx, &types.DeviceStatusUpdated{
		DeviceID:     deviceID,
		DeviceStatus: ds,
		Tenant:       d.Tenant.Name,
		Timestamp:    timestamp,
	})

}

func (a *app) GetDevice(ctx context.Context, deviceID string) (types.Device, error) {
	d, err := a.store.GetDeviceFromID(deviceID)
	if err != nil {
		return types.Device{}, err
	}

	status, err := a.store.GetLatestStatus(deviceID)
	if err != nil {
		return types.Device{}, err
	}

	return MapToModel(d, status), nil
}

func (a *app) GetDeviceByEUI(ctx context.Context, eui string) (types.Device, error) {
	d, err := a.store.GetDeviceFromDevEUI(eui)
	if err != nil {
		return types.Device{}, err
	}

	status, err := a.store.GetLatestStatus(d.DeviceId)
	if err != nil {
		return types.Device{}, err
	}

	return MapToModel(d, status), nil
}

func (a *app) GetDevices(ctx context.Context, tenants []string) ([]types.Device, error) {
	devices, err := a.store.GetAll(tenants...)
	if err != nil {
		return nil, err
	}

	models := make([]types.Device, 0)

	for _, d := range devices {
		status, err := a.store.GetLatestStatus(d.DeviceId)
		if err != nil {
			return nil, err
		}
		models = append(models, MapToModel(d, status))
	}

	return models, nil
}

func (a *app) CreateDevice(ctx context.Context, d types.Device) error {
	_, err := a.store.CreateDevice(d.DevEUI, d.DeviceID, d.Name, d.Description, d.Environment, d.SensorType.Name, d.Tenant, d.Location.Latitude, d.Location.Longitude, d.Types, d.Active)
	if err != nil {
		return err
	}

	return a.messenger.PublishOnTopic(ctx, &types.DeviceCreated{
		DeviceID:  d.DeviceID,
		Tenant:    d.Tenant,
		Timestamp: time.Now().UTC(),
	})
}

func (a *app) UpdateDevice(ctx context.Context, deviceID string, fields map[string]interface{}) error {
	d, err := a.store.UpdateDevice(deviceID, fields)
	if err != nil {
		return err
	}

	return a.messenger.PublishOnTopic(ctx, &types.DeviceUpdated{
		DeviceID:  deviceID,
		Tenant:    d.Tenant.Name,
		Timestamp: time.Now().UTC(),
	})
}

func (a *app) GetTenants(ctx context.Context) ([]string, error) {
	t := a.store.GetAllTenants()
	return t, nil
}

func (a *app) GetEnvironments(ctx context.Context) ([]types.Environment, error) {
	environments, err := a.store.ListEnvironments()
	if err != nil {
		return nil, err
	}

	return MapTo(environments, func(e database.Environment) types.Environment {
		return types.Environment{ID: e.ID, Name: e.Name}
	}), nil
}

func (a *app) SetStatus(ctx context.Context, deviceID string, message types.DeviceStatus) error {
	s := database.Status{
		DeviceID:     deviceID,
		BatteryLevel: message.BatteryLevel,
		Status:       message.Code,
		Messages:     strings.Join(message.Messages, ","),
		Timestamp:    message.Timestamp,
	}

	err := a.store.SetStatusIfChanged(s)
	if err != nil {
		return err
	}

	d, err := a.store.GetDeviceFromID(deviceID)
	if err != nil {
		return err
	}

	return a.messenger.PublishOnTopic(ctx, &types.DeviceStatusUpdated{
		DeviceID:     deviceID,
		DeviceStatus: message,
		Tenant:       d.Tenant.Name,
		Timestamp:    time.Now().UTC(),
	})
}
