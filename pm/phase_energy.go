package pm

import (
	"github.com/goburrow/modbus"

	. "pm_sch/convert"
)

func get_phasenergy(client modbus.Client) (data []int64, err error) {

	data_hex, err := client.ReadHoldingRegisters(3517, 12)
	if err == nil {
		data = BytesToInt64s(BIG_ENDIAN, HIGH_WORD_FIRST, data_hex)
	}

	return data, err
}
