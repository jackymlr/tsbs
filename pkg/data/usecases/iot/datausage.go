package iot

import (
	"github.com/timescale/tsbs/pkg/data"
	"github.com/timescale/tsbs/pkg/data/usecases/common"
	"math/rand"
	"time"
)

var (
	labelDatausage        = []byte("datausage")
	labelTx        = []byte("tx")
	labelRx       = []byte("rx")
	trxUD            = common.UD(1597504, 9452686)

	datausageFields = []common.LabeledDistributionMaker{
		{
			Label: labelTx,
			DistributionMaker: func() common.Distribution {
				return common.FP(
					common.CWD(trxUD, 0, 10000, rand.Float64()*5000),
					0,
				)
			},
		},
		{
			Label: labelRx,
			DistributionMaker: func() common.Distribution {
				return common.FP(
					common.CWD(trxUD, 0, 10000, rand.Float64()*5000),
					0,
				)
			},
		},
	}
)

// DatausageMeasurement represents a subset of truck measurement datausage.
type DatausageMeasurement struct {
	*common.SubsystemMeasurement
}

// ToPoint serializes DatausageMeasurement to serialize.Point.
func (m *DatausageMeasurement) ToPoint(p *data.Point) {
	p.SetMeasurementName(labelDatausage)
	copy := m.Timestamp
	p.SetTimestamp(&copy)

	for i, d := range m.Distributions {
		p.AppendField(datausageFields[i].Label, float64(d.Get()))
	}
}

// NewDatausagesMeasurement creates a new DatausagesMeasurement with start time.
func NewDatausagesMeasurement(start time.Time) *DatausageMeasurement {
	sub := common.NewSubsystemMeasurementWithDistributionMakers(start, datausageFields)

	return &DatausageMeasurement{
		SubsystemMeasurement: sub,
	}
}
