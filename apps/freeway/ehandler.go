package freeway

import (
	"time"

	"github.com/jrecuero/go-cli/tools"
)

// EHandler represents ...
type EHandler struct {
	freeway  *Freeway
	devices  map[string]IDevice
	running  bool
	delay    int
	workerCb func(IDevice)
}

// SetDelay is ...
func (ehdlr *EHandler) SetDelay(delay int) *EHandler {
	ehdlr.delay = delay
	return ehdlr
}

// SetWorkerCb is ...
func (ehdlr *EHandler) SetWorkerCb(cb func(IDevice)) *EHandler {
	ehdlr.workerCb = cb
	return ehdlr
}

// SetFreeway is ...
func (ehdlr *EHandler) SetFreeway(freeway *Freeway) *EHandler {
	ehdlr.freeway = freeway
	return ehdlr
}

// GetFreeway is ...
func (ehdlr *EHandler) GetFreeway() *Freeway {
	return ehdlr.freeway
}

// AddDevice is ...
func (ehdlr *EHandler) AddDevice(dev IDevice) *EHandler {
	ehdlr.devices[dev.GetName()] = dev
	return ehdlr
}

// GetDevices is ...
func (ehdlr *EHandler) GetDevices() map[string]IDevice {
	return ehdlr.devices
}

// GetDeviceByName is ...
func (ehdlr *EHandler) GetDeviceByName(name string) IDevice {
	return ehdlr.devices[name]
}

// worker is ...
func (ehdlr *EHandler) worker(device IDevice) {
	for ehdlr.running {
		device.FreewayTraverse()
		//tools.ToDisplay("worker device: %s\n", device)
		if ehdlr.workerCb != nil {
			ehdlr.workerCb(device)
		}
		time.Sleep(time.Duration(ehdlr.delay) * time.Millisecond)
	}
}

// Setup is ...
func (ehdlr *EHandler) Setup() {
	for _, device := range ehdlr.devices {
		device.NewLocation(ehdlr.GetFreeway())
		tools.ToDisplay("setup device %s\n", device)
	}
}

// Start is ...
func (ehdlr *EHandler) Start() {
	ehdlr.running = true
	for _, device := range ehdlr.devices {
		go ehdlr.worker(device)
	}
	tools.ToDisplay("EHandler started ...\n")
	for ehdlr.running {
		time.Sleep(time.Duration(ehdlr.delay) * time.Millisecond)
	}
	tools.ToDisplay("EHandler stopped\n")
}

// Stop is ...
func (ehdlr *EHandler) Stop() {
	ehdlr.running = false
}

// NewEHandler is ...
func NewEHandler() *EHandler {
	return &EHandler{
		devices: make(map[string]IDevice),
		running: false,
		delay:   10,
	}
}
