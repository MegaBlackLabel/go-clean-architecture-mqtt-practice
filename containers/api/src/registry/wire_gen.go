// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package registry

import (
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/adapter/controllers"
	pubsub2 "github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/adapter/gateways/pubsub"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/infrastructure/handler"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/infrastructure/pubsub"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/infrastructure/pubsub/mqtt"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/usecase/interactor"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/utils/logs/lggr"
)

// Injectors from wire.go:

func InitialiseMqtt(config pubsub.MqttConfig) (handler.MessageHandler, error) {
	que, err := pubsub.NewMqttConnect(config)
	if err != nil {
		return nil, err
	}
	logger := lggr.NewLogger()
	pubSub := mqtt.NewMqtt(que, logger)
	messageRepository := pubsub2.NewMessageMqtt(pubSub)
	messageInteractor := interactor.NewMessageInteractor(messageRepository)
	messageController := controllers.NewMessageController(messageInteractor)
	messageHandler := handler.NewMessageHandler(messageController)
	return messageHandler, nil
}