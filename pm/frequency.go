package pm

import (
	"github.com/goburrow/modbus"

	. "pm_sch/convert"
)

func get_frequency(client modbus.Client) (data []float32, err error) {

	data_hex, err := client.ReadHoldingRegisters(3109, 1)
	if err == nil {
		data = BytesToFloat32s(BIG_ENDIAN, HIGH_WORD_FIRST, data_hex)
	}

	return data, err
}
