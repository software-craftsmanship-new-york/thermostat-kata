package hvac

import "math/rand"

/////////////////////////////////////////////////////////////////////

func New() Hardware {
	return &HVAC{}
}

/////////////////////////////////////////////////////////////////////

// Hardware is a manufacture-specified interface for
// all hardware controlled by the thermostat.
type Hardware interface {
	AmbientTemperature() int

	SetHeater(on bool)
	SetCooler(on bool)
	SetBlower(on bool)
	SetColdAlarm(on bool)
	SetHeatAlarm(on bool)

	IsHeating() bool
	IsCooling() bool
	IsBlowing() bool
	ColdAlarm() bool
	HeatAlarm() bool
}

/////////////////////////////////////////////////////////////////////

// HVAC is a useless implementation of the Hardware interface,
// which will be replaced by the manufacture implementation later.
type HVAC struct {
	heating   bool
	cooling   bool
	blowing   bool
	coldAlarm bool
	heatAlarm bool
}

func (this *HVAC) AmbientTemperature() int { return rand.Intn(80) + 50 }

func (this *HVAC) SetHeater(on bool)    { this.heating = on }
func (this *HVAC) SetCooler(on bool)    { this.cooling = on }
func (this *HVAC) SetBlower(on bool)    { this.blowing = on }
func (this *HVAC) SetColdAlarm(on bool) { this.coldAlarm = on }
func (this *HVAC) SetHeatAlarm(on bool) { this.heatAlarm = on }

func (this *HVAC) IsHeating() bool { return this.heating }
func (this *HVAC) IsCooling() bool { return this.cooling }
func (this *HVAC) IsBlowing() bool { return this.blowing }
func (this *HVAC) ColdAlarm() bool { return this.coldAlarm }
func (this *HVAC) HeatAlarm() bool { return this.heatAlarm }

/////////////////////////////////////////////////////////////////////
