package driver

import (
	"fmt"
	"sync"

	"device-watchdog-go/internal/pkg/logger"

	"github.com/edgexfoundry/device-sdk-go/v4/pkg/interfaces"
	sdkModel "github.com/edgexfoundry/device-sdk-go/v4/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

var once sync.Once
var driver *Driver

type Driver struct {
	Logger  logger.LoggingClient         // 自定义logger客户端
	AsyncCh chan<- *sdkModel.AsyncValues // 异步上报通道
}

func (d *Driver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties,
	reqs []sdkModel.CommandRequest) (responses []*sdkModel.CommandValue, err error) {
	return nil, nil
}

func (d *Driver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties,
	reqs []sdkModel.CommandRequest, params []*sdkModel.CommandValue) error {

	return nil
}

func (d *Driver) Initialize(sdk interfaces.DeviceServiceSDK) error {
	d.Logger = logger.NewClient("DEBUG")
	d.AsyncCh = sdk.AsyncValuesChannel()
	return nil
}

func (d *Driver) Start() error {
	d.Logger.Debugf("Device %s is started")
	return nil
}

func (d *Driver) Stop(force bool) error {
	d.Logger.Debugf("Device is stopped")
	return nil
}

func (d *Driver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.Logger.Debugf("Device %s is added", deviceName)
	return nil
}

func (d *Driver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.Logger.Debugf("Device %s is updated", deviceName)
	return nil
}

func (d *Driver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	d.Logger.Debugf("Device %s is removed", deviceName)
	return nil
}

func (d *Driver) Discover() error {
	return fmt.Errorf("driver's Discover function isn't implemented")
}

func (d *Driver) ValidateDevice(device models.Device) error {

	fmt.Printf("triggers device's protocol properties validation: %v", device)

	return nil
}

func NewProtocolDriver() interfaces.ProtocolDriver {
	once.Do(func() { //	确保某段代码只执行一次
		driver = new(Driver)
	})
	return driver
}
