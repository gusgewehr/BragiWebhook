package repository

import (
	"BragiWebhooks/domain"
	"BragiWebhooks/infrastructure"
	"encoding/json"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type receivedTextMessageRepository struct {
	Amqp   *infrastructure.Amqp
	Logger *zap.Logger
}

func (r *receivedTextMessageRepository) Send(ctx *gin.Context, message domain.ReceivedTextMessage) (*domain.ReceivedTextMessage, error) {

	body, err := json.Marshal(message)
	if err != nil {
		r.Logger.Error("failed to marshal received text message", zap.Error(err))
		return nil, err
	}

	err = r.Amqp.Ch.PublishWithContext(ctx,
		"",                // exchange
		r.Amqp.Queue.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		r.Logger.Error("failed to publish received text message", zap.Error(err))
		return nil, err
	}

	return &message, nil
}

func NewExampleRepository(amqp *infrastructure.Amqp, logger *zap.Logger) domain.ReceivedTextMessageRepository {
	return &receivedTextMessageRepository{Amqp: amqp, Logger: logger}
}
