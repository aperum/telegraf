package fritzstats

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

const (
	measurement = "fritzstats"
)

type Fritz struct {
	Url string
}

var sampleConfig = `
  ## address of the fritzbox to monitor
  url = "localhost:2947"
`

func (f *Fritz) Description() string {
	return "Get statistics from a AVM FritzBox"
}

func (g *Gpsd) SampleConfig() string {
	return sampleConfig
}

func (g *Gpsd) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("fritzstats", func() telegraf.Input {
		return &Fritz{
			"Url":		"fritz.box",
		}
	})
}
