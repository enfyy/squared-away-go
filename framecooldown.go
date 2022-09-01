package main

type FrameCooldown struct {
	Counter int
	Max     int
	Active  bool
}

// Tick - Ticks the cooldown and returns true when cooldown is ongoing, else false
func (cooldown *FrameCooldown) Tick() bool {
	if cooldown.Active {
		cooldown.Counter++
		if cooldown.Counter < cooldown.Max {
			return true
		} else {
			cooldown.Counter = 0
			cooldown.Active = false
		}
	}
	return false
}
