package main

var TemperatureStation = SensorReading{sensor: TemperatureSensor}

var TemperatureChannel = make(chan SensorReading, 1)

func TemperatureMeasure() *SensorReading {
	return (TemperatureStation.Measure())
}

var HumidityStation = SensorReading{sensor: HumiditySensor}

var HumidityChannel = make(chan SensorReading, 1)

func HumidityMesure() *SensorReading {
	return (HumidityStation.Measure())
}

var WindStation = SensorReading{sensor: WindSensor}

var WindChannel = make(chan SensorReading, 1)

func WindMesure() *SensorReading {
	return (WindStation.Measure())
}
