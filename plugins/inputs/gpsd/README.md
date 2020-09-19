# gpsd plugin

The gpsd plugin collects metrics from the gpsd daemon.

### Configuration:

```toml
[[inputs.gpsd]]
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
```

### Measurements & Fields

- gpsd_satcount
  - report_time (string)
  - used (float)
  - visible (float)
- gpsd_sky
  - report_time (string)
  - xdop (float)
  - ydop (float)
  - vdop (float)
  - tdop (float)
  - hdop (float)
  - pdop (float)
  - gdop (float)
- gpsd_tpv
  - report_time (string)
  - mode (int)
  - ept (float)
  - lat (float)
  - lon (float)
  - alt (float)
  - epx (float)
  - epy (float)
  - epv (float)
  - track (float)
  - speed (float)
  - climb (float)
  - epd (float)
  - eps (float)
  - epc (float)
- gpsd_gst
  - report_time (string)
  - rms (float)
  - major (float)
  - minor (float)
  - orient (float)
  - lat (float)
  - lon (float)
  - alt (float)
- gpsd_att
  - report_time (string)
  - heading (float)
  - magst (string)
  - pitch (float)
  - pitchst (string)
  - yaw (float)
  - yawst (string)
  - roll (float)
  - rollst (string)
  - dip (float)
  - maglen (float)
  - magx (float)
  - magy (float)
  - magz (float)
  - acclen (float)
  - accx (float)
  - accy (float)
  - accz (float)
  - gyrox (float)
  - gyroy (float)
  - depth (float)
  - temperature (float)
- gpsd_pps
  - realsec (float)
  - realmusec (float)
  - clocksec (float)
  - clockmusec (float)

### Tags

  - All measurements have the following tags:
    - device

### Example Output:
```toml
gpsd_satcount,device=/dev/ttyAMA0,host=rpi report_time=1594395790000000000,used=5,visible=11 1594395787000000000
gpsd_sky,device=/dev/ttyAMA0,host=rpi report_time=1594395153414630305,xdop=2.03,ydop=4.3,vdop=0.99,tdop=7.27,hdop=2.35,pdop=2.55,gdop=11.29 1594395153414630308
```
# gpsd plugin

The gpsd plugin collects metrics from the gpsd daemon.

### Configuration:

```toml
[[inputs.gpsd]]
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
```

### Measurements & Fields

- gpsd_satcount
  - device (string)
  - report_time (string)
  - used (float)
  - visible (float)
- gpsd_sky
  - report_time (string)
  - xdop (float)
  - ydop (float)
  - vdop (float)
  - tdop (float)
  - hdop (float)
  - pdop (float)
  - gdop (float)
- gpsd_tpv
  - report_time (string)
  - mode (int)
  - ept (float)
  - lat (float)
  - lon (float)
  - alt (float)
  - epx (float)
  - epy (float)
  - epv (float)
  - track (float)
  - speed (float)
  - climb (float)
  - epd (float)
  - eps (float)
  - epc (float)
- gpsd_gst
  - report_time (string)
  - rms (float)
  - major (float)
  - minor (float)
  - orient (float)
  - lat (float)
  - lon (float)
  - alt (float)
- gpsd_att
  - report_time (string)
  - heading (float)
  - magst (string)
  - pitch (float)
  - pitchst (string)
  - yaw (float)
  - yawst (string)
  - roll (float)
  - rollst (string)
  - dip (float)
  - maglen (float)
  - magx (float)
  - magy (float)
  - magz (float)
  - acclen (float)
  - accx (float)
  - accy (float)
  - accz (float)
  - gyrox (float)
  - gyroy (float)
  - depth (float)
  - temperature (float)
- gpsd_pps
  - realsec (float)
  - realmusec (float)
  - clocksec (float)
  - clockmusec (float)

### Tags

  - All measurements have the following tags:
    - device

### Example Output:
```toml
gpsd_satcount,device=/dev/ttyAMA0,host=rpi report_time=1594395790000000000,used=5,visible=11 1594395787000000000
gpsd_sky,device=/dev/ttyAMA0,host=rpi report_time=1594395153414630305,xdop=2.03,ydop=4.3,vdop=0.99,tdop=7.27,hdop=2.35,pdop=2.55,gdop=11.29 1594395153414630308
```
# gpsd plugin

The gpsd plugin collects metrics from the gpsd daemon.

### Configuration:

```toml
[[inputs.gpsd]]
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
```

### Measurements & Fields

- gpsd_satcount
  - report_time (string)
  - used (float)
  - visible (float)
- gpsd_sky
  - report_time (string)
  - xdop (float)
  - ydop (float)
  - vdop (float)
  - tdop (float)
  - hdop (float)
  - pdop (float)
  - gdop (float)
- gpsd_tpv
  - report_time (string)
  - mode (int)
  - ept (float)
  - lat (float)
  - lon (float)
  - alt (float)
  - epx (float)
  - epy (float)
  - epv (float)
  - track (float)
  - speed (float)
  - climb (float)
  - epd (float)
  - eps (float)
  - epc (float)
- gpsd_gst
  - report_time (string)
  - rms (float)
  - major (float)
  - minor (float)
  - orient (float)
  - lat (float)
  - lon (float)
  - alt (float)
- gpsd_att
  - report_time (string)
  - heading (float)
  - magst (string)
  - pitch (float)
  - pitchst (string)
  - yaw (float)
  - yawst (string)
  - roll (float)
  - rollst (string)
  - dip (float)
  - maglen (float)
  - magx (float)
  - magy (float)
  - magz (float)
  - acclen (float)
  - accx (float)
  - accy (float)
  - accz (float)
  - gyrox (float)
  - gyroy (float)
  - depth (float)
  - temperature (float)
- gpsd_pps
  - realsec (float)
  - realmusec (float)
  - clocksec (float)
  - clockmusec (float)

### Tags

  - All measurements have the following tags:
    - device

### Example Output:
```toml
gpsd_satcount,device=/dev/ttyAMA0,host=rpi report_time=1594395790000000000,used=5,visible=11 1594395787000000000
gpsd_sky,device=/dev/ttyAMA0,host=rpi report_time=1594395153414630305,xdop=2.03,ydop=4.3,vdop=0.99,tdop=7.27,hdop=2.35,pdop=2.55,gdop=11.29 1594395153414630308
```
