package main

import (
	"log"

	. "pm_sch/json"
	. "pm_sch/pm"
)

func init() {
	Read_Config(".", "yaml", "config")
}

func main() {
	payload, err := Schneider()
	if err != nil {
		log.Printf("[ERROR] func main, error get data from Schneider: " + string(err.Error()))
	} else {
		log.Printf("Power_Meter Schneider   :\n%s", payload)
	}
}
