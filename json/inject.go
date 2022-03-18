package json

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type Voltage struct {
	State  bool
	L1_L2  string
	L2_L3  string
	L3_L1  string
	LL_avg string
	L1_N   string
	L2_N   string
	L3_N   string
	LN_avg string
}

type Current struct {
	State     bool
	Phase_1   string
	Phase_2   string
	Phase_3   string
	Phase_avg string
}

type Active struct {
	State   bool
	Phase_1 string
	Phase_2 string
	Phase_3 string
	Total   string
}

type Reactive struct {
	State bool
	Total string
}

type Apparent struct {
	State bool
	Total string
}

type Power struct {
	Active   Active
	Reactive Reactive
	Apparent Apparent
}

type PowerFactor struct {
	State       bool
	PowerFactor string
}

type Frequency struct {
	State     bool
	Frequency string
}

type Active_Energy struct {
	State_1        bool
	Total_Import   string
	Total_Export   string
	State_2        bool
	Import_Phase_1 string
	Import_Phase_2 string
	Import_Phase_3 string
}

type Reactive_Energy struct {
	State        bool
	Total_Import string
	Total_Export string
}

var Data_PM struct {
	Timestamp       string
	Voltage         Voltage
	Current         Current
	Frequency       Frequency
	Power           Power
	PowerFactor     PowerFactor
	Active_Energy   Active_Energy
	Reactive_Energy Reactive_Energy
}

func Inject_Json(timestamp string, state [10]bool, data_volt, data_current, data_freq,
	data_active, data_reactive, data_apparent, data_pf []float32,
	data_tactivenergy, data_treactivenergy, data_phasenergy []int64) (payload string, err error) {

	//log.Printf("timestamp                    :%s", timestamp)
	Data_PM.Timestamp = timestamp

	// log.Printf("state[0]                    :%v", state[0])
	// log.Printf("voltage                      :%v", data_volt)
	Data_PM.Voltage.State = state[0]
	if state[0] {
		Data_PM.Voltage.L1_L2 = fmt.Sprintf("%.2f", data_volt[0])
		Data_PM.Voltage.L2_L3 = fmt.Sprintf("%.2f", data_volt[1])
		Data_PM.Voltage.L3_L1 = fmt.Sprintf("%.2f", data_volt[2])
		Data_PM.Voltage.LL_avg = fmt.Sprintf("%.2f", data_volt[3])
		Data_PM.Voltage.L1_N = fmt.Sprintf("%.2f", data_volt[4])
		Data_PM.Voltage.L2_N = fmt.Sprintf("%.2f", data_volt[5])
		Data_PM.Voltage.L3_N = fmt.Sprintf("%.2f", data_volt[6])
		Data_PM.Voltage.LN_avg = fmt.Sprintf("%.2f", data_volt[8])
	} else {
		Data_PM.Voltage.L1_L2 = "NA"
		Data_PM.Voltage.L2_L3 = "NA"
		Data_PM.Voltage.L3_L1 = "NA"
		Data_PM.Voltage.LL_avg = "NA"
		Data_PM.Voltage.L1_N = "NA"
		Data_PM.Voltage.L2_N = "NA"
		Data_PM.Voltage.L3_N = "NA"
		Data_PM.Voltage.LN_avg = "NA"
	}

	// log.Printf("state[1]                     :%v", state[1])
	// log.Printf("current                      :%v", data_current)
	Data_PM.Current.State = state[1]
	if state[1] {
		Data_PM.Current.Phase_1 = fmt.Sprintf("%.2f", data_current[0])
		Data_PM.Current.Phase_2 = fmt.Sprintf("%.2f", data_current[1])
		Data_PM.Current.Phase_3 = fmt.Sprintf("%.2f", data_current[2])
		Data_PM.Current.Phase_avg = fmt.Sprintf("%.2f", data_current[5])
	} else {
		Data_PM.Current.Phase_1 = "NA"
		Data_PM.Current.Phase_2 = "NA"
		Data_PM.Current.Phase_3 = "NA"
		Data_PM.Current.Phase_avg = "NA"
	}

	// log.Printf("state[2]                     :%v", state[2])
	// log.Printf("frequency                    :%v", data_freq)
	Data_PM.Frequency.State = state[2]
	if state[2] {
		Data_PM.Frequency.Frequency = fmt.Sprintf("%.2f", data_freq[0])
	} else {
		Data_PM.Frequency.Frequency = "NA"
	}

	// log.Printf("state[3]                     :%v", state[3])
	// log.Printf("active                       :%v", data_active)
	Data_PM.Power.Active.State = state[3]
	if state[3] {
		Data_PM.Power.Active.Phase_1 = fmt.Sprintf("%.5f", data_active[0])
		Data_PM.Power.Active.Phase_2 = fmt.Sprintf("%.5f", data_active[1])
		Data_PM.Power.Active.Phase_3 = fmt.Sprintf("%.5f", data_active[2])
		Data_PM.Power.Active.Total = fmt.Sprintf("%.5f", data_active[3])
	} else {
		Data_PM.Power.Active.Phase_1 = "NA"
		Data_PM.Power.Active.Phase_2 = "NA"
		Data_PM.Power.Active.Phase_3 = "NA"
		Data_PM.Power.Active.Total = "NA"
	}

	// log.Printf("state[4]                     :%v", state[4])
	// log.Printf("reactive                     :%v", data_reactive)
	Data_PM.Power.Reactive.State = state[4]
	if state[4] {
		Data_PM.Power.Reactive.Total = fmt.Sprintf("%.5f", data_reactive[0])
	} else {
		Data_PM.Power.Reactive.Total = "NA"
	}

	// log.Printf("state[5]                     :%v", state[5])
	// log.Printf("apparent                     :%v", data_apparent)
	Data_PM.Power.Apparent.State = state[5]
	if state[5] {
		Data_PM.Power.Apparent.Total = fmt.Sprintf("%.5f", data_reactive[0])
	} else {
		Data_PM.Power.Apparent.Total = "NA"
	}

	// log.Printf("state[6]                     :%v", state[6])
	// log.Printf("power factor                 :%v", data_pf)
	Data_PM.PowerFactor.State = state[6]
	if state[6] {
		Data_PM.PowerFactor.PowerFactor = fmt.Sprintf("%.2f", data_pf[0])
	} else {
		Data_PM.PowerFactor.PowerFactor = "NA"
	}

	// log.Printf("state[7]                     :%v", state[7])
	// log.Printf("total active energy          :%v", data_tactivenergy)
	Data_PM.Active_Energy.State_1 = state[7]
	if state[7] {
		Data_PM.Active_Energy.Total_Import = strconv.FormatInt(data_tactivenergy[0], 10)
		Data_PM.Active_Energy.Total_Export = strconv.FormatInt(data_tactivenergy[1], 10)
	} else {
		Data_PM.Active_Energy.Total_Import = "NA"
		Data_PM.Active_Energy.Total_Export = "NA"
	}

	// log.Printf("state[9]                     :%v", state[9])
	// log.Printf("data phase energy            :%v", data_phasenergy)
	Data_PM.Active_Energy.State_2 = state[9]
	if state[9] {
		Data_PM.Active_Energy.Import_Phase_1 = strconv.FormatInt(data_phasenergy[0], 10)
		Data_PM.Active_Energy.Import_Phase_2 = strconv.FormatInt(data_phasenergy[1], 10)
		Data_PM.Active_Energy.Import_Phase_3 = strconv.FormatInt(data_phasenergy[2], 10)
	} else {
		Data_PM.Active_Energy.Import_Phase_1 = "NA"
		Data_PM.Active_Energy.Import_Phase_2 = "NA"
		Data_PM.Active_Energy.Import_Phase_3 = "NA"
	}

	// log.Printf("state[8]                     :%v", state[8])
	// log.Printf("total reactive energy        :%v", data_treactivenergy)
	Data_PM.Reactive_Energy.State = state[8]
	if state[8] {
		Data_PM.Reactive_Energy.Total_Import = strconv.FormatInt(data_treactivenergy[0], 10)
		Data_PM.Reactive_Energy.Total_Export = strconv.FormatInt(data_treactivenergy[1], 10)
	} else {
		Data_PM.Reactive_Energy.Total_Import = "NA"
		Data_PM.Reactive_Energy.Total_Export = "NA"
	}

	var write []byte
	write, err = json.MarshalIndent(Data_PM, " ", " ")
	if err != nil {
		log.Printf("[ERROR] func Inject_Json, create json Power_Meter  : " + string(err.Error()))
	} else {
		payload = string(write)
	}

	return payload, err
}
