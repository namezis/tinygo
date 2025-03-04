// +build nano_rp2040

// This contains the pin mappings for the Arduino Nano RP2040 Connect board.
//
// Sometimes the board is not detected even when the board is connected to your computer.
// To solve this, place a jumper wire between the REC and GND pins, then connect the board to your computer.
//
// For more information, see: https://store.arduino.cc/nano-rp2040-connect
// Also
// - Datasheets: https://docs.arduino.cc/hardware/nano-rp2040-connect
// - Nano RP2040 Connect technical reference: https://docs.arduino.cc/tutorials/nano-rp2040-connect/rp2040-01-technical-reference
//
package machine

// Digital Pins
const (
	D2  Pin = GPIO25
	D3  Pin = GPIO15
	D4  Pin = GPIO16
	D5  Pin = GPIO17
	D6  Pin = GPIO18
	D7  Pin = GPIO19
	D8  Pin = GPIO20
	D9  Pin = GPIO21
	D10 Pin = GPIO5
	D11 Pin = GPIO7
	D12 Pin = GPIO4
	D13 Pin = GPIO6
	D14 Pin = GPIO26
	D15 Pin = GPIO27
	D16 Pin = GPIO28
	D17 Pin = GPIO29
	D18 Pin = GPIO12
	D19 Pin = GPIO13
)

// Analog pins
const (
	A0 Pin = ADC0
	A1 Pin = ADC1
	A2 Pin = ADC2
	A3 Pin = ADC3
)

// Onboard LED
const (
	LED = GPIO6
)

// I2C pins
const (
	SDA_PIN Pin = GPIO12
	SCL_PIN Pin = GPIO13
)

// SPI pins
const (
	SPI0_SCK_PIN Pin = GPIO6
	SPI0_SDO_PIN Pin = GPIO7
	SPI0_SDI_PIN Pin = GPIO4
)

// NINA-W102 Pins
const (
	NINA_SCK Pin = GPIO14
	NINA_SDO Pin = GPIO11
	NINA_SDI Pin = GPIO8

	NINA_CS     Pin = GPIO9
	NINA_ACK    Pin = GPIO10
	NINA_GPIO0  Pin = GPIO0
	NINA_RESETN Pin = GPIO3

	NINA_TX Pin = GPIO9
	NINA_RX Pin = GPIO8
)

// Onboard crystal oscillator frequency, in MHz.
const (
	xoscFreq = 12 // MHz
)

// USB CDC identifiers
// https://github.com/arduino/ArduinoCore-mbed/blob/master/variants/NANO_RP2040_CONNECT/pins_arduino.h
const (
	usb_STRING_PRODUCT      = "Nano RP2040 Connect"
	usb_STRING_MANUFACTURER = "Arduino"
)

var (
	usb_VID uint16 = 0x2341
	usb_PID uint16 = 0x005e
)
