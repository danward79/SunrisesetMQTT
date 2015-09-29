# SunrisesetMQTT

Simple mqtt client that publishes topics that count down the sunset or sunrise..

Sunrise and Sunset MQTT events at -2, -1, 0, +1 hours. So an MQTT event at 2 hours, 1 hour before, at sunrise or sunset and one hour after.

```bash
home/sunrise/state -2
home/sunrise/state -1
home/sunrise/state 0
home/sunrise/state 1
```

### Command Line

The tool takes the command line switches below:

- x decimal longitude *default is set for Melbourne, Australia*
- y decimal latitude *default is set for Melbourne, Australia*
- s server IP:Port eg. 192.168.0.7:1883 *default is ":1883"*
- r retain MQTT state *default is set to false*


```bash
./sunrise -s 192.168.0.7:1883
```
