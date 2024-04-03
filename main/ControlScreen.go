package main

import (
	lcd "github.com/wjessop/lcm1602_lcd"
	"golang.org/x/exp/io/i2c"
	"log"
)

func main() {
	// Configure this line with the device location and address of your device
	lcdDevice, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x27)
	if err != nil {
		log.Fatal(err)
	}
	defer lcdDevice.Close()

	lcdDisplay, err := lcd.NewLCM1602LCD(lcdDevice)
	if err != nil {
		log.Fatal(err)
	}

	// Write a string to row 1, position 0 (ie, the start of the line)
	if err := lcdDisplay.WriteString("Hello World!", 1, 0); err != nil {
		log.Fatal(err)
	}

	// Write a string to row 2, position 7
	if err := lcdDisplay.WriteString("(>'.'<)", 2, 7); err != nil {
		log.Fatal(err)
	}

	if err := lcdDisplay.Clear(); err != nil {
		log.Fatal(err)
	}
}
