package sensor
import (
	"math"
	"github.com/b00lduck/raspberry-datalogger-dataservice-client"
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"strings"
	"strconv"
)

type Thermometer struct {
	oldValue float64
	precision float64
	code string
	filename string
}

const thermoDevicePrefix = "/sys/devices/w1_bus_master1/"
const thermoDeviceSuffix = "/w1_slave"

func NewThermometer(code string, filename string, precision float64) Thermometer {
	return Thermometer{
		oldValue: 0,
		precision: precision,
		code: code,
		filename: thermoDevicePrefix + filename + thermoDeviceSuffix,
	}
}

func (t *Thermometer) Process() {

	rawData, err := ioutil.ReadFile(t.filename)
	if err != nil {
		log.Error(err)
		return
	}

	foo := strings.Split(string(rawData), "t=")[1]
	bar := strings.Split(foo, "\n")[0]

	tempInt, err := strconv.Atoi(bar)
	if err != nil {
		log.Error(err)
		return
	}

	temp := float64(tempInt) / 1000000

	t.setNewReading(temp)
}

func (t *Thermometer) setNewReading(reading float64) {

	// precision reduction
	limitedPrecisionValue := round(reading / t.precision) * t.precision

	if math.Abs(float64(limitedPrecisionValue - t.oldValue)) > t.precision {
		if err := client.SendThermometerReading(t.code, limitedPrecisionValue); err != nil {
			log.Error(err)
		} else {
			t.oldValue = limitedPrecisionValue
		}
	}
}

func round(x float64) float64 {
	return math.Floor(x + 0.5)
}