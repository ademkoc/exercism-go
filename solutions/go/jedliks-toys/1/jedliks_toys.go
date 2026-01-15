package jedlik

import "fmt"

func (c *Car) Drive() {
	if c.battery > c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}
func (c *Car) CanFinish(trackDistance int) bool {
	step := trackDistance / c.speed
	requiredBatteryDraim := step * c.batteryDrain
	// The battery shouldn't be completely depleted. At least 1% should remain.
	if c.battery > requiredBatteryDraim {
		return true
	}
	return false
}
