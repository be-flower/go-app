package conf

import (
	"encoding/json"
	"fmt"
	"go-app/logger"
	"go-app/serve/Iot_mqtt"

	"github.com/spf13/viper"
)

var (
	Conf *Config
)

type MqttCloud struct {
	Host     string
	Port     int
	ClientId string
	UserName string
	PassWord string
	SubList  []string
	PubList  []string
	Qos      int
}

type MqttEdge struct {
	Host     string
	Port     int
	ClientId string
	UserName string
	PassWord string
	SubList  []string
	PubList  []string
	Qos      int
}

type RegisterTable struct {
	StartAddr int
	DataLen   int
	Type      string
	Name      string
}

type Device struct {
	Register      string
	Topic         string
	RegisterTable []RegisterTable
}

type TcpModbus struct {
	Enable   bool
	Host     string
	Port     int
	SlaveID  int
	Interval int
	Devices  []Device
}

type RtuModbus struct {
	Enable   bool
	Device   string
	BaudRate int
	DataBits int
	Parity   string
	StopBits int
	SlaveID  int
	Interval int
	Devices  []Device
}

type Config struct {
	ServeName string
	Serial    string
	MqttCloud MqttCloud
	MqttEdge  MqttEdge
	Tcpmodbus TcpModbus
	Rtumodbus RtuModbus
}

func NewConfig() *Config {
	return &Config{
		ServeName: "",
		Serial:    "",
		MqttCloud: MqttCloud{
			Host:     "127.0.0.1",
			Port:     1883,
			ClientId: "",
			UserName: "",
			PassWord: "",
			SubList:  nil,
			PubList:  nil,
			Qos:      0,
		},
		MqttEdge: MqttEdge{
			Host:     "127.0.0.1",
			Port:     1883,
			ClientId: "",
			UserName: "",
			PassWord: "",
			SubList:  nil,
			PubList:  nil,
			Qos:      0,
		},
		Tcpmodbus: TcpModbus{
			Enable:   false,
			Host:     "192.168.6.6",
			Port:     502,
			SlaveID:  1,
			Interval: 3,
			Devices:  nil,
		},
		Rtumodbus: RtuModbus{
			Enable:   false,
			Device:   "",
			BaudRate: 9600,
			DataBits: 8,
			Parity:   "N",
			StopBits: 1,
			SlaveID:  1,
			Interval: 3,
			Devices:  nil,
		},
	}
}

func InitConfig() {

	Conf = NewConfig()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Fatalf("Config file not found")
		} else {
			fmt.Println(err.Error())
		}
	}

	err := viper.Unmarshal(Conf)
	if err != nil {
		logger.Fatalf("unable to decode into struct, %v", err)
	}

	// 获取额外配置
	err = extraConfig()
	if err != nil {
		logger.Fatalf("get extra config fail!")
	}

	// 打印配置信息
	configuration, _ := json.Marshal(Conf)
	logger.Infof("Using conf: %v", string(configuration))
}

func extraConfig() error {
	// 上研需求，获取mqtt配置信息
	if Conf.ServeName == "iotmqtt" {
		err := Iot_mqtt.GetIotMqttConfig(Conf)
		if err != nil {
			logger.Errorf("GetIotMqttConfig error:", err)
			return err
		}
	}

	return nil
}
