//+build wireinject

package registry

// The build tag makes sure the stub is not built in the final build.
import (
	"github.com/google/wire"

	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/adapter/controllers"
	pubsubAdapter "github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/adapter/gateways/pubsub"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/infrastructure/handler"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/infrastructure/pubsub"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/infrastructure/pubsub/mqtt"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/usecase/interactor"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/utils/logs/lggr"
)

func InitialiseMqtt(config pubsub.MqttConfig) (handler.MessageHandler, error) {
	wire.Build(
		lggr.NewLogger,
		pubsub.NewMqttConnect,
		mqtt.NewMqtt,
		pubsubAdapter.NewMessageMqtt,
		interactor.NewMessageInteractor,
		controllers.NewMessageController,
		handler.NewMessageHandler,
	)
	return nil, nil
}
