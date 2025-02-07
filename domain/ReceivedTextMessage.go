package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReceivedTextMessage struct {
	gorm.Model
	Object string `json:"object"`
	Entry  []struct {
		Id      string `json:"id"`
		Changes []struct {
			Value struct {
				MessagingProduct string `json:"messaging_product"`
				Metadata         struct {
					DisplayPhoneNumber string `json:"display_phone_number"`
					PhoneNumberId      string `json:"phone_number_id"`
				} `json:"metadata"`
				Contacts []struct {
					Profile struct {
						Name string `json:"name"`
					} `json:"profile"`
					WaId string `json:"wa_id"`
				} `json:"contacts"`
				Messages []struct {
					From      string `json:"from"`
					Id        string `json:"id"`
					Timestamp string `json:"timestamp"`
					Text      struct {
						Body string `json:"body"`
					} `json:"text"`
					Type string `json:"type"`
				} `json:"messages"`
			} `json:"value"`
			Field string `json:"field"`
		} `json:"changes"`
	} `json:"entry"`
}

type ReceivedTextMessageRepository interface {
	Send(ctx *gin.Context, message ReceivedTextMessage) (*ReceivedTextMessage, error)
}

type ReceivedTextMessageUseCase interface {
	Send(ctx *gin.Context, message ReceivedTextMessage) (*ReceivedTextMessage, error)
}
