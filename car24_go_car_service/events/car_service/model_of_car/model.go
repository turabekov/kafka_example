package car

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"gitlab.udevs.io/car24/car24_go_car_service/config"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/event"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
	"gitlab.udevs.io/car24/car24_go_car_service/storage"
)

type ModelService struct {
	cfg     config.Config
	log     logger.Logger
	storage storage.StorageI
	kafka   event.KafkaI
	bot     *tgbotapi.BotAPI
}

func New(cfg config.Config, log logger.Logger, db *sqlx.DB, kafka event.KafkaI) *ModelService {
	return &ModelService{
		cfg:     cfg,
		log:     log,
		storage: storage.NewStoragePg(db),
		kafka:   kafka,
	}
}

func (c *ModelService) RegisterConsumers() {
	modelRoute := "v1.car_service.model."

	c.kafka.AddConsumer(
		modelRoute+"create", // consumer name
		modelRoute+"create", // topic
		modelRoute+"create", // group id
		c.Create,          // handlerFunction
	)

	c.kafka.AddConsumer(
		modelRoute+"update", // consumer name
		modelRoute+"update", // topic
		modelRoute+"update", // group id
		c.Update,          // handlerFunction,
	)

	c.kafka.AddConsumer(
		modelRoute+"delete", // consumer name
		modelRoute+"delete", // topic
		modelRoute+"delete", // group id
		c.Delete,          // handlerFunction,
	)

}
