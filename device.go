package windevice

import (
	"syscall"

	"github.com/gentlemanautomaton/windevice/deviceproperty"
	"github.com/gentlemanautomaton/windevice/diflagex"
	"github.com/gentlemanautomaton/windevice/drivertype"
	"github.com/gentlemanautomaton/windevice/installstate"
	"github.com/gentlemanautomaton/windevice/setupapi"
)

// Device provides access to Windows device information while executing a query.
// It can be copied by value.
//
// Device stores a system handle internally and shouldn't be used outside of a
// query callback.
type Device struct {
	devices syscall.Handle
	data    setupapi.DevInfoData
}

// Sys returns low-level information about the device.
func (device Device) Sys() (devices syscall.Handle, data setupapi.DevInfoData) {
	return device.devices, device.data
}

// Drivers returns a driver set that contains drivers affiliated with the
// device.
func (device Device) Drivers(q DriverQuery) DriverSet {
	return DriverSet{
		devices: device.devices,
		device:  device.data, // TODO: Clone the data first?
		query:   q,
	}
}

// InstalledDriver returns a driver set that contains the device's currently
// installed driver.
func (device Device) InstalledDriver() DriverSet {
	return DriverSet{
		devices: device.devices,
		device:  device.data, // TODO: Clone the data first?
		query: DriverQuery{
			Type:    drivertype.ClassDriver,
			FlagsEx: diflagex.InstalledDriver | diflagex.AllowExcludedDrivers,
		},
	}
}

// DeviceInstanceID returns the device instance ID of the device.
func (device Device) DeviceInstanceID() (string, error) {
	return setupapi.GetDeviceInstanceID(device.devices, device.data)
}

// Description returns the description of the device.
func (device Device) Description() (string, error) {
	return setupapi.GetDeviceRegistryString(device.devices, device.data, deviceproperty.Description)
}

// HardwareID returns the set of hardware IDs associated with the device.
func (device Device) HardwareID() ([]string, error) {
	return setupapi.GetDeviceRegistryStrings(device.devices, device.data, deviceproperty.HardwareID)
}

// CompatibleID returns the set of compatible IDs associated with the device.
func (device Device) CompatibleID() ([]string, error) {
	return setupapi.GetDeviceRegistryStrings(device.devices, device.data, deviceproperty.CompatibleID)
}

// Service returns the service for the device.
func (device Device) Service() (string, error) {
	return setupapi.GetDeviceRegistryString(device.devices, device.data, deviceproperty.Service)
}

// Class returns the class name of the device.
func (device Device) Class() (string, error) {
	return setupapi.GetDeviceRegistryString(device.devices, device.data, deviceproperty.Class)
}

// ClassGUID returns a string representation of the globally unique identifier
// of the device's class.
func (device Device) ClassGUID() (string, error) {
	return setupapi.GetDeviceRegistryString(device.devices, device.data, deviceproperty.ClassGUID)
}

// ConfigFlags returns the configuration flags for the device.
func (device Device) ConfigFlags() (uint32, error) {
	return setupapi.GetDeviceRegistryUint32(device.devices, device.data, deviceproperty.ConfigFlags)
}

// DriverRegName returns the registry name of the device's driver.
func (device Device) DriverRegName() (string, error) {
	return setupapi.GetDeviceRegistryString(device.devices, device.data, deviceproperty.DriverRegName)
}

// Manufacturer returns the manufacturer of the device.
func (device Device) Manufacturer() (string, error) {
	return setupapi.GetDeviceRegistryString(device.devices, device.data, deviceproperty.Manufacturer)
}

// FriendlyName returns the friendly name of the device.
func (device Device) FriendlyName() (string, error) {
	return setupapi.GetDeviceRegistryString(device.devices, device.data, deviceproperty.FriendlyName)
}

// LocationInformation returns the location information for the device.
func (device Device) LocationInformation() (string, error) {
	return setupapi.GetDeviceRegistryString(device.devices, device.data, deviceproperty.LocationInformation)
}

// PhysicalDeviceObjectName returns the physical object name of the device.
func (device Device) PhysicalDeviceObjectName() (string, error) {
	return setupapi.GetDeviceRegistryString(device.devices, device.data, deviceproperty.PhysicalDeviceObjectName)
}

// EnumeratorName returns the name of the device's enumerator.
func (device Device) EnumeratorName() (string, error) {
	return setupapi.GetDeviceRegistryString(device.devices, device.data, deviceproperty.EnumeratorName)
}

// DevType returns the type of the device.
func (device Device) DevType() (uint32, error) {
	return setupapi.GetDeviceRegistryUint32(device.devices, device.data, deviceproperty.DevType)
}

// Characteristics returns the characteristics of the device.
func (device Device) Characteristics() (uint32, error) {
	return setupapi.GetDeviceRegistryUint32(device.devices, device.data, deviceproperty.Characteristics)
}

// InstallState returns the installation state of the device.
func (device Device) InstallState() (installstate.State, error) {
	state, err := setupapi.GetDeviceRegistryUint32(device.devices, device.data, deviceproperty.InstallState)
	return installstate.State(state), err
}
