package main

import (
	"log"
	"time"

	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/devices/adafruit/charlcd"
)

func main() {
	// Initialize the I2C bus.
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	// Initialize the LCD.
	lcd, err := charlcd.NewI2C(bus, 0x27, 16, 2)
	if err != nil {
		log.Fatal(err)
	}
	defer lcd.Halt()

	// Display some text.
	lcd.ClearDisplay()
	lcd.SetCursor(0, 0)
	lcd.Print("Hello,")
	lcd.SetCursor(0, 1)
	lcd.Print("Raspberry Pi!")

	// Wait for a while.
	time.Sleep(5 * time.Second)
}
