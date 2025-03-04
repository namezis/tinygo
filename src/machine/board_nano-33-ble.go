// +build nano_33_ble

// This contains the pin mappings for the Arduino Nano 33 BLE [Sense] boards.
//
// Flashing the board requires special version of bossac.
//
// This executable can be obtained two ways:
// 1) In Arduino IDE, install support for the board ("Arduino Mbed OS Nano Boards")
//    Search for "tools/bossac/1.9.1-arduino2/bossac" in Arduino IDEs directory
// 2) Download https://downloads.arduino.cc/packages/package_index.json
//    Search for "bossac-1.9.1-arduino2" in that file
//    Download tarball for your OS and unpack it
//
// Once you have the executable, make it accessible in your PATH as "bossac_arduino2".
//
// It is possible to replace original bossac with this new one (this only adds support for nrf chip).
// In that case make "bossac_arduino2" symlink on it, for the board target to be able to find it.
//
// For more information, see:
// - https://store.arduino.cc/arduino-nano-33-ble
// - https://store.arduino.cc/arduino-nano-33-ble-sense
//
package machine

const HasLowFrequencyCrystal = true

// Digital Pins
const (
	D2  Pin = P1_11
	D3  Pin = P1_12
	D4  Pin = P1_15
	D5  Pin = P1_13
	D6  Pin = P1_14
	D7  Pin = P0_23
	D8  Pin = P0_21
	D9  Pin = P0_27
	D10 Pin = P1_02
	D11 Pin = P1_01
	D12 Pin = P1_08
	D13 Pin = P0_13
)

// Analog pins
const (
	A0 Pin = P0_04
	A1 Pin = P0_05
	A2 Pin = P0_30
	A3 Pin = P0_29
	A4 Pin = P0_31
	A5 Pin = P0_02
	A6 Pin = P0_28
	A7 Pin = P0_03
)

// Onboard LEDs
const (
	LED         = LED_BUILTIN
	LED1        = LED_RED
	LED2        = LED_GREEN
	LED3        = LED_BLUE
	LED_BUILTIN = P0_13
	LED_RED     = P0_24
	LED_GREEN   = P0_16
	LED_BLUE    = P0_06
)

// UART0 pins
const (
	UART_RX_PIN = P1_10
	UART_TX_PIN = P1_03
)

// I2C pins
const (
	SDA_PIN = P0_31
	SCL_PIN = P0_02
)

// SPI pins
const (
	SPI0_SCK_PIN = P0_13
	SPI0_SDO_PIN = P1_01
	SPI0_SDI_PIN = P1_08
)

// USB CDC identifiers
const (
	usb_STRING_PRODUCT      = "Nano 33 BLE"
	usb_STRING_MANUFACTURER = "Arduino"
)

var (
	usb_VID uint16 = 0x2341
	usb_PID uint16 = 0x805a
)
