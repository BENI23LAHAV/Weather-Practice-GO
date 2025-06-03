package main

import "math/rand"

type SensorType int

const (
	TemperatureSensor SensorType = iota
	HumiditySensor
	WindSensor
)

var SensorsType = map[SensorType]string{
	TemperatureSensor : "Temperature",
	HumiditySensor : "Humidity",
	WindSensor: "Wind",
}


type SensorReading struct {
	sensor SensorType
	val    float32
}

func (sr *SensorReading) Measure() *SensorReading{
    switch sr.sensor {
    case TemperatureSensor:
        sr.val = -10.0 + rand.Float32()*50 
    case HumiditySensor:
        sr.val = float32(rand.Intn(101)) 
    case WindSensor:
        sr.val = float32(rand.Intn(151))
    }
return sr
}


type Sensor interface {
	Measure() *SensorReading
}
