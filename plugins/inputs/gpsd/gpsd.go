package gpsd

import (
	"strconv"
	"strings"
	"sync"

	"github.com/stratoberry/go-gpsd"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

const (
	measurement = "gpsd"
	satcount_suffix = "satcount"
	sky_suffix = "sky"
	tpv_suffix = "tpv"
	gst_suffix = "gst"
	att_suffix = "att"
	pps_suffix = "pps"
)

type Gpsd struct {
	Url string
	ReportSatCount	bool	`toml:"gather_satcount"`
	ReportSKY	bool	`toml:"gather_sky"`
	ReportTPV	bool	`toml:"gather_tpv"`
	ReportGST	bool	`toml:"gather_gst"`
	ReportATT	bool	`toml:"gather_att"`
	ReportPPS	bool	`toml:"gather_pps"`

	done chan bool
	gps  *gpsd.Session
	wg   sync.WaitGroup
}

var sampleConfig = `
  ## gpsd daemon listening adress to connect to
  url = "localhost:2947"

  ## gather count of satellites visible and used
  # gather_satcount = true

  ## gather SKY reports
  # gather_sky = false

  ## gather TPV reports
  # gather_tpv = false

  ## gather GST reports
  # gather_gst = false

  ## gather ATT reports
  # gather_att = false

  ## gather pps reports
  # gather_pps = false
`

func (g *Gpsd) Description() string {
	return "Read satellite numbers seen and used from gpsd daemon"
}

func (g *Gpsd) SampleConfig() string {
	return sampleConfig
}

func (g *Gpsd) Gather(acc telegraf.Accumulator) error {
	return nil
}

func (g *Gpsd) SetupSKYReports(acc telegraf.Accumulator) func(r interface{}) {
	return func(r interface{}) {
		var report_time string = ""

		sky := r.(*gpsd.SKYReport)
		if ( !sky.Time.IsZero() ) {
			report_time = strconv.FormatInt(sky.Time.UnixNano(), 10)
		}

		if g.ReportSatCount {
			var used, visible int

			for _, sat := range sky.Satellites {
				if sat.Used {
					used++
				}
			}
			visible = len(sky.Satellites)

			fields := map[string]interface{}{
				"device":	sky.Device,
				"report_time":	report_time,
				"visible": 	visible,
				"used":    	used,
			}
			inputName := strings.Join([]string{measurement, satcount_suffix}, "_")
			acc.AddCounter(inputName, fields, nil)
		}

		if g.ReportSKY {
			fields := map[string]interface{}{
				"device":	sky.Device,
				"report_time":	report_time,
				"Xdop":		sky.Xdop,
				"Ydop":		sky.Ydop,
				"Vdop":		sky.Vdop,
				"Tdop":		sky.Tdop,
				"Hdop":		sky.Hdop,
				"Pdop":		sky.Pdop,
				"Gdop":		sky.Gdop,
			}
			inputName := strings.Join([]string{measurement, sky_suffix}, "_")
			acc.AddCounter(inputName, fields, nil)
		}
	}
}

func (g *Gpsd) SetupTPVReports(acc telegraf.Accumulator) func(r interface{}) {
	return func(r interface{}) {
		var report_time string = ""

		tpv := r.(*gpsd.TPVReport)
		if ( !tpv.Time.IsZero() ) {
			report_time = strconv.FormatInt(tpv.Time.UnixNano(), 10)
		}

		fields := map[string]interface{}{
			"device":	tpv.Device,
			"report_time":	report_time,
			"mode":		tpv.Mode,
			"ept":		tpv.Ept,
			"lat":		tpv.Lat,
			"lon":		tpv.Lon,
			"alt":		tpv.Alt,
			"epx":		tpv.Epx,
			"epy":		tpv.Epy,
			"epv":		tpv.Epv,
			"track":	tpv.Track,
			"speed":	tpv.Speed,
			"climb":	tpv.Climb,
			"epd":		tpv.Epd,
			"eps":		tpv.Eps,
			"epc":		tpv.Epc,
		}
		inputName := strings.Join([]string{measurement, tpv_suffix}, "_")
		acc.AddCounter(inputName, fields, nil)
	}
}

func (g *Gpsd) SetupGSTReports(acc telegraf.Accumulator) func(r interface{}) {
	return func(r interface{}) {
		var report_time string = ""

		gst := r.(*gpsd.GSTReport)
		if ( !gst.Time.IsZero() ) {
			report_time = strconv.FormatInt(gst.Time.UnixNano(), 10)
		}

		fields := map[string]interface{}{
			"device":	gst.Device,
			"report_time":	report_time,
			"rms":		gst.Rms,
			"major":	gst.Major,
			"minor":	gst.Minor,
			"orient":	gst.Orient,
			"lat":		gst.Lat,
			"lon":		gst.Lon,
			"alt":		gst.Alt,
		}
		inputName := strings.Join([]string{measurement, gst_suffix}, "_")
		acc.AddCounter(inputName, fields, nil)
	}
}

func (g *Gpsd) SetupATTReports(acc telegraf.Accumulator) func(r interface{}) {
	return func(r interface{}) {
		var report_time string = ""

		att := r.(*gpsd.ATTReport)
		if ( !att.Time.IsZero() ) {
			report_time = strconv.FormatInt(att.Time.UnixNano(), 10)
		}

		fields := map[string]interface{}{
			"device":	att.Device,
			"report_time":	report_time,
			"heading":	att.Heading,
			"magst":	att.MagSt,
			"pitch":	att.Pitch,
			"pitchst":	att.PitchSt,
			"yaw":		att.Yaw,
			"yawst":	att.YawSt,
			"roll":		att.Roll,
			"rollst":	att.RollSt,
			"dip":		att.Dip,
			"maglen":	att.MagLen,
			"magx":		att.MagX,
			"magy":		att.MagY,
			"magz":		att.MagZ,
			"acclen":	att.AccLen,
			"accx":		att.AccX,
			"accy":		att.AccY,
			"accz":		att.AccZ,
			"gyrox":	att.GyroX,
			"gyroy":	att.GyroY,
			"depth":	att.Depth,
			"temperature":	att.Temperature,	
		}
		inputName := strings.Join([]string{measurement, att_suffix}, "_")
		acc.AddCounter(inputName, fields, nil)
	}
}

func (g *Gpsd) SetupPPSReports(acc telegraf.Accumulator) func(r interface{}) {
	return func(r interface{}) {
		pps := r.(*gpsd.PPSReport)

		fields := map[string]interface{}{
			"device":	pps.Device,
			"realsec":	pps.RealSec,
			"realmusec":	pps.RealMusec,
			"clocksec":	pps.ClockSec,
			"clockmusec":	pps.ClockMusec,
		}
		inputName := strings.Join([]string{measurement, pps_suffix}, "_")
		acc.AddCounter(inputName, fields, nil)
	}
}

func (g *Gpsd) Start(acc telegraf.Accumulator) error {
	var err error

	g.gps, err = g.createGPSDSession(g.Url)
	if err != nil {
		return err
	}

	if ( g.ReportSatCount || g.ReportSKY ) {
		skyfilter := g.SetupSKYReports(acc)
		g.gps.AddFilter("SKY", skyfilter)
	}

	if ( g.ReportTPV ) {
		tpvfilter := g.SetupTPVReports(acc)
		g.gps.AddFilter("TPV", tpvfilter)
	}

	if ( g.ReportGST ) {
		gstfilter := g.SetupGSTReports(acc)
		g.gps.AddFilter("GST", gstfilter)
	}

	if ( g.ReportATT ) {
		attfilter := g.SetupATTReports(acc)
		g.gps.AddFilter("ATT", attfilter)
	}

	if ( g.ReportPPS ) {
		ppsfilter := g.SetupPPSReports(acc)
		g.gps.AddFilter("PPS", ppsfilter)
	}

	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		g.done = g.gps.Watch()
		<-g.done
	}()

	return nil
}

func (g *Gpsd) Stop() {
	g.done <- true
	g.wg.Wait()
}

func init() {
	inputs.Add("gpsd", func() telegraf.Input {
		return &Gpsd{
			ReportSatCount: true,
			ReportSKY:	false,
			ReportTPV:	false,
			ReportGST:	false,
			ReportATT:	false,
			ReportPPS:	false,
		}
	})
}

func (g *Gpsd) createGPSDSession(url string) (*gpsd.Session, error) {
	return gpsd.Dial(url)
}
