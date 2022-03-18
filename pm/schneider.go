package pm

import (
	"log"
	"time"

	"github.com/goburrow/modbus"

	. "pm_sch/json"
)

func Schneider() (payload string, err error) {

	// =========== init handler ===========
	handler := modbus.NewRTUClientHandler(Parameter.Port)
	handler.BaudRate = Parameter.Baudrate
	handler.DataBits = Parameter.DataBits
	handler.Parity = Parameter.Parity
	handler.StopBits = Parameter.StopBits
	handler.SlaveId = Parameter.SlaveId
	handler.Timeout = Parameter.Timeout * time.Millisecond
	// =========== init handler ===========

	var (
		state                                                                                  [10]bool
		data_volt, data_current, data_freq, data_active, data_reactive, data_apparent, data_pf []float32
		data_tactivenergy, data_treactivenergy, data_phasenergy                                []int64
	)

	timestamp := get_time(Parameter.Timezone)
	err = handler.Connect()
	if err != nil {
		log.Println("[ERROR] func Schneider, error open port " + Parameter.Port)
	} else {
		client := modbus.NewClient(handler)

		//0
		data_volt, err = get_voltage(client)
		if err != nil {
			state[0] = false
			log.Println("[ERROR] func Schneider, get_voltage modbus timeout")
		} else {
			state[0] = true
		}

		//1
		data_current, err = get_current(client)
		if err != nil {
			state[1] = false
			log.Println("[ERROR] func Schneider, get_current modbus timeout")
		} else {
			state[1] = true
		}

		//2
		data_freq, err = get_frequency(client)
		if err != nil {
			state[2] = false
			log.Println("[ERROR] func Schneider, get_frequency modbus timeout")
		} else {
			state[2] = true
		}

		//3
		data_active, err = get_active(client)
		if err != nil {
			state[3] = false
			log.Println("[ERROR] func Schneider, get_active modbus timeout")
		} else {
			state[3] = true
		}

		//4
		data_reactive, err = get_reactive(client)
		if err != nil {
			state[4] = false
			log.Println("[ERROR] func Schneider, get_reactive modbus timeout")
		} else {
			state[4] = true
		}

		//5
		data_apparent, err = get_apparent(client)
		if err != nil {
			state[5] = false
			log.Println("[ERROR] func Schneider, get_apparent modbus timeout")
		} else {
			state[5] = true
		}

		//6
		data_pf, err = get_pf(client)
		if err != nil {
			state[6] = false
			log.Println("[ERROR] func Schneider, get_pf modbus timeout")
		} else {
			state[6] = true
		}

		//7
		data_tactivenergy, err = get_totalactivenergy(client)
		if err != nil {
			state[7] = false
			log.Println("[ERROR] func Schneider, get_totalactivenergy modbus timeout")
		} else {
			state[7] = true
		}

		//8
		data_treactivenergy, err = get_totalreactivenergy(client)
		if err != nil {
			state[8] = false
			log.Println("[ERROR] func Schneider, get_totalreactivenergy modbus timeout")
		} else {
			state[8] = true
		}

		//9
		data_phasenergy, err = get_phasenergy(client)
		if err != nil {
			state[9] = false
			log.Println("[ERROR] func Schneider, get_phasenergy modbus timeout")
		} else {
			state[9] = true
		}

	}
	handler.Close()

	payload, err = Inject_Json(timestamp, state, data_volt, data_current, data_freq, data_active, data_reactive, data_apparent,
		data_pf, data_tactivenergy, data_treactivenergy, data_phasenergy)
	if err != nil {
		log.Printf("[ERROR] func Schneider, Inject_Json  : " + string(err.Error()))
	}

	return payload, err
}

func get_time(timezone string) (timestamp string) {
	loc, _ := time.LoadLocation(timezone)
	timeutcplus := time.Now().In(loc)

	timestamp = timeutcplus.Format(time.RFC3339)
	return timestamp
}
