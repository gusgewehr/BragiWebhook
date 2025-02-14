package usecase

import (
	"BragiWebhooks/domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type receivedTextMessageUseCase struct {
	ReceivedTextMessageRepository domain.ReceivedTextMessageRepository
	Logger                        *zap.Logger
}

func (ru *receivedTextMessageUseCase) Send(ctx *gin.Context, message domain.ReceivedMessage) (*domain.ReceivedMessage, error) {

	_, err := ru.ReceivedTextMessageRepository.Send(ctx, message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func NewExampleUseCase(logger *zap.Logger, exampleRepository domain.ReceivedTextMessageRepository) domain.ReceivedTextMessageUseCase {
	return &receivedTextMessageUseCase{
		ReceivedTextMessageRepository: exampleRepository,
		Logger:                        logger,
	}
}
