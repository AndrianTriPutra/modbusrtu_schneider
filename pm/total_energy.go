package pm

import (
	"github.com/goburrow/modbus"

	. "pm_sch/convert"
)

func get_totalactivenergy(client modbus.Client) (data []int64, err error) {

	data_hex, err := client.ReadHoldingRegisters(3203, 8)
	if err == nil {
		data = BytesToInt64s(BIG_ENDIAN, HIGH_WORD_FIRST, data_hex)
	}

	return data, err
}

func get_totalreactivenergy(client modbus.Client) (data []int64, err error) {

	data_hex, err := client.ReadHoldingRegisters(3219, 8)
	if err == nil {
		data = BytesToInt64s(BIG_ENDIAN, HIGH_WORD_FIRST, data_hex)
	}

	return data, err
}
