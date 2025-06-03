package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	temperatureMeasures := make([]SensorReading, 10)
	humidityMeasures := make([]SensorReading, 10)
	windMeasures := make([]SensorReading, 10)

	var wg sync.WaitGroup
	temperature := make(chan SensorReading, 1)
	humidity := make(chan SensorReading, 1)
	wind := make(chan SensorReading, 1)

	for i := range 10 {
		
		wg.Add(3)
		go func(i int) {
			defer wg.Done()
			measur := *TemperatureMeasure()
			temperature <- measur

		}(i)

		go func(i int) {
			defer wg.Done()
			measur := *HumidityMesure()
			humidity <- measur

		}(i)

		go func(i int) {
			defer wg.Done()
			measur := *WindMesure()
			wind <- measur

		}(i)

		wg.Wait()
		temperatureMeasures[i] = <-temperature
		humidityMeasures[i] = <-humidity
		windMeasures[i] = <-wind
		fmt.Printf("Temperature: %.3vC\nHumidity: %v%%\nWind: %v KM\n", temperatureMeasures[i].val, humidityMeasures[i].val, windMeasures[i].val)

		time.Sleep(time.Second )
	}

}