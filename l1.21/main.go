package main

import "fmt"

type USB interface {
	ConnectWithUSB() string
}

type MicroUSBDevice struct{}

func (d *MicroUSBDevice) ConnectWithMicroUSB() string {
	return "Подключено через MicroUSB"
}

type MicroUSBToUSBAdapter struct {
	Device *MicroUSBDevice
}

func (a *MicroUSBToUSBAdapter) ConnectWithUSB() string {
	return a.Device.ConnectWithMicroUSB()
}

func ChargeWithUSB(usbDevice USB) {
	fmt.Println(usbDevice.ConnectWithUSB())
}

func main() {
	microUSBDevice := &MicroUSBDevice{}

	adapter := &MicroUSBToUSBAdapter{
		Device: microUSBDevice,
	}

	ChargeWithUSB(adapter)
}
