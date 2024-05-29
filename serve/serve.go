package serve

import (
	"go-app/conf"
	"go-app/logger"
	"go-app/serve/productionline/web"
)

func StartServe() {

	// serve
	switch conf.Conf.ServeName {
	//case "cgxi":
	//	cgxi.CgxiDealModbus(config, mqttClient)
	//case "test":
	//	modbus.DealModbus(config, mqttClient)
	//case "xinjie":
	//	xinjie.XinJieDealModbus(config, mqttClient)
	//case "digitalTwin":
	//	digitaltwin.DigitalTwinDealModbus(config, mqttClient)
	//case "iotmqtt":
	//	Iot_mqtt.IotMqttImpl(config, mqttClient)
	case "productionline_web":
		web.ProductionLineWebDeal()
	default:
		logger.Errorf("config error, no serve name")
	}

}
