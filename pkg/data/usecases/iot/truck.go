package iot

import (
	"fmt"
	"github.com/timescale/tsbs/pkg/data/usecases/common"
	"time"
)

const (
	truckNameFmt = "5f59ce7a77d0e84d4b1df4a%d"
)

type model struct {
	Name            string
	LoadCapacity    float32
	FuelCapacity    float32
	FuelConsumption float32
}

var (
	driverChoices = []string{
		"Derek",
		"Rodney",
		"Albert",
		"Andy",
		"Seth",
		"Trish",
	}

	typeChoices = []string{
		"sim1",
		"sim2",
		"wan1",
		"wan2",
		"wireless",
	}

	oidChoices = []string{
		"5f1ec84fb48129250cdb12a2",
		"5f1f93e5b48129250cdb12b5",
		"5f1fd5f1bae5da7a7331f642",
		"5fd88a91d9c52860963643b0",
		"5febea078019ad2705652598",
	}

	modelChoices = []model{
		{
			Name:            "F-150",
			LoadCapacity:    2000,
			FuelCapacity:    200,
			FuelConsumption: 15,
		},
		{
			Name:            "G-2000",
			LoadCapacity:    5000,
			FuelCapacity:    300,
			FuelConsumption: 19,
		},
		{
			Name:            "H-2",
			LoadCapacity:    1500,
			FuelCapacity:    150,
			FuelConsumption: 12,
		},
	}

	deviceVersionChoices = []string{
		"v1.0",
		"v1.5",
		"v2.0",
		"v2.3",
	}

	// FleetChoices contains all the fleet name values for the IoT use case
	FleetChoices = []string{
		"East",
		"West",
		"North",
		"South",
	}

	TypeChoices = []string{
		"wan1",
		"wan2",
		"sim1",
		"sim2",
		"wireless",
	}
)

// Truck models a truck outfitted with an IoT device which sends back measurements.
type Truck struct {
	simulatedMeasurements []common.SimulatedMeasurement
	tags                  []common.Tag
}

// TickAll advances all Distributions of a Truck.
func (t *Truck) TickAll(d time.Duration) {
	for i := range t.simulatedMeasurements {
		t.simulatedMeasurements[i].Tick(d)
	}
}

// Measurements returns the trucks measurements.
func (t Truck) Measurements() []common.SimulatedMeasurement {
	return t.simulatedMeasurements
}

// Tags returns the truck tags.
func (t Truck) Tags() []common.Tag {
	return t.tags
}

func newTruckMeasurements(start time.Time) []common.SimulatedMeasurement {
	return []common.SimulatedMeasurement{
		// NewReadingsMeasurement(start),
		// NewDiagnosticsMeasurement(start),
		NewDatausagesMeasurement(start),
	}
}

// NewTruck creates a new truck in a simulated iot use case
func NewTruck(i int, start time.Time) common.Generator {
	truck := newTruckWithMeasurementGenerator(i, start, newTruckMeasurements)
	return &truck
}

func newTruckWithMeasurementGenerator(i int, start time.Time, generator func(time.Time) []common.SimulatedMeasurement) Truck {
	sm := generator(start)

	//m := modelChoices[rand.Intn(len(modelChoices))]

	h := Truck{
		tags: []common.Tag{
		{Key: []byte("deviceId"), Value: fmt.Sprintf(truckNameFmt, i)},
		{Key: []byte("oid"), Value: common.RandomStringSliceChoice(oidChoices)},
		{Key: []byte("type"), Value: common.RandomStringSliceChoice(typeChoices)},
		//	{Key: []byte("name"), Value: fmt.Sprintf(truckNameFmt, i)},
		//	{Key: []byte("fleet"), Value: common.RandomStringSliceChoice(FleetChoices)},
		//	{Key: []byte("driver"), Value: common.RandomStringSliceChoice(driverChoices)},
		//	{Key: []byte("model"), Value: m.Name},
		//	{Key: []byte("device_version"), Value: common.RandomStringSliceChoice(deviceVersionChoices)},
		//	{Key: []byte("load_capacity"), Value: m.LoadCapacity},
		//	{Key: []byte("fuel_capacity"), Value: m.FuelCapacity},
		//	{Key: []byte("nominal_fuel_consumption"), Value: m.FuelConsumption},
		},
		simulatedMeasurements: sm,
	}

	return h
}
