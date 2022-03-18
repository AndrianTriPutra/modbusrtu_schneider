package json

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

var Parameter struct {
	Name     string
	Version  string
	Release  string
	Timezone string

	Port     string
	Baudrate int
	DataBits int
	Parity   string
	StopBits int
	SlaveId  byte
	Timeout  time.Duration
}

func Read_Config(path, tipe, name string) (err error) {
	// Initialize Configuratior
	config := viper.New()

	// Set Configuratior Configuration
	config.SetConfigName(name)
	config.AddConfigPath(path)
	config.SetConfigType(tipe)
	config.AutomaticEnv()

	err = config.ReadInConfig()
	if err != nil {
		log.Println("[ERROR] func Read_Config, error load file : " + string(err.Error()))
	}

	Parameter.Name = config.GetString("Name")
	Parameter.Version = config.GetString("Version")
	Parameter.Release = config.GetString("Release")
	Parameter.Timezone = config.GetString("Timezone")

	Parameter.Port = config.GetString("Port")
	Parameter.Baudrate = config.GetInt("Baudrate")
	Parameter.DataBits = config.GetInt("DataBits")
	Parameter.Parity = config.GetString("Parity")
	Parameter.StopBits = config.GetInt("StopBits")
	Parameter.SlaveId = byte(config.GetInt("SlaveId"))
	Parameter.Timeout = config.GetDuration("Timeout")

	log.Println(" ====================== Config ====================== ")
	log.Printf("Name           :%s\n", Parameter.Name)
	log.Printf("Version        :%s\n", Parameter.Version)
	log.Printf("Release        :%s\n", Parameter.Release)
	log.Printf("Timezone       :%s\n", Parameter.Timezone)
	fmt.Println()
	log.Printf("Port           :%s\n", Parameter.Port)
	log.Printf("Baudrate       :%v\n", Parameter.Baudrate)
	log.Printf("DataBits       :%v\n", Parameter.DataBits)
	log.Printf("Parity         :%s\n", Parameter.Parity)
	log.Printf("StopBits       :%v\n", Parameter.StopBits)
	log.Printf("SlaveId        :%v\n", Parameter.SlaveId)
	log.Printf("Timeout        :%v\n", Parameter.Timeout)
	log.Println(" ====================== Config ====================== ")
	fmt.Println()
	return err
}
